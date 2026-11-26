package scope_internal

import (
	"iris.local/draw-test/Data/Map/Internal"
	"iris.local/draw-test/Data/Variant"
	context "iris.local/draw-test/Go_IO_Context_ffi"
	ioref "iris.local/draw-test/Go_IO_IORef_ffi"
	mvar "iris.local/draw-test/Go_IO_MVar_ffi"
	"iris.local/draw-test/Yoga/OmIO/Strom/Scope"
)

type Registered struct {
	Second func(scope.Exit, internal.Room) variant.EitherUnit
	First  int
}

type Finalizers struct {
	Open1 []Registered
	Open0 int
	Tag   uint8
}

const (
	FinalizersOpen = iota
	FinalizersClosed
)

type Scope struct {
	CancelContext    context.CancelContext
	Context          context.Context
	DetachFromParent ioref.IORef[func()]
	Finalizers       ioref.IORef[Finalizers]
}

type EitherScope struct {
	Left  variant.Exception
	Right Scope
	Tag   uint8
}

const (
	EitherScopeLeft = iota
	EitherScopeRight
)

type ScopeStopped struct {
	Scope   Scope
	Stopped mvar.MVar[struct{}]
}

type EitherScopeStopped struct {
	Left  variant.Exception
	Right ScopeStopped
	Tag   uint8
}

const (
	EitherScopeStoppedLeft = iota
	EitherScopeStoppedRight
)

type EitherArrayScopeStopped struct {
	Left  variant.Exception
	Right []ScopeStopped
	Tag   uint8
}

const (
	EitherArrayScopeStoppedLeft = iota
	EitherArrayScopeStoppedRight
)
