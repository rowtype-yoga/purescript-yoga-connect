package either_error

import (
	"iris.local/draw/Yoga/Protobuf/Error"
	fields "iris.local/draw/Yoga_Protobuf_Fields_ffi"
	"iris.local/draw/shared"
)

type EitherXY struct {
	Left  protobuferror.DecodeError
	Right shared.XY
	Tag   uint8
}

const (
	EitherXYLeft = iota
	EitherXYRight
)

type EitherFields struct {
	Right fields.Fields
	Left  protobuferror.DecodeError
	Tag   uint8
}

const (
	EitherFieldsLeft = iota
	EitherFieldsRight
)

type EitherInt struct {
	Left  protobuferror.DecodeError
	Right int
	Tag   uint8
}

const (
	EitherIntLeft = iota
	EitherIntRight
)

type DecodeErrorFailure struct {
	Failed bool
	Left   protobuferror.DecodeError
}
