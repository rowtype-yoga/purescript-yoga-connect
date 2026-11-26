package either_tuple_error

import (
	"iris.local/draw-test/Data/Tuple"
	"iris.local/draw-test/Yoga/Protobuf/Error"
)

type Either struct {
	Left  protobuferror.DecodeError
	Right tuple.TupleInt
	Tag   uint8
}

const (
	EitherLeft = iota
	EitherRight
)
