package status

import (
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/shared"
)

type Code uint8

const (
	CodeCanceled Code = iota
	CodeUnknown
	CodeInvalidArgument
	CodeDeadlineExceeded
	CodeNotFound
	CodeAlreadyExists
	CodePermissionDenied
	CodeResourceExhausted
	CodeFailedPrecondition
	CodeAborted
	CodeOutOfRange
	CodeUnimplemented
	CodeInternal
	CodeUnavailable
	CodeDataLoss
	CodeUnauthenticated
)

type ConnectError struct {
	Details []shared.ErrorDetail
	Message string
	Code    Code
}

type VariantConnectErrorException struct {
	ConnectError ConnectError
	Exception    exception.Error
	Tag          uint8
}

const (
	VariantConnectErrorExceptionConnectError = iota
	VariantConnectErrorExceptionException
)

type EitherVariantConnectErrorExceptionUnit struct {
	Right struct{}
	Left  VariantConnectErrorException
	Tag   uint8
}

const (
	EitherVariantConnectErrorExceptionUnitLeft = iota
	EitherVariantConnectErrorExceptionUnitRight
)

type EitherEitherVariantConnectErrorExceptionUnit struct {
	Left  exception.Error
	Right EitherVariantConnectErrorExceptionUnit
	Tag   uint8
}

const (
	EitherEitherVariantConnectErrorExceptionUnitLeft = iota
	EitherEitherVariantConnectErrorExceptionUnitRight
)

type EitherInt struct {
	Left  VariantConnectErrorException
	Right int
	Tag   uint8
}

const (
	EitherIntLeft = iota
	EitherIntRight
)

type MaybeVariantConnectErrorException struct {
	Just VariantConnectErrorException
	Tag  uint8
}

var NoVariantConnectErrorException = MaybeVariantConnectErrorException{Tag: maybe.TagNothing}

type EitherBoolean struct {
	Left  VariantConnectErrorException
	Tag   uint8
	Right bool
}

const (
	EitherBooleanLeft = iota
	EitherBooleanRight
)

type EitherEitherVariantConnectErrorExceptionBoolean struct {
	Left  exception.Error
	Right EitherBoolean
	Tag   uint8
}

const (
	EitherEitherVariantConnectErrorExceptionBooleanLeft = iota
	EitherEitherVariantConnectErrorExceptionBooleanRight
)

func JustVariantConnectErrorException(just VariantConnectErrorException) MaybeVariantConnectErrorException {
	return MaybeVariantConnectErrorException{Tag: maybe.TagJust, Just: just}
}

type EitherConnectError struct {
	Right struct{}
	Left  ConnectError
	Tag   uint8
}

const (
	EitherConnectErrorLeft = iota
	EitherConnectErrorRight
)

type EitherArrayEitherVariantConnectErrorExceptionUnit struct {
	Left  VariantConnectErrorException
	Right []EitherVariantConnectErrorExceptionUnit
	Tag   uint8
}

const (
	EitherArrayEitherVariantConnectErrorExceptionUnitLeft = iota
	EitherArrayEitherVariantConnectErrorExceptionUnitRight
)

type MaybeEitherVariantConnectErrorExceptionUnit struct {
	Just EitherVariantConnectErrorExceptionUnit
	Tag  uint8
}

var NoEitherVariantConnectErrorExceptionUnit = MaybeEitherVariantConnectErrorExceptionUnit{Tag: maybe.TagNothing}

func JustEitherVariantConnectErrorExceptionUnit(just EitherVariantConnectErrorExceptionUnit) MaybeEitherVariantConnectErrorExceptionUnit {
	return MaybeEitherVariantConnectErrorExceptionUnit{Tag: maybe.TagJust, Just: just}
}

type EitherArrayBoolean struct {
	Left  VariantConnectErrorException
	Right []bool
	Tag   uint8
}

const (
	EitherArrayBooleanLeft = iota
	EitherArrayBooleanRight
)

type VariantConnectErrorExceptionFailure struct {
	Failed bool
	Left   VariantConnectErrorException
}

func LeftEitherVariantConnectErrorExceptionUnit(left VariantConnectErrorException) EitherVariantConnectErrorExceptionUnit {
	return EitherVariantConnectErrorExceptionUnit{Tag: EitherVariantConnectErrorExceptionUnitLeft, Left: left}
}

func Exception(exception exception.Error) VariantConnectErrorException {
	return VariantConnectErrorException{Tag: VariantConnectErrorExceptionException, Exception: exception}
}

func RightEitherVariantConnectErrorExceptionUnit(right struct{}) EitherVariantConnectErrorExceptionUnit {
	return EitherVariantConnectErrorExceptionUnit{Tag: EitherVariantConnectErrorExceptionUnitRight, Right: right}
}

func LeftEitherBoolean(left VariantConnectErrorException) EitherBoolean {
	return EitherBoolean{Tag: EitherBooleanLeft, Left: left}
}

func ConnectErrorVariantConnectErrorException(connectError ConnectError) VariantConnectErrorException {
	return VariantConnectErrorException{Tag: VariantConnectErrorExceptionConnectError, ConnectError: connectError}
}

func RightEitherConnectError(right struct{}) EitherConnectError {
	return EitherConnectError{Tag: EitherConnectErrorRight, Right: right}
}

func LeftEitherConnectError(left ConnectError) EitherConnectError {
	return EitherConnectError{Tag: EitherConnectErrorLeft, Left: left}
}

func RightEitherBoolean(right bool) EitherBoolean {
	return EitherBoolean{Tag: EitherBooleanRight, Right: right}
}

func EitherBooleanFromResults(value bool, err VariantConnectErrorExceptionFailure) EitherBoolean {
	if err.Failed {
		return LeftEitherBoolean(err.Left)
	}
	return RightEitherBoolean(value)
}

func EitherVariantConnectErrorExceptionUnitFromResults(value struct{}, err VariantConnectErrorExceptionFailure) EitherVariantConnectErrorExceptionUnit {
	if err.Failed {
		return LeftEitherVariantConnectErrorExceptionUnit(err.Left)
	}
	return RightEitherVariantConnectErrorExceptionUnit(value)
}
