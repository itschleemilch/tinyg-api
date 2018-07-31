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

func (o *TReceiveObjects) updateOffset(dst *TOffset, src *TOffset) {
	if src != nil {
		if dst != nil {
			dst.UpdateFrom(src)
		} else {
			dst = src
		}
	}
}

func (dst *TReceiveObjects) UpdateFrom(src *TReceiveObjects) {
	if src.FirmwareVersion != nil {
		dst.FirmwareVersion = src.FirmwareVersion
	}
	if src.HardwarePlatform != nil {
		dst.HardwarePlatform = src.HardwarePlatform
	}
	if src.StatusReport != nil {
		if dst.StatusReport != nil {
			dst.StatusReport.UpdateFrom(src.StatusReport)
		} else {
			dst.StatusReport = src.StatusReport
		}
	}
	if src.ExceptionReport != nil {
		dst.ExceptionReport = src.ExceptionReport
	}
	if src.QueueReport != nil {
		dst.QueueReport = src.QueueReport
	}
	if src.RxBufferReport != nil {
		dst.RxBufferReport = src.RxBufferReport
	}
	dst.updateOffset(dst.AbsoluteMachinePosition, src.AbsoluteMachinePosition)
	dst.updateOffset(dst.WorkingPosition, src.WorkingPosition)
	dst.updateOffset(dst.SavedPositionG28, src.SavedPositionG28)
	dst.updateOffset(dst.SavedPositionG30, src.SavedPositionG30)
	dst.updateOffset(dst.OffsetG54, src.OffsetG54)
	dst.updateOffset(dst.OffsetG55, src.OffsetG55)
	dst.updateOffset(dst.OffsetG56, src.OffsetG56)
	dst.updateOffset(dst.OffsetG57, src.OffsetG57)
	dst.updateOffset(dst.OffsetG58, src.OffsetG58)
	dst.updateOffset(dst.OffsetG59, src.OffsetG59)
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

// ParseResponse parses raw json data to a TResponse.
func ParseResponse(data []byte) (rsp *TResponse, err error) {
	rsp = &TResponse{}
	err = jsjson.Unmarshal(data, rsp)
	return
}
