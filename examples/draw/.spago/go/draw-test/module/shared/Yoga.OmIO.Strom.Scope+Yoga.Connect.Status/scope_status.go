package scope_status

import (
	"github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Data/Either"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Tuple"
	context "iris.local/draw-test/Go_IO_Context_ffi"
	ioref "iris.local/draw-test/Go_IO_IORef_ffi"
	internal "iris.local/draw-test/Go_IO_Internal_ffi"
	"iris.local/draw-test/Yoga/Connect/Status"
	"iris.local/draw-test/Yoga/OmIO/Strom/Scope"
	"slices"
)

type Exit struct {
	Failed status.VariantConnectErrorException
	Tag    uint8
}

const (
	ExitEnded = iota
	ExitFailed
	ExitInterrupted
)

type Registered struct {
	Second func(Exit, struct{}) status.EitherVariantConnectErrorExceptionUnit
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

type Either struct {
	Left  status.VariantConnectErrorException
	Right Scope
	Tag   uint8
}

const (
	EitherLeft = iota
	EitherRight
)

type Tuple struct {
	First  Finalizers
	Second Finalizers
}

func Failed(failed status.VariantConnectErrorException) Exit {
	return Exit{Tag: ExitFailed, Failed: failed}
}

func Open(open0 int, open1 []Registered) Finalizers {
	return Finalizers{Tag: FinalizersOpen, Open0: open0, Open1: open1}
}

var Interrupted status.VariantConnectErrorException = status.Exception(scope.Interruption)

//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:109:31
func NewScope(parentContext context.Context) Scope {
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:109:31
	tuple2 := context.WithCancelImpl(func(context2 context.Context) func(context.CancelContext) tuple.TupleContext {
		return func(cancelContext context.CancelContext) tuple.TupleContext {
			return tuple.TupleContext{First: context2, Second: cancelContext}
		}
	}, parentContext)
	finalizers := ioref.New(Open(0, []Registered{}))
	detachFromParent := ioref.New(func() {
	})
	return Scope{Context: tuple2.First, CancelContext: tuple2.Second, Finalizers: finalizers, DetachFromParent: detachFromParent}
}

//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:164:12
func CloseScope(exit Exit, scope Scope, r struct{}) (struct{}, status.VariantConnectErrorExceptionFailure) {
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:164:12
	finalizers := ioref.AtomicModifyIORefImpl(Fst, Snd, func(state Finalizers) Tuple {
		return Tuple{First: Finalizers{Tag: FinalizersClosed}, Second: state}
	}, scope.Finalizers)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	switch finalizers.Tag {
	case FinalizersClosed:
		return struct{}{}, status.VariantConnectErrorExceptionFailure{}
	default: // FinalizersOpen
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:168:7
		context.CancelImpl(scope.CancelContext)
		effect := scope.DetachFromParent.Load()
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		effect()
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:171:24
		elements := array.Reverse(finalizers.Open1)
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:171:24
		array_, err := CloseScopeFrom(exit, r, elements, 0, []status.EitherVariantConnectErrorExceptionUnit{})
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if err.Failed {
			return struct{}{}, err
		}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
		found := status.NoEitherVariantConnectErrorExceptionUnit
		for _, either := range array_ {
			var result bool
			switch either.Tag {
			case status.EitherVariantConnectErrorExceptionUnitLeft:
				result = true
			default: // status.EitherVariantConnectErrorExceptionUnitRight
				result = false
			}
			if result {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
				found = status.JustEitherVariantConnectErrorExceptionUnit(either)
				break
			}
		}
		if found.Tag == maybe.TagJust && found.Just.Tag == status.EitherVariantConnectErrorExceptionUnitLeft {
			return struct{}{}, status.VariantConnectErrorExceptionFailure{Failed: true, Left: found.Just.Left}
		}
		return struct{}{}, status.VariantConnectErrorExceptionFailure{}
	}
}

//line scope_status.go:139
func Fst(tuple Tuple) Finalizers {
	return tuple.First
}

func Snd(tuple Tuple) Finalizers {
	return tuple.Second
}

func CloseScopeFrom(exit Exit, r struct{}, elements []Registered, index int, field0 []status.EitherVariantConnectErrorExceptionUnit) ([]status.EitherVariantConnectErrorExceptionUnit, status.VariantConnectErrorExceptionFailure) {
	field0 = slices.Clip(field0)
	for {
		exitCaptured := exit
		rCaptured := r
		if index < len(elements) {
			scrutinee := elements[index]
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
			either3, err := internal.CatchResultsImpl(func(x exception.Error) (status.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
				return status.EitherVariantConnectErrorExceptionUnit{}, either.ErrorFailure{Failed: true, Left: x}
			}, func() (status.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:22
				either2 := scrutinee.Second(exitCaptured, rCaptured)
				return either2, either.ErrorFailure{}
			})
			var result status.EitherVariantConnectErrorExceptionUnit
			if err.Failed {
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
				result = status.LeftEitherVariantConnectErrorExceptionUnit(status.Exception(err.Left))
			} else {
				result = either3
			}
//line ../../../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
			exit, r, index, field0 = exitCaptured, rCaptured, index+1, append(field0, result)
			continue
		}
		return field0, status.VariantConnectErrorExceptionFailure{}
	}
}
