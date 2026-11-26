package variant

import (
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Data/Maybe"
)

type Exception struct {
	Exception exception.Error
}

type EitherUnit struct {
	Right struct{}
	Left  Exception
	Tag   uint8
}

const (
	EitherUnitLeft = iota
	EitherUnitRight
)

type EitherEitherVariantExceptionUnit struct {
	Left  exception.Error
	Right EitherUnit
	Tag   uint8
}

const (
	EitherEitherVariantExceptionUnitLeft = iota
	EitherEitherVariantExceptionUnitRight
)

type EitherInt struct {
	Left  Exception
	Right int
	Tag   uint8
}

const (
	EitherIntLeft = iota
	EitherIntRight
)

type EitherArrayUnit struct {
	Left  Exception
	Right []struct{}
	Tag   uint8
}

const (
	EitherArrayUnitLeft = iota
	EitherArrayUnitRight
)

type EitherBoolean struct {
	Left  Exception
	Tag   uint8
	Right bool
}

const (
	EitherBooleanLeft = iota
	EitherBooleanRight
)

type EitherEitherVariantExceptionBoolean struct {
	Left  exception.Error
	Right EitherBoolean
	Tag   uint8
}

const (
	EitherEitherVariantExceptionBooleanLeft = iota
	EitherEitherVariantExceptionBooleanRight
)

type Variant uint8

type MaybeVariant struct {
	Tag  uint8
	Just Variant
}

var NoVariant = MaybeVariant{Tag: maybe.TagNothing}

type EitherArrayEitherVariantExceptionUnit struct {
	Left  Exception
	Right []EitherUnit
	Tag   uint8
}

const (
	EitherArrayEitherVariantExceptionUnitLeft = iota
	EitherArrayEitherVariantExceptionUnitRight
)

type MaybeEitherVariantExceptionUnit struct {
	Just EitherUnit
	Tag  uint8
}

var NoEitherVariantExceptionUnit = MaybeEitherVariantExceptionUnit{Tag: maybe.TagNothing}

func JustEitherVariantExceptionUnit(just EitherUnit) MaybeEitherVariantExceptionUnit {
	return MaybeEitherVariantExceptionUnit{Tag: maybe.TagJust, Just: just}
}

type ExceptionFailure struct {
	Failed bool
	Left   Exception
}

func RightEitherUnit(right struct{}) EitherUnit {
	return EitherUnit{Tag: EitherUnitRight, Right: right}
}

func LeftEitherUnit(left Exception) EitherUnit {
	return EitherUnit{Tag: EitherUnitLeft, Left: left}
}

func LeftEitherBoolean(left Exception) EitherBoolean {
	return EitherBoolean{Tag: EitherBooleanLeft, Left: left}
}

func RightEitherBoolean(right bool) EitherBoolean {
	return EitherBoolean{Tag: EitherBooleanRight, Right: right}
}

func EitherBooleanFromResults(value bool, err ExceptionFailure) EitherBoolean {
	if err.Failed {
		return LeftEitherBoolean(err.Left)
	}
	return RightEitherBoolean(value)
}

func EitherUnitFromResults(value struct{}, err ExceptionFailure) EitherUnit {
	if err.Failed {
		return LeftEitherUnit(err.Left)
	}
	return RightEitherUnit(value)
}

func EitherUnitResults(either EitherUnit) (struct{}, ExceptionFailure) {
	if either.Tag == EitherUnitLeft {
		return struct{}{}, ExceptionFailure{Failed: true, Left: either.Left}
	}
	return either.Right, ExceptionFailure{}
}
