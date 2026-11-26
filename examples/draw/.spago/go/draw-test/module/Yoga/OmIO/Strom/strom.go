package strom

import (
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Tuple"
	"iris.local/draw-test/Data/Variant"
	"iris.local/draw-test/shared"
)

type StepXY struct {
	Emit []shared.XY
	Tag  uint8
}

const (
	StepXYEmit = iota
	StepXYEnd
)

type EitherStepXY struct {
	Left  variant.Exception
	Right StepXY
	Tag   uint8
}

const (
	EitherStepXYLeft = iota
	EitherStepXYRight
)

type HaltStrategy uint8

const (
	HaltStrategyHaltBoth HaltStrategy = iota
	HaltStrategyHaltEither
	HaltStrategyHaltLeft
	HaltStrategyHaltRight
)

type StepUnit struct {
	Emit []struct{}
	Tag  uint8
}

const (
	StepUnitEmit = iota
	StepUnitEnd
)

type EitherStepUnit struct {
	Left  variant.Exception
	Right StepUnit
	Tag   uint8
}

const (
	EitherStepUnitLeft = iota
	EitherStepUnitRight
)

type StepTupleIntInt struct {
	Emit []tuple.TupleInt
	Tag  uint8
}

const (
	StepTupleIntIntEmit = iota
	StepTupleIntIntEnd
)

type EitherStepTupleIntInt struct {
	Left  variant.Exception
	Right StepTupleIntInt
	Tag   uint8
}

const (
	EitherStepTupleIntIntLeft = iota
	EitherStepTupleIntIntRight
)

type StepBoolean struct {
	Emit []bool
	Tag  uint8
}

const (
	StepBooleanEmit = iota
	StepBooleanEnd
)

type Step struct {
	Emit []any
	Tag  uint8
}

const (
	StepEmit = iota
	StepEnd
)

type Signal struct {
	Field1      []shared.XY
	FailedWith1 variant.Exception
	Field0      int
	FailedWith0 int
	Tag         uint8
}

const (
	SignalPulled = iota
	SignalFinished
	SignalFailedWith
)

type MaybeSignalNoRequestXY struct {
	Just Signal
	Tag  uint8
}

func JustSignalNoRequestXY(just Signal) MaybeSignalNoRequestXY {
	return MaybeSignalNoRequestXY{Tag: maybe.TagJust, Just: just}
}

var NoSignalNoRequestXY = MaybeSignalNoRequestXY{Tag: maybe.TagNothing}

func EmitStepXY(emit []shared.XY) StepXY {
	return StepXY{Tag: StepXYEmit, Emit: emit}
}

func EmitStepTupleIntInt(emit []tuple.TupleInt) StepTupleIntInt {
	return StepTupleIntInt{Tag: StepTupleIntIntEmit, Emit: emit}
}

func EmitStepUnit(emit []struct{}) StepUnit {
	return StepUnit{Tag: StepUnitEmit, Emit: emit}
}

func Finished(field0 int) Signal {
	return Signal{Tag: SignalFinished, Field0: field0}
}

func FailedWith(failedWith0 int, failedWith1 variant.Exception) Signal {
	return Signal{Tag: SignalFailedWith, FailedWith0: failedWith0, FailedWith1: failedWith1}
}

func EmitStepBoolean(emit []bool) StepBoolean {
	return StepBoolean{Tag: StepBooleanEmit, Emit: emit}
}

func Pulled(field0 int, field1 []shared.XY) Signal {
	return Signal{Tag: SignalPulled, Field0: field0, Field1: field1}
}

//line ../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1032:5
func OpenClosingAtEndStreamGo(chunk []bool, i int) maybe.Int {
	for {
//line ../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1032:5
		if i < len(chunk) {
//line ../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1036:17
			if chunk[i] {
				i += 1
				continue
			}
//line ../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1036:45
			return maybe.JustInt(i)
		}
		return maybe.NoInt
	}
}
