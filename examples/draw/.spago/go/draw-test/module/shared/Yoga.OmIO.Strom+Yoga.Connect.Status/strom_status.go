package strom_status

import (
	"iris.local/draw-test/Yoga/Connect/Status"
	"iris.local/draw-test/Yoga/OmIO/Strom"
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
