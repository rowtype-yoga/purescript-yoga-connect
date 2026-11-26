package either_tuple_error

import (
	"iris.local/draw/Data/Tuple"
	"iris.local/draw/Yoga/Protobuf/Error"
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
