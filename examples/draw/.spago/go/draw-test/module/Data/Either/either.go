package either

import (
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Data/Void"
	"iris.local/draw-test/shared"
)

type EitherVoid struct {
	Left  string
	Right void.Void
	Tag   uint8
}

const (
	EitherVoidLeft = iota
	EitherVoidRight
)

type EitherFnUnitAffUnit struct {
	Left  string
	Right func(struct{})
	Tag   uint8
}

const (
	EitherFnUnitAffUnitLeft = iota
	EitherFnUnitAffUnitRight
)

type EitherSeen struct {
	Left  exception.Error
	Right shared.Seen
	Tag   uint8
}

const (
	EitherSeenLeft = iota
	EitherSeenRight
)

type EitherUnit struct {
	Right struct{}
	Left  exception.Error
	Tag   uint8
}

const (
	EitherUnitLeft = iota
	EitherUnitRight
)

type OnComplete struct {
	Handler func(func(EitherUnit))
	Rethrow bool
}

type EitherFnUnit struct {
	Left   string
	Right  func(any) any
	Origin any
	Tag    uint8
}

const (
	EitherFnUnitLeft = iota
	EitherFnUnitRight
)

type EitherString struct {
	Left  exception.Error
	Right string
	Tag   uint8
}

const (
	EitherStringLeft = iota
	EitherStringRight
)

type EitherInt struct {
	Left  string
	Right int
	Tag   uint8
}

const (
	EitherIntLeft = iota
	EitherIntRight
)

func LeftEitherInt(left string) EitherInt {
	return EitherInt{Tag: EitherIntLeft, Left: left}
}

func RightEitherInt(right int) EitherInt {
	return EitherInt{Tag: EitherIntRight, Right: right}
}

type ErrorFailure struct {
	Failed bool
	Left   exception.Error
}

func LeftEitherFnUnitAffUnit(left string) EitherFnUnitAffUnit {
	return EitherFnUnitAffUnit{Tag: EitherFnUnitAffUnitLeft, Left: left}
}

func RightEitherSeen(right shared.Seen) EitherSeen {
	return EitherSeen{Tag: EitherSeenRight, Right: right}
}

func LeftEitherSeen(left exception.Error) EitherSeen {
	return EitherSeen{Tag: EitherSeenLeft, Left: left}
}

func LeftEitherUnit(left exception.Error) EitherUnit {
	return EitherUnit{Tag: EitherUnitLeft, Left: left}
}

func RightEitherUnit(right struct{}) EitherUnit {
	return EitherUnit{Tag: EitherUnitRight, Right: right}
}

func RightEitherString(right string) EitherString {
	return EitherString{Tag: EitherStringRight, Right: right}
}

func LeftEitherString(left exception.Error) EitherString {
	return EitherString{Tag: EitherStringLeft, Left: left}
}

func LeftEitherVoid(left string) *EitherVoid {
	return &EitherVoid{Tag: EitherVoidLeft, Left: left}
}
