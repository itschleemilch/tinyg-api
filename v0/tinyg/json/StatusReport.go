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

type TStatusReport struct {
	GCodeLineNo      *int               `json:"line"`
	Velocity         *float64           `json:"vel"`
	FeedRate         *float64           `json:"feed"`
	MachineState     *TMachineState     `json:"stat"`
	UnitsMode        *TUnitsMode        `json:"unit"`
	CoordinateSystem *TCoordinateSystem `json:"coor"`
	MotionMode       *TMotionMode       `json:"momo"`
	PlaneSelect      *TPlaneSelect      `json:"plan"`
	PathMode         *TPathMode         `json:"path"`
	DistanceMode     *TDistanceMode     `json:"dist"`
	FeedRateMode     *TFeedRateMode     `json:"frmo"`
	ArcDistanceMode  *int               `json:"admo"`
	WorkingPositionX *float64           `json:"posx"`
	WorkingPositionY *float64           `json:"posy"`
	WorkingPositionZ *float64           `json:"posz"`
	/* not used: WorkingPositionA *float64 `json:"posa"` */
}

func (dst *TStatusReport) UpdateFrom(src *TStatusReport) {
	if src.GCodeLineNo != nil {
		dst.GCodeLineNo = src.GCodeLineNo
	}
	if src.Velocity != nil {
		dst.Velocity = src.Velocity
	}
	if src.FeedRate != nil {
		dst.FeedRate = src.FeedRate
	}
	if src.MachineState != nil {
		dst.MachineState = src.MachineState
	}
	if src.UnitsMode != nil {
		dst.UnitsMode = src.UnitsMode
	}
	if src.CoordinateSystem != nil {
		dst.CoordinateSystem = src.CoordinateSystem
	}
	if src.MotionMode != nil {
		dst.MotionMode = src.MotionMode
	}
	if src.PlaneSelect != nil {
		dst.PlaneSelect = src.PlaneSelect
	}
	if src.PathMode != nil {
		dst.PathMode = src.PathMode
	}
	if src.DistanceMode != nil {
		dst.DistanceMode = src.DistanceMode
	}
	if src.FeedRateMode != nil {
		dst.FeedRateMode = src.FeedRateMode
	}
	if src.ArcDistanceMode != nil {
		dst.ArcDistanceMode = src.ArcDistanceMode
	}
	if src.WorkingPositionX != nil {
		dst.WorkingPositionX = src.WorkingPositionX
	}
	if src.WorkingPositionY != nil {
		dst.WorkingPositionY = src.WorkingPositionY
	}
	if src.WorkingPositionZ != nil {
		dst.WorkingPositionZ = src.WorkingPositionZ
	}
}
