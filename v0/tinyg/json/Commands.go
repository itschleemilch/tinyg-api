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

const (
	CommandEmptyLine                      string = ""
	CommandRequestMachineAbsolutePosition string = "{mpo:n}"
	CommandRequestWorkingPosition         string = "{pos:n}"
	CommandRequestG92Offset               string = "{g92:n}"
	CommandRequestG54Offset               string = "{g54:n}"
	CommandRequestG55Offset               string = "{g55:n}"
	CommandRequestG56Offset               string = "{g56:n}"
	CommandRequestG57Offset               string = "{g57:n}"
	CommandRequestG58Offset               string = "{g58:n}"
	CommandRequestG59Offset               string = "{g59:n}"
	CommandRequestPositionG28             string = "{g28:n}"
	CommandRequestPositionG30             string = "{g30:n}"
	CommandRequestStatus                  string = "{sr:n}"
	CommandRequestFirmwareVersion         string = "{fv:n}"
	CommandRequestHardwareVersion         string = "{hv:n}"
	CommandRequestHardwarePlatform        string = "{hp:n}"
	CommandRequestExceptionReport         string = "{ex:n}"
	CommandRequestQueueReport             string = "{qr:n}"
	CommandRequestRxBuffer                string = "{rx:n}"
	CommandRequestRxMode                  string = "{rxm:n}"
	CommandSetRxModeLine                  string = "{rxm:1}"
	CommandFeedHold                       string = "!"
	CommandFeedResume                     string = "~"
	CommandQueueFlush                     string = "%"
	CommandFeedHoldQueueFlush             string = "!%"
	CommandSetFlowControlCts              string = "{ex:2}"
	CommandHardwareReset                  string = "\x18"
)
