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

// TinygCommand is a string constant that can be used to perform
// a specific task.
type TinygCommand string

const (
	CommandEmptyLine                      TinygCommand = ""
	CommandRequestMachineAbsolutePosition TinygCommand = "{mpo:n}"
	CommandRequestWorkingPosition         TinygCommand = "{pos:n}"
	CommandRequestG92Offset               TinygCommand = "{g92:n}"
	CommandRequestG54Offset               TinygCommand = "{g54:n}"
	CommandRequestG55Offset               TinygCommand = "{g55:n}"
	CommandRequestG56Offset               TinygCommand = "{g56:n}"
	CommandRequestG57Offset               TinygCommand = "{g57:n}"
	CommandRequestG58Offset               TinygCommand = "{g58:n}"
	CommandRequestG59Offset               TinygCommand = "{g59:n}"
	CommandRequestPositionG28             TinygCommand = "{g28:n}"
	CommandRequestPositionG30             TinygCommand = "{g30:n}"
	CommandRequestStatus                  TinygCommand = "{sr:n}"
	CommandRequestFirmwareVersion         TinygCommand = "{fv:n}"
	CommandRequestHardwareVersion         TinygCommand = "{hv:n}"
	CommandRequestHardwarePlatform        TinygCommand = "{hp:n}"
	CommandRequestExceptionReport         TinygCommand = "{ex:n}"
	CommandRequestQueueReport             TinygCommand = "{qr:n}"
	CommandRequestRxBuffer                TinygCommand = "{rx:n}"
	CommandRequestRxMode                  TinygCommand = "{rxm:n}"
	CommandSetRxModeLine                  TinygCommand = "{rxm:1}"
	CommandFeedHold                       TinygCommand = "!"
	CommandFeedResume                                  = "~"
	CommandQueueFlush                     TinygCommand = "%"
	CommandFeedHoldQueueFlush             TinygCommand = "!%"
)
