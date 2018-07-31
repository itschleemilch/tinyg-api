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

type TResponseStatusCode int

const (
	/* Low level codes	System and comms status	*/
	StatusOk                 TResponseStatusCode = 0
	StatusError              TResponseStatusCode = 1
	StatusEagain             TResponseStatusCode = 2
	StatusNoop               TResponseStatusCode = 3
	StatusComplete           TResponseStatusCode = 4
	StatusTerminate          TResponseStatusCode = 5
	StatusReset              TResponseStatusCode = 6
	StatusEol                TResponseStatusCode = 7
	StatusEof                TResponseStatusCode = 8
	StatusFileNotOpen        TResponseStatusCode = 9
	StatusFileSizeExceeded   TResponseStatusCode = 10
	StatusNoSuchDevice       TResponseStatusCode = 11
	StatusBufferEmpty        TResponseStatusCode = 12
	StatusBufferFull         TResponseStatusCode = 13
	StatusBufferFullFatal    TResponseStatusCode = 14
	StatusInitializing       TResponseStatusCode = 15
	StatusEnteringBootLoader TResponseStatusCode = 16
	StatusFunctionIsStubbed  TResponseStatusCode = 17
	/* 18 - 19	Reserved */
	/* Internal System Errors */
	StatusInternalError              TResponseStatusCode = 20
	StatusInternalRangeError         TResponseStatusCode = 21
	StatusFloatingPointError         TResponseStatusCode = 22
	StatusDivideByZero               TResponseStatusCode = 23
	StatusInvalidAddress             TResponseStatusCode = 24
	StatusReadOnlyAddress            TResponseStatusCode = 25
	StatusInitFail                   TResponseStatusCode = 26
	StatusAlarmed                    TResponseStatusCode = 27
	StatusFailedToGetPlannerBuffer   TResponseStatusCode = 28
	StatusGenericExceptionReport     TResponseStatusCode = 29
	StatusPrepLineMoveTimeIsInfinite TResponseStatusCode = 30
	StatusPrepLineMoveTimeIsNan      TResponseStatusCode = 31
	StatusFloatIsInfinite            TResponseStatusCode = 32
	StatusFloatIsNan                 TResponseStatusCode = 33
	StatusPersistenceError           TResponseStatusCode = 34
	StatusBadStatusReportSetting     TResponseStatusCode = 35
	/* 36 – 89	Reserved */
	/* Assertion Failures	Build down from 99 until they meet system errors */
	StatusConfigAssertionFailure     TResponseStatusCode = 90
	StatusXioAssertionFailure        TResponseStatusCode = 91
	StatusEncoderAssertionFailure    TResponseStatusCode = 92
	StatusStepperAssertionFailure    TResponseStatusCode = 93
	StatusPlannerAssertionFailure    TResponseStatusCode = 94
	StatusCanonicalMachine           TResponseStatusCode = 95
	StatusControllerAssertionFailure TResponseStatusCode = 96
	StatusStackOverflow              TResponseStatusCode = 97
	StatusMemoryFault                TResponseStatusCode = 98
	StatusGenericAssertionFailure    TResponseStatusCode = 99
	/* Application and Data Input Errors */
	/* Generic Data Input Errors */
	StatusUnrecognizedName          TResponseStatusCode = 100
	StatusInvalidOrMalformedCommand TResponseStatusCode = 101
	StatusBadNumberFormat           TResponseStatusCode = 102
	StatusBadUnsupportedType        TResponseStatusCode = 103
	StatusParameterIsReadOnly       TResponseStatusCode = 104
	StatusParameterCannotBeRead     TResponseStatusCode = 105
	StatusCommandNotAccepted        TResponseStatusCode = 106
	StatusInputExceedsMaxLength     TResponseStatusCode = 107
	StatusInputLessThanMinValue     TResponseStatusCode = 108
	StatusInputExceedsMaxValue      TResponseStatusCode = 109
	StatusInputValueRangeError      TResponseStatusCode = 110
	StatusJsonSyntaxError           TResponseStatusCode = 111
	StatusJsonTooManyPairs          TResponseStatusCode = 112
	StatusJsonTooLong               TResponseStatusCode = 113
	/* 114 – 129	Reserved */
	/* Gcode Errors and Warnings	Most are from Nist */
	StatusGcodeGenericInputError    TResponseStatusCode = 130
	StatusGcodeCommandUnsupported   TResponseStatusCode = 131
	StatusMcodeCommandUnsupported   TResponseStatusCode = 132
	StatusGcodeModalGroupViolation  TResponseStatusCode = 133
	StatusGcodeAxisIsMissing        TResponseStatusCode = 134
	StatusGcodeAxisCannotBePresent  TResponseStatusCode = 135
	StatusGcodeAxisIsInvalid        TResponseStatusCode = 136
	StatusGcodeAxisIsNotConfigured  TResponseStatusCode = 137
	StatusGcodeAxisNumberIsMissing  TResponseStatusCode = 138
	StatusGcodeAxisNumberIsInvalid  TResponseStatusCode = 139
	StatusGcodeActivePlaneIsMissing TResponseStatusCode = 140
	StatusGcodeActivePlaneIsInvalid TResponseStatusCode = 141
	StatusGcodeFeedrateNotSpecified TResponseStatusCode = 142
	StatusGcodeInverseTimeMode      TResponseStatusCode = 143
	StatusGcodeRotaryAxis           TResponseStatusCode = 144
	StatusGcodeG53WithoutG0OrG1     TResponseStatusCode = 145
	StatusRequestedVelocity         TResponseStatusCode = 146
	StatusCutterCompensation        TResponseStatusCode = 147
	StatusProgrammedPoint           TResponseStatusCode = 148
	StatusSpindleSpeedBelowMinimum  TResponseStatusCode = 149
	StatusSpindleSpeedMaxExceeded   TResponseStatusCode = 150
	StatusSWordIsMissing            TResponseStatusCode = 151
	StatusSWordIsInvalid            TResponseStatusCode = 152
	StatusSpindleMustBeOff          TResponseStatusCode = 153
	StatusSpindleMustBeTurning      TResponseStatusCode = 154
	StatusArcSpecificationError     TResponseStatusCode = 155
	StatusArcAxisMissing            TResponseStatusCode = 156
	StatusArcOffsetsMissing         TResponseStatusCode = 157
	StatusArcRadius                 TResponseStatusCode = 158
	StatusArcEndpoint               TResponseStatusCode = 159
	StatusPWordIsMissing            TResponseStatusCode = 160
	StatusPWordIsInvalid            TResponseStatusCode = 161
	StatusPWordIsZero               TResponseStatusCode = 162
	StatusPWordIsNegative           TResponseStatusCode = 163
	StatusPWordIsNotAnInteger       TResponseStatusCode = 164
	StatusPWordIsNotValidToolNumber TResponseStatusCode = 165
	StatusDWordIsMissing            TResponseStatusCode = 166
	StatusDWordIsInvalid            TResponseStatusCode = 167
	StatusEWordIsMissing            TResponseStatusCode = 168
	StatusEWordIsInvalid            TResponseStatusCode = 169
	StatusHWordIsMissing            TResponseStatusCode = 170
	StatusHWordIsInvalid            TResponseStatusCode = 171
	StatusLWordIsMissing            TResponseStatusCode = 172
	StatusLWordIsInvalid            TResponseStatusCode = 173
	StatusQWordIsMissing            TResponseStatusCode = 174
	StatusQWordIsInvalid            TResponseStatusCode = 175
	StatusRWordIsMissing            TResponseStatusCode = 176
	StatusRWordIsInvalid            TResponseStatusCode = 177
	StatusTWordIsMissing            TResponseStatusCode = 178
	StatusTWordIsInvalid            TResponseStatusCode = 179
	/* 180 - 199	Reserved	reserved for Gcode errors */
	/* TinyG Errors and Warnings */
	StatusGenericError            TResponseStatusCode = 200
	StatusMinimumLengthMove       TResponseStatusCode = 201
	StatusMinimumTimeMove         TResponseStatusCode = 202
	StatusMachineAlarmed          TResponseStatusCode = 203
	StatusLimitSwitchHit          TResponseStatusCode = 204
	StatusPlannerFailedToConverge TResponseStatusCode = 205
	/* 206 - 219	Reserved */
	StatusSoftLimitExceeded     TResponseStatusCode = 220
	StatusSoftLimitExceededXmin TResponseStatusCode = 221
	StatusSoftLimitExceededXmax TResponseStatusCode = 222
	StatusSoftLimitExceededYmin TResponseStatusCode = 223
	StatusSoftLimitExceededYmax TResponseStatusCode = 224
	StatusSoftLimitExceededZmin TResponseStatusCode = 225
	StatusSoftLimitExceededZmax TResponseStatusCode = 226
	StatusSoftLimitExceededAmin TResponseStatusCode = 227
	StatusSoftLimitExceededAmax TResponseStatusCode = 228
	StatusSoftLimitExceededBmin TResponseStatusCode = 229
	StatusSoftLimitExceededBmax TResponseStatusCode = 230
	StatusSoftLimitExceededCmin TResponseStatusCode = 231
	StatusSoftLimitExceededCmax TResponseStatusCode = 232
	/* 233 – 239	Reserved */
	StatusHomingCycleFailed                 TResponseStatusCode = 240
	StatusHomingErrorBadOrNoAxis            TResponseStatusCode = 241
	StatusHomingErrorSwitchMisconfiguration TResponseStatusCode = 242
	StatusHomingErrorZeroSearchVelocity     TResponseStatusCode = 243
	StatusHomingErrorZeroLatchVelocity      TResponseStatusCode = 244
	StatusHomingErrorTravelMinMaxIdentical  TResponseStatusCode = 245
	StatusHomingErrorNegativeLatchBackoff   TResponseStatusCode = 246
	StatusHomingErrorSearchFailed           TResponseStatusCode = 247
	/* StatusReserved TResponseStatusCode = 248 */
	/* StatusReserved TResponseStatusCode = 249 */
	StatusProbeCycleFailed   TResponseStatusCode = 250
	StatusProbeEndpoint      TResponseStatusCode = 251
	StatusJoggingCycleFailed TResponseStatusCode = 252
)
