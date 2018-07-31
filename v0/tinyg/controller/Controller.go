// Copyright (C) 2018 Sebastian Schleemilch
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301  USA

package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	tgjson "github.com/itschleemilch/tinyg-api/v0/tinyg/json"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"sync"
	"time"
)

const (
	dataQueueLength             int           = 10000
	cmdQueueLength              int           = 100
	pollMachinePositionInterval time.Duration = 3 * time.Second
	pollWorkingPositionInterval time.Duration = 3 * time.Second
	pollMachineOffsetInterval   time.Duration = 5 * time.Second
)

// TinygController holds the internal hardware handels and publishes
// methods for control and getting the state of the machine.
type TinygController struct {
	port             io.ReadWriteCloser
	once             sync.Once
	dataQueue        chan string
	commandQueue     chan string
	exit             bool
	tinygState       tgjson.TResponse
	lastResponseTime time.Time
}

func NewController() (controller *TinygController, err error) {
	controller = &TinygController{}
	controller.tinygState = tgjson.TResponse{}
	return
}

// Open inits the hardware handels. A suitable portName could be "COM3" or "/dev/ttyUSB0"
func (o *TinygController) Open(portName string) (err error) {
	if o == nil {
		fmt.Errorf("Open can not be called from a nil pointer")
	}
	//o.once.Do(func() {
	o.dataQueue = make(chan string, dataQueueLength)
	o.commandQueue = make(chan string, cmdQueueLength)
	portOptions := serial.OpenOptions{
		PortName:          portName,
		BaudRate:          115200,
		DataBits:          8,
		ParityMode:        serial.PARITY_NONE,
		StopBits:          1,
		RTSCTSFlowControl: true,
		MinimumReadSize:   1,
	}
	o.port, err = serial.Open(portOptions)
	if err == nil {
		go o.serialRxLoop()
		go o.serialTxLoop()
		o.SendCommand(tgjson.CommandEmptyLine) // Flush any unfinished command at tinyg rx buffer
		o.SendCommand(tgjson.CommandSetRxModeLine)
	}
	//})
	return
}

func (o *TinygController) Close() {
	o.port.Close()
	o.exit = true
}

func (o *TinygController) serialRxLoop() {
	rxBuffer1 := make([]byte, 0)    // concat-ed reads
	rxBuffer0 := make([]byte, 1000) // single read
	for !o.exit {
		nRead, err := o.port.Read(rxBuffer0)
		if err == nil && nRead > 0 {
			// combine previous received data:
			rxBuffer1 = append(rxBuffer1, rxBuffer0[:nRead]...)
			// check for new line
			for j := 0; j < len(rxBuffer1); j++ {
				if rxBuffer1[j] == '\n' {
					jsonResponse := rxBuffer1[:j]
					if len(rxBuffer1) >= j {
						rxBuffer1 = rxBuffer1[j+1:]
					} else {
						rxBuffer1 = make([]byte, 0)
					}
					// Handle data
					data, parseErr := tgjson.ParseResponse(jsonResponse)
					if parseErr == nil {
						o.lastResponseTime = time.Now()
						o.tinygState.UpdateFrom(data) // overwrite buffered states with received values
						encState, encErr := json.Marshal(&o.tinygState)
						if encErr == nil {
							fmt.Println(string(encState))
							fmt.Println()
							fmt.Println()
						} else {
							panic(encErr)
						}
					} else {
						fmt.Println("Input Error: ", string(jsonResponse))
						fmt.Println(parseErr)
					}
				}
			}
		}
	}
}

func (o *TinygController) serialTxLoop() {
	go func() {
		time.Sleep(1 * time.Second)
		o.SendCommand(tgjson.CommandRequestExceptionReport)
		o.SendCommand(tgjson.CommandRequestFirmwareVersion)
		o.SendCommand(tgjson.CommandRequestStatus)
		o.SendCommand(tgjson.CommandRequestMachineAbsolutePosition)
		time.Sleep(2 * time.Second)
		o.SendData("G0 X5")
		time.Sleep(1 * time.Second)
		o.SendCommand(tgjson.CommandRequestStatus)
		o.SendCommand(tgjson.CommandRequestMachineAbsolutePosition)

	}()
	txChan := make(chan string, 16)
	// Separate stream processing of (high priority) commands
	go func() {
		for !o.exit {
			newCmd := <-o.commandQueue
			txChan <- newCmd
			// Check number of free packet slots on tinyg:
			if len(o.tinygState.ResponseFooter) == 3 { // check json footer format
				for o.tinygState.ResponseFooter[2] < 3 { // always keep at least 2 slot free!
					time.Sleep(100 * time.Millisecond)
				}
			} else {
				time.Sleep(100 * time.Millisecond) // no free slot information received yet
			}
		}
	}()
	// Seperate stream processing of (low priority) data commands
	go func() {
		for !o.exit {
			newData := <-o.dataQueue
			txChan <- newData
			// Check number of free packet slots on tinyg:
			if len(o.tinygState.ResponseFooter) == 3 { // check json footer format
				for o.tinygState.ResponseFooter[2] < 4 { // always keep at least 3 slot free!
					time.Sleep(100 * time.Millisecond)
				}
			} else {
				time.Sleep(500 * time.Millisecond) // no free slot information received yet
			}
		}
	}()
	// Actual output process:
	for !o.exit {
		cmd := <-txChan
		if len(cmd) > 0 {
			fmt.Println("Sending ", cmd)
			cmd = string(append([]byte(cmd), []byte{0x0A}...)) // Append new line character
			o.port.Write([]byte(cmd))
		}
	}
}

// SendCommand sends a raw command over the high priority queue. Do not send GCode!
// Returns nil if the internal queue has accepted the command.
func (o *TinygController) SendCommand(cmd tgjson.TinygCommand) error {
	select {
	case o.commandQueue <- string(cmd):
		break
	default:
		return errors.New("Buffer full")
	}
	return nil
}

// SendData sends a raw command over the low priority queue.
// Returns nil if the internal queue has accepted the command.
func (o *TinygController) SendData(cmd string) error {
	select {
	case o.dataQueue <- cmd:
		break
	default:
		return errors.New("Buffer full")
	}
	return nil
}

// TinygReset performs a software reset of the hardware.
func (o *TinygController) TinygReset() error {
	return o.SendCommand(tgjson.TinygCommand([]byte{24})) // 24	Ctrl-X	Cancel
}

func (o *TinygController) FeedHold() error {
	return o.SendCommand(tgjson.CommandFeedHold)
}

func (o *TinygController) FeedResume() error {
	return o.SendCommand(tgjson.CommandFeedResume)
}

func (o *TinygController) Online() bool {
	return time.Now().Sub(o.lastResponseTime).Seconds() < 1.0
}
