package strom_internal

import (
	"iris.local/draw-test/Data/Map/Internal"
	"iris.local/draw-test/Data/Variant"
	"iris.local/draw-test/Yoga/OmIO/Strom"
	"iris.local/draw-test/shared/Yoga.OmIO.Strom.Scope+Data.Map.Internal"
)

type EitherFnRoomEffectEitherVariantExceptionStepXY struct {
	Left  variant.Exception
	Right func(internal.Room) (strom.StepXY, variant.ExceptionFailure)
	Tag   uint8
}

const (
	EitherFnRoomEffectEitherVariantExceptionStepXYLeft = iota
	EitherFnRoomEffectEitherVariantExceptionStepXYRight
)

type StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY struct {
	Emit []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepXY, variant.ExceptionFailure), variant.ExceptionFailure)
	Tag  uint8
}

const (
	StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXYEmit = iota
	StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXYEnd
)

type EitherStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY struct {
	Left  variant.Exception
	Right StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXY
	Tag   uint8
}

const (
	EitherStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXYLeft = iota
	EitherStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepXYRight
)

type EitherFnRoomEffectEitherVariantExceptionStepTupleIntInt struct {
	Left  variant.Exception
	Right func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure)
	Tag   uint8
}

const (
	EitherFnRoomEffectEitherVariantExceptionStepTupleIntIntLeft = iota
	EitherFnRoomEffectEitherVariantExceptionStepTupleIntIntRight
)

type StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt struct {
	Emit []func(scope_internal.Scope, internal.Room) (func(internal.Room) (strom.StepTupleIntInt, variant.ExceptionFailure), variant.ExceptionFailure)
	Tag  uint8
}

const (
	StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntIntEmit = iota
	StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntIntEnd
)

type EitherStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt struct {
	Left  variant.Exception
	Right StepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntInt
	Tag   uint8
}

const (
	EitherStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntIntLeft = iota
	EitherStepFnScopeFnRoomEffectEitherVariantExceptionFnRoomEffectEitherVariantExceptionStepTupleIntIntRight
)
