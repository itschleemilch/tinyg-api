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

package json

import (
	jsjson "encoding/json"
)

// TReceiveObjects contains many of the possible return values
// from tinyg responses. They are all nullable (pointers) since
// only one item is returned per request.
type TReceiveObjects struct {
	FirmwareVersion         *float64       `json:"fv"`
	HardwarePlatform        *int           `json:"hp"`
	HardwareVersion         *int           `json:"hv"`
	StatusReport            *TStatusReport `json:"sr"`
	ExceptionReport         *int           `json:"ex"`
	QueueReport             *int           `json:"qr"`
	RxBufferReport          *int           `json:"rx"`
	AbsoluteMachinePosition *TOffset       `json:"mpo"`
	WorkingPosition         *TOffset       `json:"pos"`
	SavedPositionG28        *TOffset       `json:"g28"`
	SavedPositionG30        *TOffset       `json:"g30"`
	OffsetG54               *TOffset       `json:"g54"`
	OffsetG55               *TOffset       `json:"g55"`
	OffsetG56               *TOffset       `json:"g56"`
	OffsetG57               *TOffset       `json:"g57"`
	OffsetG58               *TOffset       `json:"g58"`
	OffsetG59               *TOffset       `json:"g59"`
	AddonOffsetG92          *TOffset       `json:"g92"`
	RxMode                  *TRxMode       `json:"rxm"`
}

func (dst *TReceiveObjects) UpdateFrom(src *TReceiveObjects) {
	if src.FirmwareVersion != nil {
		dst.FirmwareVersion = src.FirmwareVersion
	}
	if src.HardwarePlatform != nil {
		dst.HardwarePlatform = src.HardwarePlatform
	}
	if dst.StatusReport == nil {
		dst.StatusReport = &TStatusReport{}
	}
	dst.StatusReport.UpdateFrom(src.StatusReport)
	if src.ExceptionReport != nil {
		dst.ExceptionReport = src.ExceptionReport
	}
	if src.QueueReport != nil {
		dst.QueueReport = src.QueueReport
	}
	if src.RxBufferReport != nil {
		dst.RxBufferReport = src.RxBufferReport
	}
	if dst.AbsoluteMachinePosition == nil {
		dst.AbsoluteMachinePosition = &TOffset{}
	}
	if dst.WorkingPosition == nil {
		dst.WorkingPosition = &TOffset{}
	}
	if dst.SavedPositionG28 == nil {
		dst.SavedPositionG28 = &TOffset{}
	}
	if dst.SavedPositionG30 == nil {
		dst.SavedPositionG30 = &TOffset{}
	}
	if dst.OffsetG54 == nil {
		dst.OffsetG54 = &TOffset{}
	}
	if dst.OffsetG55 == nil {
		dst.OffsetG55 = &TOffset{}
	}
	if dst.OffsetG56 == nil {
		dst.OffsetG56 = &TOffset{}
	}
	if dst.OffsetG57 == nil {
		dst.OffsetG57 = &TOffset{}
	}
	if dst.OffsetG58 == nil {
		dst.OffsetG58 = &TOffset{}
	}
	if dst.OffsetG59 == nil {
		dst.OffsetG59 = &TOffset{}
	}

	dst.AbsoluteMachinePosition.UpdateFrom(src.AbsoluteMachinePosition)
	dst.WorkingPosition.UpdateFrom(src.WorkingPosition)
	dst.SavedPositionG28.UpdateFrom(src.SavedPositionG28)
	dst.SavedPositionG30.UpdateFrom(src.SavedPositionG30)
	dst.OffsetG54.UpdateFrom(src.OffsetG54)
	dst.OffsetG55.UpdateFrom(src.OffsetG55)
	dst.OffsetG56.UpdateFrom(src.OffsetG56)
	dst.OffsetG57.UpdateFrom(src.OffsetG57)
	dst.OffsetG58.UpdateFrom(src.OffsetG58)
	dst.OffsetG59.UpdateFrom(src.OffsetG59)
	if src.RxMode != nil {
		dst.RxMode = src.RxMode
	}
}

// TResponse is the central struct. It is used to store
// data received from tinyg after sending a request or command.
type TResponse struct {
	ResponseData   TReceiveObjects `json:"r"`
	ResponseFooter []int           `json:"f"`
}

func (dst *TResponse) UpdateFrom(src *TResponse) {
	dst.ResponseFooter = src.ResponseFooter
	dst.ResponseData.UpdateFrom(&src.ResponseData)
}

func (o *TResponse) Json() (jsonOut []byte) {
	jsonOut, err := jsjson.Marshal(o)
	if err != nil {
		panic(err)
	}
	return
}

// ParseResponse parses raw json data to a TResponse.
func ParseResponse(data []byte) (rsp *TResponse, err error) {
	rsp = &TResponse{}
	err = jsjson.Unmarshal(data, rsp)
	return
}
