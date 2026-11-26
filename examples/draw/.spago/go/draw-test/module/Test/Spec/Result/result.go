package result

import (
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Test/Spec/Speed"
)

type Result struct {
	Failure  exception.Error
	Success1 float64
	Tag      uint8
	Success0 speed.Speed
}

const (
	ResultSuccess = iota
	ResultFailure
)

type MaybeResult struct {
	Just Result
	Tag  uint8
}

var NoResult = MaybeResult{Tag: maybe.TagNothing}

func JustResult(just Result) MaybeResult {
	return MaybeResult{Tag: maybe.TagJust, Just: just}
}

func Failure(failure exception.Error) Result {
	return Result{Tag: ResultFailure, Failure: failure}
}

func Success(success0 speed.Speed, success1 float64) Result {
	return Result{Tag: ResultSuccess, Success0: success0, Success1: success1}
}
