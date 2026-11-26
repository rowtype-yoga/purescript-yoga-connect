package main

import (
	"github.com/i-am-the-slime/iris-go/aff"
	"github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/eventloop"
	"github.com/i-am-the-slime/iris-go/exception"
	"github.com/i-am-the-slime/iris-go/prim"
	tea "iris.local/draw-test/Bubbletea"
	"iris.local/draw-test/Bubbletea/Hooks"
	"iris.local/draw-test/Bubbletea/Theme"
	hooksffi "iris.local/draw-test/Bubbletea_Hooks_ffi"
	bubbletea "iris.local/draw-test/Bubbletea_ffi"
	"iris.local/draw-test/Data/Either"
	"iris.local/draw-test/Data/Map/Internal"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Tuple"
	"iris.local/draw-test/Draw/Client"
	"iris.local/draw-test/Draw/Server"
	"iris.local/draw-test/Effect/Aff"
	internalffi "iris.local/draw-test/Go_IO_Internal_ffi"
	io "iris.local/draw-test/Go_IO_ffi"
	"iris.local/draw-test/Test/Spec/Config"
	"iris.local/draw-test/Test/Spec/Reporter/Console"
	"iris.local/draw-test/Test/Spec/Runner"
	"iris.local/draw-test/Test/Spec/Tree"
	serverffi "iris.local/draw-test/Yoga_Connect_Server_ffi"
	serverffi2 "iris.local/draw-test/Yoga_Go_HTTP_Om_Server_ffi"
	"iris.local/draw-test/shared"
	"iris.local/draw-test/shared/Test.Spec.Runner.Event+Pipes.Internal"
	"strconv"
	"strings"
)

func main() {
	defer eventloop.Await()
//line ../../../p/aff-8.0.0/src/Effect/Aff.purs:229:12
	fiber := aff.NewFiber(func() struct{} {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
		runner.EvalSpecT(config.DefaultConfig, []*event_internal.Proxy{console.ConsoleReporter}, tree.TupleUnit{First: struct{}{}, Second: []*tree.TreeStringFnUnitAffUnit{tree.NodeTreeStringFnUnitAffUnit(either.LeftEitherFnUnitAffUnit("the shared canvas"), array.Append([]*tree.TreeStringFnUnitAffUnit{tree.LeafTreeStringFnUnitAffUnit("shows bob what ada drew, and ada what bob drew", tree.JustItem(tree.Item{IsParallelizable: maybe.NoBoolean, IsFocused: false, Example: func(around_ func(func(struct{}))) {
//line ../../../p/spec-8.1.2/src/Test/Spec.purs:147:31
			around_(func(_ struct{}) {
//line ../../../../specs/src/Test/Main.purs:32:9
				address := startCanvas()
				action := func() shared.Seen {
					ada := clientOn(address)
					bob := clientOn(address)
					drag(ada, []tuple.TupleInt{{First: 2, Second: 1}, {First: 3, Second: 1}, {First: 4, Second: 1}})
					settle()
					drag(bob, []tuple.TupleInt{{First: 0, Second: 3}})
					settle()
//line ../../../../specs/src/Test/Main.purs:81:19
					s := bubbletea.LastFrameImpl(ada)
//line ../../../../specs/src/Test/Main.purs:81:19
					s2 := bubbletea.LastFrameImpl(bob)
//line ../../../../specs/src/Test/Main.purs:42:11
					return shared.Seen{ByAda: bubbletea.StripAnsi(s), ByBob: bubbletea.StripAnsi(s2)}
				}
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:90:1
				seen := aff.MakeAffDirect(func(value either.EitherSeen) bool {
					return value.Tag == either.EitherSeenLeft
				}, func(value either.EitherSeen) exception.Error {
					return value.Left
				}, func(value either.EitherSeen) shared.Seen {
					return value.Right
				}, func(resolve func(either.EitherSeen)) func(exception.Error) {
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:90:36
					io.ToAffImpl(action, func(x shared.Seen) {
//line ../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
						resolve(either.RightEitherSeen(x))
					}, func(x exception.Error) {
//line ../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
						resolve(either.LeftEitherSeen(x))
					})
					return effectaff.NonCanceler
				})
//line ../../../../specs/src/Test/Main.purs:84:20
				elements := strings.Split(seen.ByBob, "\n")
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
				var v1 string
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
				if len(elements) > 1 {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
					v1 = elements[1]
				} else {
					v1 = ""
				}
//line ../../../p/spec-8.1.2/src/Test/Spec/Assertions.purs:50:3
				if v1 != "  ███" {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
					effectaff.ThrowError(exception.NewError(prim.ShowString(v1) + " ≠ " + prim.ShowString("  ███")))
				}
//line ../../../../specs/src/Test/Main.purs:84:20
				elements2 := strings.Split(seen.ByBob, "\n")
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
				var v12 string
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
				if len(elements2) > 3 {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
					v12 = elements2[3]
				} else {
					v12 = ""
				}
//line ../../../p/spec-8.1.2/src/Test/Spec/Assertions.purs:50:3
				if v12 != "█    " {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
					effectaff.ThrowError(exception.NewError(prim.ShowString(v12) + " ≠ " + prim.ShowString("█    ")))
				}
//line ../../../../specs/src/Test/Main.purs:84:20
				elements3 := strings.Split(seen.ByAda, "\n")
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
				var v13 string
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
				if len(elements3) > 3 {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
					v13 = elements3[3]
				} else {
					v13 = ""
				}
//line ../../../p/spec-8.1.2/src/Test/Spec/Assertions.purs:50:3
				if v13 != "█    " {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
					effectaff.ThrowError(exception.NewError(prim.ShowString(v13) + " ≠ " + prim.ShowString("█    ")))
				}
			})
		}}))}, []*tree.TreeStringFnUnitAffUnit{tree.LeafTreeStringFnUnitAffUnit("shows a client that hasn't painted yet what's on the canvas", tree.JustItem(tree.Item{IsParallelizable: maybe.NoBoolean, IsFocused: false, Example: func(around_ func(func(struct{}))) {
//line ../../../p/spec-8.1.2/src/Test/Spec.purs:147:31
			around_(func(_ struct{}) {
//line ../../../../specs/src/Test/Main.purs:48:9
				address := startCanvas()
				action := func() string {
					ada := clientOn(address)
					drag(ada, []tuple.TupleInt{{First: 1, Second: 0}, {First: 2, Second: 0}})
					settle()
					cleo := clientOn(address)
					settle()
//line ../../../../specs/src/Test/Main.purs:81:19
					s := bubbletea.LastFrameImpl(cleo)
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
					return bubbletea.StripAnsi(s)
				}
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:90:1
				seenByCleo := aff.MakeAffDirect(func(value either.EitherString) bool {
					return value.Tag == either.EitherStringLeft
				}, func(value either.EitherString) exception.Error {
					return value.Left
				}, func(value either.EitherString) string {
					return value.Right
				}, func(resolve func(either.EitherString)) func(exception.Error) {
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:90:36
					io.ToAffImpl(action, func(x string) {
//line ../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
						resolve(either.RightEitherString(x))
					}, func(x exception.Error) {
//line ../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
						resolve(either.LeftEitherString(x))
					})
					return effectaff.NonCanceler
				})
//line ../../../../specs/src/Test/Main.purs:84:20
				elements := strings.Split(seenByCleo, "\n")
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
				var v1 string
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
				if len(elements) > 0 {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:75:20
					v1 = elements[0]
				} else {
					v1 = ""
				}
//line ../../../p/spec-8.1.2/src/Test/Spec/Assertions.purs:50:3
				if v1 != " ██" {
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
					effectaff.ThrowError(exception.NewError(prim.ShowString(v1) + " ≠ " + prim.ShowString(" ██")))
				}
			})
		}}))}))}})
		return struct{}{}
	})
//line ../../../p/aff-8.0.0/src/Effect/Aff.purs:229:12
	_, _, _, _, run := aff.FiberFields(fiber, func(payload exception.Error) either.EitherUnit {
//line ../../../p/aff-8.0.0/src/Effect/Aff.purs:229:12
		return either.LeftEitherUnit(payload)
	}, func(payload struct{}) either.EitherUnit {
//line ../../../p/aff-8.0.0/src/Effect/Aff.purs:229:12
		return either.RightEitherUnit(payload)
	}, either.RightEitherUnit(struct{}{}), func(value either.OnComplete) func(func(either.EitherUnit)) {
		return value.Handler
	})
//line ../../../p/aff-8.0.0/src/Effect/Aff.purs:230:3
	run()
}

//line ../../../../specs/src/Test/Main.purs:60:1
func startCanvas() string {
//line ../../../../specs/src/Test/Main.purs:60:1
	return aff.MakeAffDirect(func(value either.EitherString) bool {
		return value.Tag == either.EitherStringLeft
	}, func(value either.EitherString) exception.Error {
		return value.Left
	}, func(value either.EitherString) string {
		return value.Right
	}, func(resolve func(either.EitherString)) func(exception.Error) {
//line ../../../p/go-io/30228bf274081dcb730b5238c1c00f3fc2363c4c/packages/go-io/src/Go/IO.purs:90:36
		io.ToAffImpl(func() string {
//line ../../../../specs/src/Test/Main.purs:61:3
			room := server.NewRoom()
//line ../../../../src/Draw/Server.purs:31:13
			server2 := internalffi.LiftEffectImpl(func() serverffi.Server {
//line ../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:81:16
				return serverffi.CreateServerImpl()
			})
//line ../../../../src/Draw/Server.purs:32:3
			internalffi.LiftEffectImpl(func() struct{} {
//line ../../../../src/Draw/Server.purs:32:3
				server.Draw(internal.Room{Room: room}, prim.Concat("/draw.v1.Canvas/", prim.StringToUpper(prim.CodeUnitTake(1, "draw")), prim.CodeUnitDrop(1, "draw")), server2)
				return struct{}{}
			})
//line ../../../../src/Draw/Server.purs:33:3
			scrutinee := shared.ListenOptions{Host: "127.0.0.1", Port: 0}
			host := scrutinee.Host
//line ../../../p/yoga-connect-go/9008eae956dc0dfcad5be9f47da59248dc854086/packages/yoga-connect-go/src/Yoga/Connect/Server.purs:112:16
			bound2 := serverffi.ListenImpl(either.LeftEitherInt, either.RightEitherInt, host, scrutinee.Port, server2)
			var bound serverffi2.BoundAddress
			switch bound2.Tag {
			case either.EitherIntLeft:
//line ../../../p/either-6.1.0/src/Data/Either.purs:209:23
				panic(exception.NewError(bound2.Left))
			default: // either.EitherIntRight
				bound = serverffi2.BoundAddress{Address: host, Port: bound2.Right}
			}
//line ../../../../specs/src/Test/Main.purs:63:3
			return "http://127.0.0.1:" + strconv.Itoa(bound.Port)
		}, func(x string) {
//line ../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
			resolve(either.RightEitherString(x))
		}, func(x exception.Error) {
//line ../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
			resolve(either.LeftEitherString(x))
		})
		return effectaff.NonCanceler
	})
}

//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:464:16
func clientOn(address string) bubbletea.NativeRunning[tea.Snapshot] {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:464:16
	outbox := hooksffi.NewOutbox[bubbletea.Cmd[hooks.ActionUnit]]()
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:464:16
	scope := hooks.NewScope()
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:506:12
	env := hooks.Env{Outbox: outbox, Scope: scope, Theme: theme.CharmtonePantera}
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:506:12
	frame := client.CanvasApp(address, hooks.Self{Props: hooks.ActionUnit{Action: func(_ hooks.Env) {
	}}, Emit: hooks.AbsurdSpin}, env)
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:507:11
	init_ := clientOnDrain(env)
	view := func() bubbletea.ViewType {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:194:30
		return frame.Action(env)
	}
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea.purs:545:16
	return bubbletea.StartProgramImpl(40, 12, tea.Snapshot{Snapshot: struct{}{}}, init_, func(native bubbletea.NativeMsg) bubbletea.Cmd[hooks.ActionUnit] {
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
		return clientOnDrain(env)
	}, view)
}

//line ../../../../specs/src/Test/Main.purs:74:3
func drag(running bubbletea.NativeRunning[tea.Snapshot], spots []tuple.TupleInt) {
//line ../../../../specs/src/Test/Main.purs:74:3
	if len(spots) > 0 {
//line ../../../p/foldable-traversable-6.0.0/src/Data/Foldable.purs:140:24
		dragMouse(running, tea.MouseActionClick, spots[0])
	}
//line ../../../p/effect-4.0.0/src/Effect.purs:32:3
	for _, tuple := range array.Slice(1, len(spots), spots) {
//line ../../../p/effect-4.0.0/src/Effect.purs:32:3
		dragMouse(running, tea.MouseActionMotion, tuple)
	}
//line ../../../p/arrays-7.3.0/src/Data/Array.purs:328:18
	index := len(spots) - 1
//line ../../../p/prelude-6.0.2/src/Data/Function.purs:29:14
	if index >= 0 {
//line ../../../p/foldable-traversable-6.0.0/src/Data/Foldable.purs:140:24
		dragMouse(running, tea.MouseActionRelease, spots[index])
	}
}

//line ../../../../specs/src/Test/Main.purs:88:10
func settle() {
//line ../../../../specs/src/Test/Main.purs:88:10
	io.ThreadDelayImpl(150.0)
}

//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:525:15
func clientOnDrain(env hooks.Env) bubbletea.Cmd[hooks.ActionUnit] {
//line ../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Hooks.purs:525:15
	array := env.Outbox.Drain()
//line ../../../p/prelude-6.0.2/src/Control/Monad.purs:86:3
	return bubbletea.BatchImpl(array)
}

//line ../../../../specs/src/Test/Main.purs:78:27
func dragMouse(running bubbletea.NativeRunning[tea.Snapshot], action tea.MouseAction, tuple tuple.TupleInt) {
//line ../../../../specs/src/Test/Main.purs:78:27
	bubbletea.SendNativeImpl(running, tea.ToNative(tea.Mouse(tea.MouseEvent{X: tuple.First, Y: tuple.Second, Button: tea.MouseButtonLeftButton, Action: action, Ctrl: false, Alt: false, Shift: false})))
}
