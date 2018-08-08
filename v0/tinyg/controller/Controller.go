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
	"bufio"
	"github.com/golang/glog"
	"github.com/itschleemilch/huanyango/v1/vfdio"
	tgjson "github.com/itschleemilch/tinyg-api/v0/tinyg/json"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	linesToSendDefault          int32         = 4
	lineQueueLength             int           = 10000
	pollMachinePositionInterval time.Duration = 5000 * time.Millisecond
	pollWorkingPositionInterval time.Duration = 5000 * time.Millisecond
	pollMachineOffsetInterval   time.Duration = 7000 * time.Millisecond
	pollFullState               time.Duration = 5000 * time.Millisecond
)

// TinygController holds the internal hardware handels and publishes
// methods for control and getting the state of the machine.
type TinygController struct {
	port               io.ReadWriteCloser
	writeLock          sync.Mutex
	initOnce           sync.Once
	lineQueue          chan string
	lineQueueEmptyFlag bool
	linesToSend        int32
	lineQueueLock      sync.Mutex
	exit               bool
	tinygState         tgjson.TResponse
	lastResponseTime   time.Time
	// Optional reference to a Huanyan VFD
	VfdOutput *vfdio.HyInverter
}

func NewController() (controller *TinygController, err error) {
	controller = &TinygController{}
	controller.tinygState = tgjson.TResponse{}
	return
}

// Open inits the hardware handels. A suitable portName could be "COM3" or "/dev/ttyUSB0"
func (o *TinygController) Open(portName string) (err error) {
	if o == nil {
		glog.Error("Open can not be called from a nil pointer")
	}
	o.initOnce.Do(func() {
		o.lineQueue = make(chan string, lineQueueLength)
		o.linesToSend = linesToSendDefault
		atomic.StoreInt32(&o.linesToSend, linesToSendDefault)
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
			go o.statePolling()
			if o.VfdOutput != nil {
				o.VfdOutput.GCode("S0 M5")
			}
		} else {
			glog.Error(err)
		}
	})
	return
}

func (o *TinygController) Close() {
	o.port.Close()
	o.exit = true
	o.initOnce = sync.Once{}
	o.writeLock = sync.Mutex{}
	o.lineQueueLock = sync.Mutex{}
}

func (o *TinygController) serialRxLoop() {
	lineScanner := bufio.NewScanner(o.port)
	for !o.exit && lineScanner.Scan() {
		// Handle data
		jsonResponse := lineScanner.Text()
		glog.Infoln("Tinyg Output: ", jsonResponse)
		data, parseErr := tgjson.ParseResponse([]byte(jsonResponse))
		if parseErr == nil {
			o.lastResponseTime = time.Now()
			if len(data.ResponseFooter) == 3 { // only react to fixed-interval status reports
				linesToSendNow := atomic.LoadInt32(&o.linesToSend)
				if linesToSendNow < linesToSendDefault {
					atomic.AddInt32(&o.linesToSend, 1)
				}
			}
			o.tinygState.UpdateFrom(data) // overwrite buffered states with received values
		} else {
			glog.Warning("Input Error: ", jsonResponse, parseErr)
		}
	}
}

func (o *TinygController) handleVfdCommand(cmd string) {
	// If a VFD is connected, output G-Code to the VFD processor
	if o.VfdOutput != nil {
		glog.Infoln("Vfd detected")
		o.VfdOutput.GCodeWaiting(cmd)
		vfdOk, _, _ := o.VfdOutput.Processed()
		glog.Infoln("Vfd waiting for processing...")
		for !vfdOk {
			time.Sleep(5 * time.Millisecond)
			vfdOk, _, _ = o.VfdOutput.Processed()
		}
		glog.Infoln("Vfd handled.")
	}
}

func (o *TinygController) serialTxLoop() {
	for !o.exit {
		cmd := <-o.lineQueue
		if o.lineQueueEmptyFlag {
			for len(o.lineQueue) > 0 {
				cmd = <-o.lineQueue
			}
			o.lineQueueEmptyFlag = false
		} else {
			if len(cmd) > 0 {
				for o.linesToSend <= 0 {
					time.Sleep(10 * time.Millisecond)
				}
				o.handleVfdCommand(cmd)
				// Serial Output
				cmd = string(append([]byte(cmd), []byte{0x0A}...)) // Append new line character
				if !o.lineQueueEmptyFlag {                         // Again, check for flush flag
					glog.Infoln("TX: '", cmd, "'")
					o.writeLock.Lock()
					o.port.Write([]byte(cmd))
					o.writeLock.Unlock()
					atomic.AddInt32(&o.linesToSend, -1)
				}
			}
		}
	}

}

func (o *TinygController) Flush() {
	o.lineQueueEmptyFlag = true
	o.writeLock.Lock()
	o.port.Write([]byte{0x04})        // Send ^D flush command
	o.port.Write([]byte("{clr:n}\n")) // Clear alarms
	o.writeLock.Unlock()
	atomic.StoreInt32(&o.linesToSend, linesToSendDefault)
}

// RefreshState sends all required commands to Tinyg for reconstructing
// the full machine state.
func (o *TinygController) RefreshState() {
	go func() {
		//o.SendCommandWaiting(tgjson.CommandEmptyLine) // Flush any unfinished command at tinyg rx buffer
		//o.SendCommandWaiting(tgjson.CommandSetRxModeLine)
		//o.SendCommandWaiting(tgjson.CommandSetFlowControlCts)
		/*
			o.SendCommandWaiting(tgjson.CommandRequestStatus)
			o.SendCommandWaiting(tgjson.CommandRequestG54Offset)
			o.SendCommandWaiting(tgjson.CommandRequestG55Offset)
			o.SendCommandWaiting(tgjson.CommandRequestG56Offset)
			o.SendCommandWaiting(tgjson.CommandRequestG57Offset)
			o.SendCommandWaiting(tgjson.CommandRequestG58Offset)
			o.SendCommandWaiting(tgjson.CommandRequestG59Offset)
			o.SendCommandWaiting(tgjson.CommandRequestG92Offset)
			o.SendCommandWaiting(tgjson.CommandRequestExceptionReport)
			o.SendCommandWaiting(tgjson.CommandRequestQueueReport)
			o.SendCommandWaiting(tgjson.CommandRequestRxBuffer)
			o.SendCommandWaiting(tgjson.CommandRequestHardwarePlatform)
			o.SendCommandWaiting(tgjson.CommandRequestHardwareVersion)
			o.SendCommandWaiting(tgjson.CommandRequestPositionG28)
			o.SendCommandWaiting(tgjson.CommandRequestPositionG28)
			o.SendCommandWaiting(tgjson.CommandRequestWorkingPosition)
			o.SendCommandWaiting(tgjson.CommandRequestMachineAbsolutePosition)
		*/
	}()
}

func (o *TinygController) statePolling() {
	tickAbsMachineCoords := time.NewTicker(pollMachinePositionInterval)
	defer tickAbsMachineCoords.Stop()
	tickWorkingCoords := time.NewTicker(pollMachineOffsetInterval)
	defer tickWorkingCoords.Stop()
	tickOffsets := time.NewTicker(pollMachineOffsetInterval)
	defer tickOffsets.Stop()
	tickFullState := time.NewTicker(pollFullState)
	defer tickFullState.Stop()
	for !o.exit {
		select {
		case <-tickAbsMachineCoords.C:
			o.write(tgjson.CommandRequestMachineAbsolutePosition, false)
			break
		case <-tickWorkingCoords.C:
			o.write(tgjson.CommandRequestWorkingPosition, false)
			break
		case <-tickOffsets.C:
			if o.tinygState.ResponseData.StatusReport != nil &&
				o.tinygState.ResponseData.StatusReport.CoordinateSystem != nil {
				switch *o.tinygState.ResponseData.StatusReport.CoordinateSystem {
				case tgjson.CoordinateSystemG54:
					o.write(tgjson.CommandRequestG54Offset, false)
					break
				case tgjson.CoordinateSystemG55:
					o.write(tgjson.CommandRequestG55Offset, false)
					break
				case tgjson.CoordinateSystemG56:
					o.write(tgjson.CommandRequestG56Offset, false)
					break
				case tgjson.CoordinateSystemG57:
					o.write(tgjson.CommandRequestG57Offset, false)
					break
				case tgjson.CoordinateSystemG58:
					o.write(tgjson.CommandRequestG58Offset, false)
					break
				case tgjson.CoordinateSystemG59:
					o.write(tgjson.CommandRequestG59Offset, false)
					break
				default:
					break
				}
			}
			o.write(tgjson.CommandRequestG92Offset, false)
			break
		case <-tickFullState.C:
			o.write(tgjson.CommandRequestStatus, false)
			break
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func eraseGcodeComments(in string) string {
	out1 := strings.Split(in, "(")
	out2 := strings.Split(out1[0], ";")
	return out2[0]
}

func (o *TinygController) writeLines(cmds []string, queue bool) (inserted bool) {
	o.lineQueueLock.Lock()
	if queue {
		inserted = true

	} else {
		if len(o.lineQueue) == 0 {
			inserted = true
		}
	}
	if inserted {
		for _, cmd := range cmds {
			cmd = eraseGcodeComments(cmd)
			cmd = strings.TrimSuffix(cmd, "\n")
			cmd = strings.TrimSuffix(cmd, "\r")
			cmd = strings.TrimSpace(cmd)
			o.lineQueue <- cmd
		}
	}
	o.lineQueueLock.Unlock()
	return
}

func (o *TinygController) write(cmd string, queue bool) bool {
	return o.writeLines([]string{cmd}, queue)
}

// SendData sends a raw command over the low priority queue.
// Returns true if the internal queue has accepted the command.
func (o *TinygController) Write(cmd string) bool {
	return o.write(cmd, true)
}

func (o *TinygController) WriteLines(cmds []string) bool {
	return o.writeLines(cmds, true)
}

// TinygReset performs a software reset of the hardware.
func (o *TinygController) TinygReset() error {
	o.writeLock.Lock()
	o.port.Write([]byte{24}) // CTRL-X
	o.writeLock.Unlock()
	time.Sleep(5 * time.Second)
	o.Flush()
	return nil
}

func (o *TinygController) FeedHold() error {
	o.writeLock.Lock()
	o.port.Write([]byte{'!'})
	o.writeLock.Unlock()
	return nil
}

func (o *TinygController) FeedResume() error {
	o.writeLock.Lock()
	o.port.Write([]byte{'~'})
	o.writeLock.Unlock()
	return nil

}

func (o *TinygController) Online() bool {
	return time.Now().Sub(o.lastResponseTime).Seconds() < 1.0
}

func (o *TinygController) StateJson() []byte {
	return o.tinygState.Json()
}
