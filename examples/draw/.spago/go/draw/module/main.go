package main

import (
	arraypkg "github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/direct"
	"github.com/i-am-the-slime/iris-go/eventloop"
	"github.com/i-am-the-slime/iris-go/exception"
	"github.com/i-am-the-slime/iris-go/prim"
	tea "iris.local/draw/Bubbletea"
	"iris.local/draw/Bubbletea/Hooks"
	"iris.local/draw/Bubbletea/Key"
	"iris.local/draw/Bubbletea/Theme"
	hooksffi "iris.local/draw/Bubbletea_Hooks_ffi"
	bubbletea "iris.local/draw/Bubbletea_ffi"
	"iris.local/draw/Data/Either"
	"iris.local/draw/Data/Map/Internal"
	"iris.local/draw/Data/Maybe"
	"iris.local/draw/Data/Set"
	"iris.local/draw/Data/Tuple"
	"iris.local/draw/Go/IO"
	chanffi "iris.local/draw/Go_IO_Chan_ffi"
	context "iris.local/draw/Go_IO_Context_ffi"
	ioref "iris.local/draw/Go_IO_IORef_ffi"
	internalffi "iris.local/draw/Go_IO_Internal_ffi"
	selectffi "iris.local/draw/Go_IO_Select_ffi"
	ioffi "iris.local/draw/Go_IO_ffi"
	"iris.local/draw/Yoga/Connect/Om/Client"
	om "iris.local/draw/Yoga/Connect/Status"
	"iris.local/draw/Yoga/OmIO/Strom"
	clientffi "iris.local/draw/Yoga_Connect_Client_ffi"
	"iris.local/draw/shared"
	"iris.local/draw/shared/Data.Maybe+Data.Tuple"
	"iris.local/draw/shared/Yoga.OmIO.Strom+Yoga.Connect.Status"
	"iris.local/draw/shared/Yoga.OmIO.Strom.Scope+Yoga.Connect.Status"
)

func main() {
	defer eventloop.Await()
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:109:1
	ioffi.LaunchImpl(func() {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:464:16
		outbox := hooksffi.NewOutbox[bubbletea.Cmd[hooks.ActionUnit]]()
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:464:16
		scope := hooks.NewScope()
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:506:12
		env := hooks.Env{Outbox: outbox, Scope: scope, Theme: theme.CharmtonePantera}
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:506:12
		frame := canvasApp("http://127.0.0.1:50053", hooks.Self{Props: hooks.ActionUnit{Action: func(_ hooks.Env) {
		}}, Emit: hooks.AbsurdSpin}, env)
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:507:11
		init_ := mainDrain(env)
		view := func() bubbletea.ViewType {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:194:30
			return frame.Action(env)
		}
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:545:16
		running := bubbletea.StartProgramImpl(0, 0, tea.Snapshot{Snapshot: struct{}{}}, init_, func(native bubbletea.NativeMsg) bubbletea.Cmd[hooks.ActionUnit] {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:162:21
			kind := bubbletea.KindImpl(native)
			var result hooks.Event
			switch kind {
			case bubbletea.KindKeyPress:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:164:33
				result = hooks.KeyPress(tea.KeyOf(native))
			case bubbletea.KindKeyRelease:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:165:35
				result = hooks.KeyRelease(tea.KeyOf(native))
			case bubbletea.KindMouseClick:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:166:35
				result = hooks.Mouse(tea.MouseOf(tea.MouseActionClick, native))
			case bubbletea.KindMouseRelease:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:167:37
				result = hooks.Mouse(tea.MouseOf(tea.MouseActionRelease, native))
			case bubbletea.KindMouseWheel:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:168:35
				result = hooks.Mouse(tea.MouseOf(tea.MouseActionWheel, native))
			case bubbletea.KindMouseMotion:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:169:36
				result = hooks.Mouse(tea.MouseOf(tea.MouseActionMotion, native))
			case bubbletea.KindResize:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:170:31
				result = hooks.Resize(shared.Size{Width: bubbletea.WidthImpl(native), Height: bubbletea.HeightImpl(native)})
			case bubbletea.KindPaste:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:171:30
				result = hooks.Paste(bubbletea.ContentImpl(native))
			case bubbletea.KindFocus:
				result = hooks.Event{Tag: hooks.EventFocus}
			case bubbletea.KindBlur:
				result = hooks.Event{Tag: hooks.EventBlur}
			case bubbletea.KindClipboard:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:174:34
				result = hooks.Clipboard(bubbletea.ContentImpl(native))
			default:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:162:21
				result = hooks.Native(native)
			}
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:162:21
			event := bubbletea.MessageImpl(result, hooks.Message, native)
			switch event.Tag {
			case hooks.EventMessage:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:194:30
				event.Message.Action(env)
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:429:19
				array := scope.Watchers.Registered()
				for _, action := range array {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:194:30
					action.Action(env)
				}
			default:
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:426:22
				array := scope.Listeners.Registered()
				for _, listener := range array {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:426:80
					scrutinee := listener(event)
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:194:30
					scrutinee.Action(env)
				}
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:429:19
				array2 := scope.Watchers.Registered()
				for _, action := range array2 {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:194:30
					action.Action(env)
				}
			}
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:510:44
			return mainDrain(env)
		}, view)
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:565:39
		either2 := bubbletea.WaitImpl(either.Left, func(snapshot tea.Snapshot) either.Either {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:565:39
			return either.Right(snapshot.Snapshot)
		}, running)
//line ../../../p/either-6.1.0/src/Data/Either.purs:209:1
		switch either2.Tag {
		case either.EitherLeft:
//line ../../../p/either-6.1.0/src/Data/Either.purs:209:23
			io.ThrowIO(exception.NewError(either2.Left))
		default: // either.EitherRight
		}
	})
}

//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:208:32
func canvasApp(address string, ignored hooks.Self, env hooks.Env) hooks.ActionView {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:208:32
	canvas := ioref.New(internal.SharedMapLeaf)
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	mine := chanffi.NewBufferedChanImpl[tuple.TupleInt](256)
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	theirs := chanffi.NewBufferedChanImpl[tuple.TupleInt](1024)
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	env.Outbox.Push(bubbletea.PerformWith(func(context context.Context) hooks.ActionUnit {
//line ../../../../src/Draw/Client.purs:38:43
		drawAlong(address, mine, theirs, context)
		return hooks.ActionUnit{Action: func(_ hooks.Env) {
		}}
	}))
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	env.Outbox.Push(showNextShow(theirs, canvas))
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	env.Scope.Listeners.Register(func(value hooks.Event) hooks.ActionUnit {
		if value.Tag == hooks.EventMouse && value.Mouse.Button == tea.MouseButtonLeftButton {
			if value.Mouse.Action != tea.MouseActionRelease {
				return hooks.ActionUnit{Action: func(_ hooks.Env) {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:149:46
					chanffi.TrySendImpl(mine, tuple.TupleInt{First: value.Mouse.X, Second: value.Mouse.Y})
				}}
			}
		}
		return hooks.ActionUnit{Action: func(_ hooks.Env) {
		}}
	})
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
	env.Scope.Listeners.Register(func(value hooks.Event) hooks.ActionUnit {
		if value.Tag == hooks.EventKeyPress {
			pressed := value.FieldA0
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:140:38
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
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:149:46
		spots := canvas.Load()
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
		array := canvasAppUpTo(spots, tuple.Snd)
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
		mapped := make([]string, len(array))
		for index, y := range array {
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
			array2 := canvasAppUpTo(spots, tuple.Fst)
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
			mapped2 := make([]string, len(array2))
			for index2, x := range array2 {
//line ../../../../src/Draw/Client.purs:60:97
				var result string
//line ../../../../src/Draw/Client.purs:60:97
				if canvasAppGo(y, x, spots) {
					result = "█"
				} else {
					result = " "
				}
//line ../../../../src/Draw/Client.purs:60:97
				mapped2[index2] = result
			}
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
			mapped[index] = prim.StringJoinWith("", mapped2)
		}
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
		return bubbletea.AltScreen(bubbletea.MouseModeImpl(1, bubbletea.ViewImpl(prim.StringJoinWith("\n", mapped))))
	}}
}

//line ../../../../src/Draw/Client.purs:49:17
func drawAlong(address string, mine chanffi.Chan[tuple.TupleInt], theirs chanffi.Chan[tuple.TupleInt], outside context.Context) (struct{}, om.VariantConnectErrorExceptionFailure) {
//line ../../../../src/Draw/Client.purs:49:17
	connection := internalffi.LiftEffectImpl(func() clientffi.Connection {
//line ../../../../src/Draw/Client.purs:49:17
		return clientffi.ConnectImpl(address)
	})
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:91:12
	scope := scope_status.NewScope(outside)
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
	either3, err2 := internalffi.CatchResultsImpl(func(x exception.Error) (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
		return om.EitherVariantConnectErrorExceptionUnit{}, either.ErrorFailure{Failed: true, Left: x}
	}, func() (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
//line ../../../p/transformers-6.1.0/src/Control/Monad/Reader/Trans.purs:70:5
		f, err := client.Client(connection).Draw(func(scope2 scope_status.Scope, _ struct{}) (func(struct{}) (strom.StepXY, om.VariantConnectErrorExceptionFailure), om.VariantConnectErrorExceptionFailure) {
			x := func() maybe_tuple.MaybeMaybeTupleIntInt {
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom.purs:225:15
				return selectffi.Select2Impl(selectffi.OnReceiveImpl(chanffi.ReceiveEndImpl(mine), func(value tuple.TupleInt, open bool) maybe_tuple.MaybeMaybeTupleIntInt {
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:49
					var result maybe_tuple.MaybeTupleIntInt
					if open {
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:62
						result = maybe_tuple.JustTupleIntInt(value)
					} else {
						result = maybe_tuple.NoTupleIntInt
					}
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/Select.purs:43:49
					return maybe_tuple.JustMaybeTupleIntInt(result)
				}), selectffi.OnReceiveImpl(context.DoneImpl(scope2.Context), func(_ struct{}, _ bool) maybe_tuple.MaybeMaybeTupleIntInt {
					return maybe_tuple.NoMaybeTupleIntInt
				}))
			}
			return func(_ struct{}) (strom.StepXY, om.VariantConnectErrorExceptionFailure) {
//line ../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
				maybe2 := x()
				if maybe2.Tag == maybe.TagNothing {
					return strom.StepXY{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: scope_status.Interrupted}
				}
				if maybe2.Tag == maybe.TagJust && maybe2.Just.Tag == maybe.TagNothing {
					return strom.StepXY{Tag: strom.StepXYEnd}, om.VariantConnectErrorExceptionFailure{}
				}
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
				array := []tuple.TupleInt{maybe2.Just.Just}
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
				mapped := make([]shared.XY, len(array))
				for index := range array {
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
					mapped[index] = shared.XY{X: array[index].First, Y: array[index].Second}
				}
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
				return strom.EmitStepXY(mapped), om.VariantConnectErrorExceptionFailure{}
			}, om.VariantConnectErrorExceptionFailure{}
		})(scope, struct{}{})
//line ../../../p/either-6.1.0/src/Data/Either.purs:209:1
		var either2 om.EitherVariantConnectErrorExceptionUnit
		if err.Failed {
//line ../../../p/either-6.1.0/src/Data/Either.purs:209:1
			either2 = om.LeftEitherVariantConnectErrorExceptionUnit(err.Left)
		} else {
//line ../../../p/either-6.1.0/src/Data/Either.purs:210:24
			either2 = om.EitherVariantConnectErrorExceptionUnitFromResults(strom_status.FoldPulledStepFF(theirs, scope, f, struct{}{}, struct{}{}))
		}
		return either2, either.ErrorFailure{}
	})
	var field om.EitherVariantConnectErrorExceptionUnit
	if err2.Failed {
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
		field = om.LeftEitherVariantConnectErrorExceptionUnit(om.Exception(err2.Left))
	} else {
		field = either3
	}
//line ../../../p/either-6.1.0/src/Data/Either.purs:209:1
	switch field.Tag {
	case om.EitherVariantConnectErrorExceptionUnitRight:
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
		either4, err := internalffi.CatchResultsImpl(func(x exception.Error) (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
			return om.EitherVariantConnectErrorExceptionUnit{}, either.ErrorFailure{Failed: true, Left: x}
		}, func() (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:22
			either2 := om.EitherVariantConnectErrorExceptionUnitFromResults(scope_status.CloseScope(scope_status.Exit{Tag: scope_status.ExitEnded}, scope, struct{}{}))
			return either2, either.ErrorFailure{}
		})
		switch field.Tag {
		case om.EitherVariantConnectErrorExceptionUnitLeft:
			return struct{}{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: field.Left}
		default: // om.EitherVariantConnectErrorExceptionUnitRight
			var result om.EitherVariantConnectErrorExceptionUnit
			if err.Failed {
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
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
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:184:23
		boolean := context.IsCancelledImpl(scope.Context)
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
		either4, err := internalffi.CatchResultsImpl(func(x exception.Error) (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
			return om.EitherVariantConnectErrorExceptionUnit{}, either.ErrorFailure{Failed: true, Left: x}
		}, func() (om.EitherVariantConnectErrorExceptionUnit, either.ErrorFailure) {
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			var result bool
			switch failure.Tag {
			case om.VariantConnectErrorExceptionConnectError:
				result = false
			default: // om.VariantConnectErrorExceptionException
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:215:70
				result = exception.Message(failure.Exception) == "the stream was interrupted"
			}
			var result2 scope_status.Exit
			if boolean || result {
				result2 = scope_status.Exit{Tag: scope_status.ExitInterrupted}
			} else {
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
				result2 = scope_status.Failed(failure)
			}
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:105:22
			either2 := om.EitherVariantConnectErrorExceptionUnitFromResults(scope_status.CloseScope(result2, scope, struct{}{}))
			return either2, either.ErrorFailure{}
		})
		switch field.Tag {
		case om.EitherVariantConnectErrorExceptionUnitLeft:
			return struct{}{}, om.VariantConnectErrorExceptionFailure{Failed: true, Left: field.Left}
		default: // om.EitherVariantConnectErrorExceptionUnitRight
			var result om.EitherVariantConnectErrorExceptionUnit
			if err.Failed {
//line ../../../p/yoga-strom-io/098ab498b3d32cd8ff19c1fcf1292bad04856cb3/packages/yoga-strom-io/src/Yoga/OmIO/Strom/Scope.purs:222:14
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

//line ../../../../src/Draw/Client.purs:56:24
func showNextShow(theirs chanffi.Chan[tuple.TupleInt], ioRef ioref.IORef[*internal.Map]) bubbletea.Cmd[hooks.ActionUnit] {
//line ../../../../src/Draw/Client.purs:56:24
	action := func() hooks.ActionUnit {
//line ../../../../src/Draw/Client.purs:56:37
		maybe2 := chanffi.ReceiveImpl(maybe_tuple.NoTupleIntInt, maybe_tuple.JustTupleIntInt, chanffi.ReceiveEndImpl(theirs))
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:9
		var result hooks.ActionUnit
		switch maybe2.Tag {
		case maybe.TagNothing:
			result = hooks.ActionUnit{Action: func(_ hooks.Env) {
			}}
		default: // maybe.TagJust
//line ../../../../src/Draw/Client.purs:56:105
			item := showNextShow(theirs, ioRef)
			result = hooks.ActionUnit{Action: func(env hooks.Env) {
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO/IORef.purs:38:22
				m := ioRef.Load()
//line ../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
				ioRef.Store(showNextShowGo(maybe2.Just, m))
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:158:40
				env.Outbox.Push(item)
			}}
		}
		return result
	}
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:361:18
	return bubbletea.PerformWith(func(_ context.Context) hooks.ActionUnit {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:361:36
		return action()
	})
}

//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:525:15
func mainDrain(env hooks.Env) bubbletea.Cmd[hooks.ActionUnit] {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:525:15
	array := env.Outbox.Drain()
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
	return bubbletea.BatchImpl(array)
}

//line ../../../../src/Draw/Client.purs:63:56
func canvasAppUpTo(spots *internal.Map, coordinate func(tuple.TupleInt) int) []int {
//line ../../../../src/Draw/Client.purs:63:56
	array := arraypkg.Map(coordinate, direct.FromFoldable(func(f func(tuple.TupleInt) func(*direct.ConsList[tuple.TupleInt]) *direct.ConsList[tuple.TupleInt]) func(*direct.ConsList[tuple.TupleInt]) func(*internal.Map) *direct.ConsList[tuple.TupleInt] {
		return func(x *direct.ConsList[tuple.TupleInt]) func(*internal.Map) *direct.ConsList[tuple.TupleInt] {
//line ../../../../src/Draw/Client.purs:63:56
			return set.Foldr(f, x)
		}
	}, spots))
//line ../../../p/maybe-6.0.0/src/Data/Maybe.purs:238:1
	if len(array) == 0 {
		return []int{}
	}
//line ../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
	return arraypkg.Range(0, foldLanes(array, array[0]))
}

//line main.go:471
func canvasAppGo(y int, x int, value *internal.Map) bool {
	for {
		if value == internal.SharedMapLeaf {
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

//line ../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
func showNextShowGo(spot tuple.TupleInt, value *internal.Map) *internal.Map {
//line ../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
	if value == internal.SharedMapLeaf {
//line ../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
		return internal.Node(1, 1, spot, struct{}{}, internal.SharedMapLeaf, internal.SharedMapLeaf)
	}
	mk := value.Node2
	ml := value.Node4
	mr := value.Node5
//line ../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
	left0 := spot.First
	left1 := spot.Second
	right0 := mk.First
	right1 := mk.Second
	if left0 < right0 {
//line ../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:472:15
		return internal.UnsafeBalancedNode(mk, struct{}{}, showNextShowGo(spot, ml), mr)
	}
	if left0 == right0 {
		if left1 < right1 {
//line ../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:472:15
			return internal.UnsafeBalancedNode(mk, struct{}{}, showNextShowGo(spot, ml), mr)
		}
		if left1 == right1 {
//line ../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
			return internal.Node(value.Node0, value.Node1, spot, struct{}{}, ml, mr)
		}
		return internal.UnsafeBalancedNode(mk, struct{}{}, ml, showNextShowGo(spot, mr))
	}
//line ../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:473:15
	return internal.UnsafeBalancedNode(mk, struct{}{}, ml, showNextShowGo(spot, mr))
}
