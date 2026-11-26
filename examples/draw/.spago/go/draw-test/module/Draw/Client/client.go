package client

import (
	arraypkg "github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/direct"
	"github.com/i-am-the-slime/iris-go/exception"
	"github.com/i-am-the-slime/iris-go/prim"
	tea "iris.local/draw-test/Bubbletea"
	"iris.local/draw-test/Bubbletea/Hooks"
	"iris.local/draw-test/Bubbletea/Key"
	bubbletea "iris.local/draw-test/Bubbletea_ffi"
	"iris.local/draw-test/Data/Either"
	"iris.local/draw-test/Data/Map/Internal"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Set"
	"iris.local/draw-test/Data/Tuple"
	chanffi "iris.local/draw-test/Go_IO_Chan_ffi"
	context "iris.local/draw-test/Go_IO_Context_ffi"
	ioref "iris.local/draw-test/Go_IO_IORef_ffi"
	internalffi "iris.local/draw-test/Go_IO_Internal_ffi"
	selectffi "iris.local/draw-test/Go_IO_Select_ffi"
	"iris.local/draw-test/Yoga/Connect/Om/Client"
	om "iris.local/draw-test/Yoga/Connect/Status"
	"iris.local/draw-test/Yoga/OmIO/Strom"
	clientffi "iris.local/draw-test/Yoga_Connect_Client_ffi"
	"iris.local/draw-test/shared"
	"iris.local/draw-test/shared/Data.Maybe+Data.Tuple"
	"iris.local/draw-test/shared/Yoga.OmIO.Strom.Scope+Yoga.Connect.Status"
)

//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:208:32
func CanvasApp(address string, ignored hooks.Self, env hooks.Env) hooks.ActionView {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:208:32
	canvas := ioref.New(internal.SharedMapTupleIntIntLeaf)
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	mine := chanffi.NewBufferedChanImpl[tuple.TupleInt](256)
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	theirs := chanffi.NewBufferedChanImpl[tuple.TupleInt](1024)
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	env.Outbox.Push(bubbletea.PerformWith(func(context context.Context) hooks.ActionUnit {
//line ../../../../../../src/Draw/Client.purs:38:43
		DrawAlong(address, mine, theirs, context)
		return hooks.ActionUnit{Action: func(_ hooks.Env) {
		}}
	}))
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	env.Outbox.Push(ShowNextShow(theirs, canvas))
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	env.Scope.Listeners.Register(func(value hooks.Event) hooks.ActionUnit {
		if value.Tag == hooks.EventMouse && value.Mouse.Button == tea.MouseButtonLeftButton {
			if value.Mouse.Action != tea.MouseActionRelease {
				return hooks.ActionUnit{Action: func(_ hooks.Env) {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:149:46
					chanffi.TrySendImpl(mine, tuple.TupleInt{First: value.Mouse.X, Second: value.Mouse.Y})
				}}
			}
		}
		return hooks.ActionUnit{Action: func(_ hooks.Env) {
		}}
	})
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	env.Scope.Listeners.Register(func(value hooks.Event) hooks.ActionUnit {
		if value.Tag == hooks.EventKeyPress {
			pressed := value.FieldA0
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
			array := []key.KeyCode{key.Char('q')}
			found := false
			for _, code := range array {
				var result bool
				if pressed.Code.Tag == key.KeyCodeChar && code.Tag == key.KeyCodeChar {
					result = pressed.Code.Char == code.Char
				} else if pressed.Code.Tag == key.KeyCodeEnter && code.Tag == key.KeyCodeEnter {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeTab && code.Tag == key.KeyCodeTab {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeEsc && code.Tag == key.KeyCodeEsc {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeSpace && code.Tag == key.KeyCodeSpace {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeBackspace && code.Tag == key.KeyCodeBackspace {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeDelete && code.Tag == key.KeyCodeDelete {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeUp && code.Tag == key.KeyCodeUp {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeDown && code.Tag == key.KeyCodeDown {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeLeft && code.Tag == key.KeyCodeLeft {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeRight && code.Tag == key.KeyCodeRight {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeHome && code.Tag == key.KeyCodeHome {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeEnd && code.Tag == key.KeyCodeEnd {
					result = true
				} else if pressed.Code.Tag == key.KeyCodePgUp && code.Tag == key.KeyCodePgUp {
					result = true
				} else if pressed.Code.Tag == key.KeyCodePgDown && code.Tag == key.KeyCodePgDown {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeInsert && code.Tag == key.KeyCodeInsert {
					result = true
				} else if pressed.Code.Tag == key.KeyCodeF && code.Tag == key.KeyCodeF || pressed.Code.Tag == key.KeyCodeOther && code.Tag == key.KeyCodeOther {
					result = pressed.Code.Field0 == code.Field0
				} else {
					result = false
				}
				if result && (!pressed.Ctrl && !pressed.Alt) {
					found = true
					break
				}
			}
			if found {
				return hooks.Quit
			}
		}
		return hooks.ActionUnit{Action: func(_ hooks.Env) {
		}}
	})
	return hooks.ActionView{Action: func(_ hooks.Env) bubbletea.ViewType {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:149:46
		spots := canvas.Load()
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
		array := CanvasAppUpTo(spots, tuple.Snd)
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
		mapped := make([]string, len(array))
		for index, y := range array {
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
			array2 := CanvasAppUpTo(spots, tuple.Fst)
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
			mapped2 := make([]string, len(array2))
			for index2, x := range array2 {
//line ../../../../../../src/Draw/Client.purs:60:97
				var result string
//line ../../../../../../src/Draw/Client.purs:60:97
				if CanvasAppGo(y, x, spots) {
					result = "█"
				} else {
					result = " "
				}
//line ../../../../../../src/Draw/Client.purs:60:97
				mapped2[index2] = result
			}
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
			mapped[index] = prim.StringJoinWith("", mapped2)
		}
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
		return bubbletea.AltScreen(bubbletea.MouseModeImpl(1, bubbletea.ViewImpl(prim.StringJoinWith("\n", mapped))))
	}}
}

//line ../../../../../../src/Draw/Client.purs:49:17
func DrawAlong(address string, mine chanffi.Chan[tuple.TupleInt], theirs chanffi.Chan[tuple.TupleInt], outside context.Context) (struct{}, om.VariantConnectErrorExceptionFailure) {
//line ../../../../../../src/Draw/Client.purs:49:17
	connection := internalffi.LiftEffectImpl(func() clientffi.Connection {
//line ../../../../../../src/Draw/Client.purs:49:17
		return clientffi.ConnectImpl(address)
	})
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:91:12
	scope := scope_status.NewScope(outside)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
	either3, err2 := internalffi.CatchResultsImpl(func(x exception.Error) (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
		return om.EitherVariantConnectErrorExceptionUnit{}, either.ErrorFailure{Failed: true, Left: x}
	}, func() (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		f, err := client.Client(connection).Draw(func(scope2 scope_status.Scope, _ struct{}) (func(struct{}) (strom.StepXY, om.VariantConnectErrorExceptionFailure), om.VariantConnectErrorExceptionFailure) {
			x := func() maybe_tuple.MaybeMaybeTupleIntInt {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:225:15
				return selectffi.Select2Impl(selectffi.OnReceiveImpl(chanffi.ReceiveEndImpl(mine), func(value tuple.TupleInt, open bool) maybe_tuple.MaybeMaybeTupleIntInt {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:49
					var result maybe_tuple.MaybeTupleIntInt
					if open {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:62
						result = maybe_tuple.JustTupleIntInt(value)
					} else {
						result = maybe_tuple.NoTupleIntInt
					}
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:49
					return maybe_tuple.JustMaybeTupleIntInt(result)
				}), selectffi.OnReceiveImpl(context.DoneImpl(scope2.Context), func(_ struct{}, _ bool) maybe_tuple.MaybeMaybeTupleIntInt {
					return maybe_tuple.NoMaybeTupleIntInt
				}))
			}
			return func(_ struct{}) (strom.StepXY, om.VariantConnectErrorExceptionFailure) {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
				maybe2 := x()
				if maybe2.Tag == maybe.TagNothing {
					return strom.StepXY{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: scope_status.Interrupted}
				}
				if maybe2.Tag == maybe.TagJust && maybe2.Just.Tag == maybe.TagNothing {
					return strom.StepXY{Tag: strom.StepXYEnd}, om.VariantConnectErrorExceptionFailure{}
				}
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
				array := []tuple.TupleInt{maybe2.Just.Just}
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
				mapped := make([]shared.XY, len(array))
				for index := range array {
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
					mapped[index] = shared.XY{X: array[index].First, Y: array[index].Second}
				}
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
				return strom.EmitStepXY(mapped), om.VariantConnectErrorExceptionFailure{}
			}, om.VariantConnectErrorExceptionFailure{}
		})(scope, struct{}{})
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		var either2 om.EitherVariantConnectErrorExceptionUnit
		if err.Failed {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
			either2 = om.LeftEitherVariantConnectErrorExceptionUnit(err.Left)
		} else {
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:210:24
			either2 = om.EitherVariantConnectErrorExceptionUnitFromResults(FoldPulledStepFF(theirs, scope, f, struct{}{}, struct{}{}))
		}
		return either2, either.ErrorFailure{}
	})
	var field om.EitherVariantConnectErrorExceptionUnit
	if err2.Failed {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
		field = om.LeftEitherVariantConnectErrorExceptionUnit(om.Exception(err2.Left))
	} else {
		field = either3
	}
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
	switch field.Tag {
	case om.EitherVariantConnectErrorExceptionUnitRight:
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
		either4, err := internalffi.CatchResultsImpl(func(x exception.Error) (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
			return om.EitherVariantConnectErrorExceptionUnit{}, either.ErrorFailure{Failed: true, Left: x}
		}, func() (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:22
			either2 := om.EitherVariantConnectErrorExceptionUnitFromResults(scope_status.CloseScope(scope_status.Exit{Tag: scope_status.ExitEnded}, scope, struct{}{}))
			return either2, either.ErrorFailure{}
		})
		switch field.Tag {
		case om.EitherVariantConnectErrorExceptionUnitLeft:
			return struct{}{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: field.Left}
		default: // om.EitherVariantConnectErrorExceptionUnitRight
			var result om.EitherVariantConnectErrorExceptionUnit
			if err.Failed {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
				result = om.LeftEitherVariantConnectErrorExceptionUnit(om.Exception(err.Left))
			} else {
				result = either4
			}
			switch result.Tag {
			case om.EitherVariantConnectErrorExceptionUnitLeft:
				return struct{}{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: result.Left}
			default: // om.EitherVariantConnectErrorExceptionUnitRight
				return field.Right, om.VariantConnectErrorExceptionFailure{}
			}
		}
	default: // om.EitherVariantConnectErrorExceptionUnitLeft
		failure := field.Left
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(scope.Context)
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
		either4, err := internalffi.CatchResultsImpl(func(x exception.Error) (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
			return om.EitherVariantConnectErrorExceptionUnit{}, either.ErrorFailure{Failed: true, Left: x}
		}, func() (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			var result bool
			switch failure.Tag {
			case om.VariantConnectErrorExceptionConnectError:
				result = false
			default: // om.VariantConnectErrorExceptionException
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
				result = exception.Message(failure.Exception) == "the stream was interrupted"
			}
			var result2 scope_status.Exit
			if boolean || result {
				result2 = scope_status.Exit{Tag: scope_status.ExitInterrupted}
			} else {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
				result2 = scope_status.Failed(failure)
			}
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			either2 := om.EitherVariantConnectErrorExceptionUnitFromResults(scope_status.CloseScope(result2, scope, struct{}{}))
			return either2, either.ErrorFailure{}
		})
		switch field.Tag {
		case om.EitherVariantConnectErrorExceptionUnitLeft:
			return struct{}{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: field.Left}
		default: // om.EitherVariantConnectErrorExceptionUnitRight
			var result om.EitherVariantConnectErrorExceptionUnit
			if err.Failed {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
				result = om.LeftEitherVariantConnectErrorExceptionUnit(om.Exception(err.Left))
			} else {
				result = either4
			}
			switch result.Tag {
			case om.EitherVariantConnectErrorExceptionUnitLeft:
				return struct{}{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: result.Left}
			default: // om.EitherVariantConnectErrorExceptionUnitRight
				return field.Right, om.VariantConnectErrorExceptionFailure{}
			}
		}
	}
}

//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
func FoldPulledStepFF(c chanffi.Chan[tuple.TupleInt], scope scope_status.Scope, pull func(struct{}) (strom.StepXY, om.VariantConnectErrorExceptionFailure), accumulated struct{}, r struct{}) (struct{}, om.VariantConnectErrorExceptionFailure) {
	for {
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		context2 := scope.Context
//line ../../../../../p/yoga-strom-io/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(context2)
//line ../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
		if boolean {
			return struct{}{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: scope_status.Interrupted}
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
			return accumulated, om.VariantConnectErrorExceptionFailure{}
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

//line ../../../../../../src/Draw/Client.purs:56:24
func ShowNextShow(theirs chanffi.Chan[tuple.TupleInt], ioRef ioref.IORef[*internal.MapTupleIntInt]) bubbletea.Cmd[hooks.ActionUnit] {
//line ../../../../../../src/Draw/Client.purs:56:24
	action := func() hooks.ActionUnit {
//line ../../../../../../src/Draw/Client.purs:56:37
		maybe2 := chanffi.ReceiveImpl(maybe_tuple.NoTupleIntInt, maybe_tuple.JustTupleIntInt, chanffi.ReceiveEndImpl(theirs))
//line ../../../../../p/prelude-6.0.2/src/Control/Monad.purs:86:9
		var result hooks.ActionUnit
		switch maybe2.Tag {
		case maybe.TagNothing:
			result = hooks.ActionUnit{Action: func(_ hooks.Env) {
			}}
		default: // maybe.TagJust
//line ../../../../../../src/Draw/Client.purs:56:105
			item := ShowNextShow(theirs, ioRef)
			result = hooks.ActionUnit{Action: func(env hooks.Env) {
//line ../../../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/IORef.purs:38:22
				m := ioRef.Load()
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
				ioRef.Store(set.InsertGo(maybe2.Just, m))
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:158:40
				env.Outbox.Push(item)
			}}
		}
		return result
	}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:361:18
	return bubbletea.PerformWith(func(_ context.Context) hooks.ActionUnit {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:361:36
		return action()
	})
}

//line ../../../../../../src/Draw/Client.purs:63:56
func CanvasAppUpTo(spots *internal.MapTupleIntInt, coordinate func(tuple.TupleInt) int) []int {
//line ../../../../../../src/Draw/Client.purs:63:56
	array := arraypkg.Map(coordinate, direct.FromFoldable(func(f func(tuple.TupleInt) func(*direct.ConsList[tuple.TupleInt]) *direct.ConsList[tuple.TupleInt]) func(*direct.ConsList[tuple.TupleInt]) func(*internal.MapTupleIntInt) *direct.ConsList[tuple.TupleInt] {
		return func(x *direct.ConsList[tuple.TupleInt]) func(*internal.MapTupleIntInt) *direct.ConsList[tuple.TupleInt] {
//line ../../../../../../src/Draw/Client.purs:63:56
			return set.Foldr(f, x)
		}
	}, spots))
//line ../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:238:1
	if len(array) == 0 {
		return []int{}
	}
//line ../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
	return arraypkg.Range(0, foldLanes(array, array[0]))
}

//line client.go:384
func CanvasAppGo(y int, x int, value *internal.MapTupleIntInt) bool {
	for {
		if value == internal.SharedMapTupleIntIntLeaf {
			return false
		}
		mk := value.Node2
		ml := value.Node4
		mr := value.Node5
		right0 := mk.First
		right1 := mk.Second
		if x < right0 {
			value = ml
			continue
		}
		if x == right0 {
			if y < right1 {
				value = ml
				continue
			}
			if y == right1 {
				return true
			}
			value = mr
			continue
		}
		value = mr
	}
}

func FoldPulledStepFFForEachFrom(chunk []shared.XY, c chanffi.Chan[tuple.TupleInt], index int, r struct{}) (struct{}, om.VariantConnectErrorExceptionFailure) {
	for {
		if index < len(chunk) {
//line ../../../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
			scrutinee := chunk[index]
//line ../../../../../../src/Draw/Client.purs:50:103
			chanffi.SendImpl(c, tuple.TupleInt{First: scrutinee.X, Second: scrutinee.Y})
			index += 1
			continue
		}
		return struct{}{}, om.VariantConnectErrorExceptionFailure{}
	}
}
