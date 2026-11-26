package strom_status

import (
	"iris.local/draw/Data/Tuple"
	chanffi "iris.local/draw/Go_IO_Chan_ffi"
	context "iris.local/draw/Go_IO_Context_ffi"
	"iris.local/draw/Yoga/Connect/Status"
	"iris.local/draw/Yoga/OmIO/Strom"
	"iris.local/draw/shared"
	"iris.local/draw/shared/Yoga.OmIO.Strom.Scope+Yoga.Connect.Status"
)

type EitherFnUnitEffectEitherVariantConnectErrorExceptionStepXY struct {
	Left  status.VariantConnectErrorException
	Right func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure)
	Tag   uint8
}

const (
	EitherFnUnitEffectEitherVariantConnectErrorExceptionStepXYLeft = iota
	EitherFnUnitEffectEitherVariantConnectErrorExceptionStepXYRight
)

type EitherFnUnitEffectEitherVariantConnectErrorExceptionStepBoolean struct {
	Left  status.VariantConnectErrorException
	Right func(struct{}) (strom.StepBoolean, status.VariantConnectErrorExceptionFailure)
	Tag   uint8
}

const (
	EitherFnUnitEffectEitherVariantConnectErrorExceptionStepBooleanLeft = iota
	EitherFnUnitEffectEitherVariantConnectErrorExceptionStepBooleanRight
)

//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func FoldPulledStepFF(c chanffi.Chan[tuple.TupleInt], scope scope_status.Scope, pull func(struct{}) (strom.StepXY, status.VariantConnectErrorExceptionFailure), accumulated struct{}, r struct{}) (struct{}, status.VariantConnectErrorExceptionFailure) {
	for {
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		context2 := scope.Context
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(context2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if boolean {
			return struct{}{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: scope_status.Interrupted}
		}
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		step, err := pull(r)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if err.Failed {
			return struct{}{}, err
		}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
		switch step.Tag {
		case strom.StepXYEnd:
			return accumulated, status.VariantConnectErrorExceptionFailure{}
		default: // strom.StepXYEmit
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
			unit, err2 := FoldPulledStepFFForEachFrom(step.Emit, c, 0, r)
			if err2.Failed {
				return struct{}{}, err2
			}
			accumulated = unit
			continue
		}
	}
}

//line strom_status.go:69
func FoldPulledStepFFForEachFrom(chunk []shared.XY, c chanffi.Chan[tuple.TupleInt], index int, r struct{}) (struct{}, status.VariantConnectErrorExceptionFailure) {
	for {
		if index < len(chunk) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			scrutinee := chunk[index]
//line ../../../../../../src/Draw/Client.purs:50:103
			chanffi.SendImpl(c, tuple.TupleInt{First: scrutinee.X, Second: scrutinee.Y})
			index += 1
			continue
		}
		return struct{}{}, status.VariantConnectErrorExceptionFailure{}
	}
}
