package console

import (
	"github.com/i-am-the-slime/iris-go/alloc"
	arraypkg "github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/exception"
	"github.com/i-am-the-slime/iris-go/prim"
	"github.com/i-am-the-slime/iris-go/spec"
	"iris.local/draw-test/Ansi/Codes"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Ordering"
	"iris.local/draw-test/Test/Spec/Reporter/Base"
	specresult "iris.local/draw-test/Test/Spec/Result"
	"iris.local/draw-test/Test/Spec/Runner/Event"
	"iris.local/draw-test/Test/Spec/Style"
	"iris.local/draw-test/Test/Spec/Summary"
	"iris.local/draw-test/Test/Spec/Tree"
	"iris.local/draw-test/shared/Test.Spec.Runner.Event+Pipes.Internal"
	"strconv"
)

type PrintAction struct {
	Field0 string
	Field1 specresult.Result
	Tag    uint8
}

const (
	PrintActionPrintTest = iota
	PrintActionPrintPending
)

type Proxy struct {
	Request0 struct{}
	Request1 func(*event.Event) Proxy
	Field0   *event.Event
	Field1   func(struct{}) Proxy
	M        func() Proxy
	Tag      uint8
}

const (
	ProxyRequest = iota
	ProxyRespond
	ProxyM
	ProxyPure
)

func Pure(field0 *event.Event) Proxy {
	return Proxy{Tag: ProxyPure, Field0: field0}
}

var SharedMapIterIterLeaf = &base.MapIter{Tag: base.MapIterIterLeaf}

func Request(request0 struct{}, request1 func(*event.Event) Proxy) Proxy {
	return Proxy{Tag: ProxyRequest, Request0: request0, Request1: request1}
}

func PrintTest(field0 string, field1 specresult.Result) PrintAction {
	return PrintAction{Tag: PrintActionPrintTest, Field0: field0, Field1: field1}
}

func PrintPending(field0 string) PrintAction {
	return PrintAction{Tag: PrintActionPrintPending, Field0: field0}
}

func Request2(request0 struct{}, request1 func(*event.Event) *event_internal.Proxy) *event_internal.Proxy {
	return &event_internal.Proxy{Tag: event_internal.ProxyRequest, Request0: request0, Request1: request1}
}

func Respond(respond0 *event.Event, respond1 func(struct{}) *event_internal.Proxy) *event_internal.Proxy {
	return &event_internal.Proxy{Tag: event_internal.ProxyRespond, Respond0: respond0, Respond1: respond1}
}

func M(m func() *event_internal.Proxy) *event_internal.Proxy {
	return &event_internal.Proxy{Tag: event_internal.ProxyM, M: m}
}

var ConsoleReporter = InitializeConsoleReporter()

//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
func InitializeConsoleReporter() *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
	return ConsoleReporterGoGoGoPureGoM(func() base.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return base.Pure(InitialState)
	})
}

//line console.go:91
var InitialState base.State = base.State{RunningItems: base.SharedMapLeaf, LastPrintedSuitePath: tree.NoArrayPathItem}

//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
func Print_(path []tree.PathItem, a PrintAction, s base.State) base.TupleTupleUnitState {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
	var result base.TupleTupleUnitState
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
	if s.LastPrintedSuitePath.Tag == maybe.TagJust && arraypkg.Eq(tree.Eq, s.LastPrintedSuitePath.Just, path) {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
		result = PrintJoin(a, base.TupleUnit{First: struct{}{}, Second: s}, "")
	} else {
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:93:11
		bound := make([]string, 0, len(path))
		for xIndex := range path {
//line ../../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
			switch path[xIndex].Name.Tag {
			case maybe.TagNothing:
			default: // maybe.TagJust
//line ../../../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
				bound = append(bound, path[xIndex].Name.Just)
			}
		}
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:93:11
		str := prim.StringJoinWith(" » ", bound)
		updated := s
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:93:11
		updated.LastPrintedSuitePath = tree.JustArrayPathItem(path)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:93:11
		array := arraypkg.Append(style.Bold, style.Magenta)
		accumulator := codes.SharedListGraphicsParamNil
		var listGraphicsParamCells alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:93:11
		for index := len(array) - 1; index >= 0; index-- {
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:93:11
			graphicsParam := array[index]
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:93:11
			accumulator = listGraphicsParamCells.New(codes.ListGraphicsParam{Cons0: graphicsParam, Cons1: accumulator})
		}
		var result2 string
		if accumulator == codes.SharedListGraphicsParamNil {
			result2 = str
		} else {
			x := accumulator.Cons0
			xs := accumulator.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = prim.AppendString(text, str)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: codes.GraphicsParam{Tag: codes.GraphicsParamReset}, Second: codes.SharedListGraphicsParamNil}))
			result2 = string(text)
		}
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:93:11
		result = PrintJoin(a, base.TupleUnit{First: struct{}{}, Second: updated}, result2+"\n")
	}
	return base.TupleTupleUnitState{First: result.First, Second: result.Second}
}

//line console.go:152
var Await Proxy = Request(struct{}{}, Pure)

func ConsoleReporterOptsPrintFinishedItem(tuple tree.TupleArrayPathItem) func(base.RunningItem) func(base.State) base.TupleTupleUnitState {
	path := tuple.First
	name := tuple.Second
	return func(value base.RunningItem) func(base.State) base.TupleTupleUnitState {
		if value.Tag == base.RunningItemRunningTest && value.RunningTest.Tag == maybe.TagJust {
			a := PrintTest(name, value.RunningTest.Just)
			return func(s base.State) base.TupleTupleUnitState {
				return Print_(path, a, s)
			}
		}
		if value.Tag == base.RunningItemRunningPending {
			a := PrintPending(name)
			return func(s base.State) base.TupleTupleUnitState {
				return Print_(path, a, s)
			}
		}
		return func(s base.State) base.TupleTupleUnitState {
			return base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: s}, Second: ""}
		}
	}
}

//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
func ConsoleReporterGo(x base.State) *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
	return ConsoleReporterGoGo(x, Await)
}

//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func ConsoleReporterGoGo(x base.State, p Proxy) *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyRequest:
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return Request2(p.Request0, func(x2 *event.Event) *event_internal.Proxy {
			return ConsoleReporterGoGo(x, p.Request1(x2))
		})
	case ProxyRespond:
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return Respond(p.Field0, func(x2 struct{}) *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return ConsoleReporterGoGo(x, p.Field1(x2))
		})
	case ProxyM:
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return M(func() *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return ConsoleReporterGoGo(x, proxy)
		})
	default: // ProxyPure
		r := p.Field0
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return ConsoleReporterGoGoGoRespond(r, x, r)
	}
}

//line console.go:213
func ConsoleReporterGoGoGoRespond(r *event.Event, x base.State, b *event.Event) *event_internal.Proxy {
	return Respond(b, func(x2 struct{}) *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
		return ConsoleReporterGoGoGoPure(r, x, x2)
	})
}

//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
func ConsoleReporterGoGoGoPure(r *event.Event, x base.State, r2 struct{}) *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
	graphicsParam := codes.GraphicsParam{Tag: codes.GraphicsParamReset}
	nonEmptyListGraphicsParam := codes.NonEmpty{First: graphicsParam, Second: codes.SharedListGraphicsParamNil}
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
	escapeCode := codes.Graphics(nonEmptyListGraphicsParam)
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
	return ConsoleReporterGoGoGoPureGoM(func() base.Proxy {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
		var result base.TupleTupleUnitState
		if r.Tag == event.EventSuite && r.FieldA0 == event.ExecutionSequential {
			result = base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: x}, Second: ""}
		} else if r.Tag == event.EventSuite && r.FieldA0 == event.ExecutionParallel {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			v := base.RunningSuite(false)
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			result = ConsoleReporterGoGoGoPureModifyRunningItems(func(value *base.Map) *base.Map {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
				return ConsoleReporterGoGoGoPureGo2(r.FieldA1, v, value)
			}, x)
		} else if r.Tag == event.EventSuiteEnd {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			result = ConsoleReporterGoGoGoPureModifyRunningItems(func(value *base.Map) *base.Map {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
				return ConsoleReporterGoGoGoPureGo3(r.FieldB0, value)
			}, x)
		} else if r.Tag == event.EventTest && r.FieldA0 == event.ExecutionSequential {
			result = base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: x}, Second: ""}
		} else if r.Tag == event.EventTest && r.FieldA0 == event.ExecutionParallel {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			v := base.RunningTest(specresult.NoResult)
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			result = ConsoleReporterGoGoGoPureModifyRunningItems(func(value *base.Map) *base.Map {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
				return ConsoleReporterGoGoGoPureGo2(r.FieldA1, v, value)
			}, x)
		} else if r.Tag == event.EventTestEnd {
			loc := r.FieldB0
			res := r.FieldB1
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:12
			scrutinee := ConsoleReporterGoGoGoPureGo4(loc, x.RunningItems)
			var result2 base.TupleTupleUnitState
			if scrutinee.Tag == maybe.TagJust && scrutinee.Just.Tag == base.RunningItemRunningTest {
//line ../../../../../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
				v := base.RunningTest(specresult.JustResult(res))
//line ../../../../../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
				result2 = ConsoleReporterGoGoGoPureModifyRunningItems(func(value *base.Map) *base.Map {
//line ../../../../../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
					return ConsoleReporterGoGoGoPureGo2(loc, v, value)
				}, x)
			} else if scrutinee.Tag == maybe.TagNothing {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:12
				result2 = ConsoleReporterOptsPrintFinishedItem(loc)(base.RunningTest(specresult.JustResult(res)))(x)
			} else {
				result2 = base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: x}, Second: ""}
			}
			result = base.TupleTupleUnitState{First: result2.First, Second: result2.Second}
		} else if r.Tag == event.EventPending {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			var result2 base.TupleTupleUnitState
			if x.RunningItems == base.SharedMapLeaf {
				result2 = base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: x}, Second: ""}
			} else {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
				result2 = ConsoleReporterGoGoGoPureModifyRunningItems(func(value *base.Map) *base.Map {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
					return ConsoleReporterGoGoGoPureGo5(r.FieldB0, value)
				}, x)
			}
			result = base.TupleTupleUnitState{First: result2.First, Second: result2.Second}
		} else {
			result = base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: x}, Second: ""}
		}
		s_ := result.First.Second
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:30
		var result3 base.TupleTupleUnitState
		if r.Tag == event.EventPending {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			var result2 base.TupleTupleUnitState
			if s_.RunningItems == base.SharedMapLeaf {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
				result2 = Print_(r.FieldB0.First, PrintPending(r.FieldB0.Second), s_)
			} else {
				result2 = base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: s_}, Second: ""}
			}
			result3 = base.TupleTupleUnitState{First: result2.First, Second: result2.Second}
		} else if r.Tag == event.EventEnd {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:30
			scrutinee := summary.Summarize(r.End)
			failed := scrutinee.Failed
			passed := scrutinee.Passed
			pending := scrutinee.Pending
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			total := passed + failed
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:60:20
			var result2 string
			switch total {
			case 1:
				result2 = "test"
			default:
				result2 = "tests"
			}
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:60:20
			amount := strconv.Itoa(passed) + "/" + strconv.Itoa(total) + " " + result2 + " passed"
			var color []codes.GraphicsParam
			if failed > 0 {
				color = style.Red
			} else {
				color = style.Dim
			}
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
			accumulator := codes.SharedListGraphicsParamNil
			var listGraphicsParamCells alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
			for index := len(color) - 1; index >= 0; index-- {
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
				graphicsParam2 := color[index]
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
				accumulator = listGraphicsParamCells.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator})
			}
			var result4 string
			if accumulator == codes.SharedListGraphicsParamNil {
				result4 = amount
			} else {
				x2 := accumulator.Cons0
				xs := accumulator.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
				var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
				text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x2, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
				text = prim.AppendString(text, amount)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
				text = codes.AppendEscapeCodeToString(text, escapeCode)
				result4 = string(text)
			}
			field := result4 + "\n"
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
			if pending > 0 {
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
				var result5 string
				switch pending {
				case 1:
					result5 = "test"
				default:
					result5 = "tests"
				}
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
				str := strconv.Itoa(pending) + " " + result5 + " pending"
				accumulator2 := codes.SharedListGraphicsParamNil
				var listGraphicsParamCells2 alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
				for index := len(style.Yellow) - 1; index >= 0; index-- {
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
					graphicsParam2 := style.Yellow[index]
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
					accumulator2 = listGraphicsParamCells2.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator2})
				}
				var result6 string
				if accumulator2 == codes.SharedListGraphicsParamNil {
					result6 = str
				} else {
					x2 := accumulator2.Cons0
					xs := accumulator2.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
					var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
					text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x2, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
					text = prim.AppendString(text, str)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
					text = codes.AppendEscapeCodeToString(text, escapeCode)
					result6 = string(text)
				}
//line ../../../../../../../p/prelude-6.0.2/src/Data/Semigroup.purs:40:3
				result3 = ConsoleReporterGoGoGoPureJoin(field, base.TupleUnit{First: struct{}{}, Second: s_}, result6+"\n")
			} else {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
				result3 = ConsoleReporterGoGoGoPureJoin(field, base.TupleUnit{First: struct{}{}, Second: s_}, "")
			}
		} else {
			result3 = base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: s_}, Second: ""}
		}
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Base.purs:98:19
		spec.Write(prim.Append(result.Second, result3.Second))
		state := result3.First.Second
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Base.purs:98:19
		return base.Pure(state)
	})
}

//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func ConsoleReporterGoGoGoPureGo(p base.Proxy) *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case base.ProxyRequest:
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return Request2(p.Request0, func(x *event.Event) *event_internal.Proxy {
			return ConsoleReporterGoGoGoPureGo(p.Request1(x))
		})
	case base.ProxyRespond:
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return Respond(p.Respond0, func(x struct{}) *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return ConsoleReporterGoGoGoPureGo(p.Respond1(x))
		})
	case base.ProxyM:
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return M(func() *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return ConsoleReporterGoGoGoPureGo(proxy)
		})
	default: // base.ProxyPure
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:51:25
		return ConsoleReporterGo(p.Pure)
	}
}

//line console.go:442
func ConsoleReporterGoGoGoPureGoM(m func() base.Proxy) *event_internal.Proxy {
	return M(func() *event_internal.Proxy {
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
		proxy := m()
//line ../../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
		return ConsoleReporterGoGoGoPureGo(proxy)
	})
}

//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Base.purs:154:28
func ConsoleReporterGoGoGoPureModifyRunningItems(f func(*base.Map) *base.Map, s base.State) base.TupleTupleUnitState {
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Base.purs:154:28
	nextRunningItems := f(s.RunningItems)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Base.purs:154:28
	allFinished := ConsoleReporterGoGoGoPureModifyRunningItemsGo(nextRunningItems)
	updated := s
	var result *base.Map
	if allFinished {
		result = base.SharedMapLeaf
	} else {
		result = nextRunningItems
	}
	updated.RunningItems = result
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:12
	var result2 base.TupleTupleUnitState
	if allFinished {
		function := func(x base.TupleTupleArrayPathItemString) func(func(base.State) base.TupleTupleUnitState) func(base.State) base.TupleTupleUnitState {
//line ../../../../../../../p/tuples-7.0.0/src/Data/Tuple.purs:131:25
			a := ConsoleReporterOptsPrintFinishedItem(x.First)(x.Second)
			return func(b func(base.State) base.TupleTupleUnitState) func(base.State) base.TupleTupleUnitState {
				return func(s2 base.State) base.TupleTupleUnitState {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
					scrutinee := a(s2)
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:7
					scrutinee2 := b(scrutinee.First.Second)
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
					return base.TupleTupleUnitState{First: base.TupleUnit{First: scrutinee2.First.First, Second: scrutinee2.First.Second}, Second: prim.Append(scrutinee.Second, scrutinee2.Second)}
				}
			}
		}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:12
		result2 = arraypkg.Foldr(func(tuple base.TupleTupleArrayPathItemString, f2 func(base.State) base.TupleTupleUnitState) func(base.State) base.TupleTupleUnitState {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:12
			return function(tuple)(f2)
		}, func(s2 base.State) base.TupleTupleUnitState {
			return base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: s2}, Second: ""}
		}, arraypkg.Unfoldr(func(value base.MaybeTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem) bool {
			return value.Tag == maybe.TagNothing
		}, func(value base.MaybeTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem) base.TupleTupleTupleArrayPathItemStringRunningItem {
			return value.Just
		}, func(value base.TupleTupleTupleArrayPathItemStringRunningItem) base.TupleTupleArrayPathItemString {
			return value.First
		}, func(value base.TupleTupleTupleArrayPathItemStringRunningItem) *base.MapIter {
			return value.Second
		}, base.StepUnfoldrGo, base.IterNode(nextRunningItems, SharedMapIterIterLeaf)))(updated)
	} else {
		result2 = base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: updated}, Second: ""}
	}
	return base.TupleTupleUnitState{First: result2.First, Second: result2.Second}
}

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:181:5
func ConsoleReporterGoGoGoPureModifyRunningItemsGo(value *base.Map) bool {
	for {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:181:5
		if value == base.SharedMapLeaf {
			return true
		}
		v := value.Node3
		l := value.Node4
		r := value.Node5
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:184:9
		if ConsoleReporterGoGoGoPureModifyRunningItemsGo(l) {
			var result bool
			switch v.Tag {
			case base.RunningItemRunningPending:
				result = true
			case base.RunningItemRunningTest:
				switch v.RunningTest.Tag {
				case maybe.TagNothing:
					result = false
				default: // maybe.TagJust
					result = true
				}
			default: // base.RunningItemRunningSuite
				result = v.RunningSuite
			}
			if result {
				value = r
				continue
			}
			return false
		}
		return false
	}
}

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
func ConsoleReporterGoGoGoPureGo2(loc tree.TupleArrayPathItem, v base.RunningItem, value *base.Map) *base.Map {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
	if value == base.SharedMapLeaf {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
		return base.Node(1, 1, loc, v, base.SharedMapLeaf, base.SharedMapLeaf)
	}
	mk := value.Node2
	mv := value.Node3
	ml := value.Node4
	mr := value.Node5
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
	scrutinee := tree.Compare(loc, mk)
	switch scrutinee {
	case ordering.OrderingLT:
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:472:15
		return base.UnsafeBalancedNode(mk, mv, ConsoleReporterGoGoGoPureGo2(loc, v, ml), mr)
	case ordering.OrderingGT:
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:473:15
		return base.UnsafeBalancedNode(mk, mv, ml, ConsoleReporterGoGoGoPureGo2(loc, v, mr))
	default: // ordering.OrderingEQ
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
		return base.Node(value.Node0, value.Node1, loc, v, ml, mr)
	}
}

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:524:3
func ConsoleReporterGoGoGoPureGo3(loc tree.TupleArrayPathItem, value *base.Map) *base.Map {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:524:3
	if value == base.SharedMapLeaf {
		return base.SharedMapLeaf
	}
	mk := value.Node2
	mv := value.Node3
	ml := value.Node4
	mr := value.Node5
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:527:7
	scrutinee := tree.Compare(loc, mk)
	switch scrutinee {
	case ordering.OrderingLT:
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:528:15
		return base.UnsafeBalancedNode(mk, mv, ConsoleReporterGoGoGoPureGo3(loc, ml), mr)
	case ordering.OrderingGT:
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:529:15
		return base.UnsafeBalancedNode(mk, mv, ml, ConsoleReporterGoGoGoPureGo3(loc, mr))
	default: // ordering.OrderingEQ
		switch mv.Tag {
		case base.RunningItemRunningSuite:
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:531:11
			return base.Node(value.Node0, value.Node1, mk, base.RunningSuite(true), ml, mr)
		default:
			if ml == base.SharedMapLeaf {
				return mr
			}
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:767:5
			scrutinee2 := base.UnsafeSplitLast(ml.Node2, ml.Node3, ml.Node4, ml.Node5)
			return base.UnsafeBalancedNode(scrutinee2.First, scrutinee2.Second, scrutinee2.Third, mr)
		}
	}
}

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:280:3
func ConsoleReporterGoGoGoPureGo4(loc tree.TupleArrayPathItem, value *base.Map) base.MaybeRunningItem {
	for {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:280:3
		if value == base.SharedMapLeaf {
			return base.NoRunningItem
		}
		mk := value.Node2
		mv := value.Node3
		ml := value.Node4
		mr := value.Node5
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:283:7
		scrutinee := tree.Compare(loc, mk)
		switch scrutinee {
		case ordering.OrderingLT:
			value = ml
			continue
		case ordering.OrderingGT:
			value = mr
			continue
		default: // ordering.OrderingEQ
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:286:15
			return base.JustRunningItem(mv)
		}
	}
}

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
func ConsoleReporterGoGoGoPureGo5(loc tree.TupleArrayPathItem, value *base.Map) *base.Map {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
	runningItem := base.RunningItem{Tag: base.RunningItemRunningPending}
	if value == base.SharedMapLeaf {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
		return base.Node(1, 1, loc, runningItem, base.SharedMapLeaf, base.SharedMapLeaf)
	}
	mk := value.Node2
	mv := value.Node3
	ml := value.Node4
	mr := value.Node5
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
	scrutinee := tree.Compare(loc, mk)
	switch scrutinee {
	case ordering.OrderingLT:
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:472:15
		return base.UnsafeBalancedNode(mk, mv, ConsoleReporterGoGoGoPureGo5(loc, ml), mr)
	case ordering.OrderingGT:
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:473:15
		return base.UnsafeBalancedNode(mk, mv, ml, ConsoleReporterGoGoGoPureGo5(loc, mr))
	default: // ordering.OrderingEQ
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
		return base.Node(value.Node0, value.Node1, loc, runningItem, ml, mr)
	}
}

//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
func ConsoleReporterGoGoGoPureJoin(field string, a base.TupleUnit, w string) base.TupleTupleUnitState {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
	accumulator := codes.SharedListGraphicsParamNil
	var listGraphicsParamCells alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
	for index := len(style.Bold) - 1; index >= 0; index-- {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		graphicsParam := style.Bold[index]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		accumulator = listGraphicsParamCells.New(codes.ListGraphicsParam{Cons0: graphicsParam, Cons1: accumulator})
	}
	var result string
	if accumulator == codes.SharedListGraphicsParamNil {
		result = "Summary"
	} else {
		x := accumulator.Cons0
		xs := accumulator.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
		var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
		text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
		text = append(text, "Summary"...)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
		text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: codes.GraphicsParam{Tag: codes.GraphicsParamReset}, Second: codes.SharedListGraphicsParamNil}))
		result = string(text)
	}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
	return base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: a.Second}, Second: prim.Concat("\n", result, "\n", field, w, "\n")}
}

//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:7
func PrintJoin(a PrintAction, a2 base.TupleUnit, w string) base.TupleTupleUnitState {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:7
	s_ := a2.Second
	graphicsParam := codes.GraphicsParam{Tag: codes.GraphicsParamReset}
	nonEmptyListGraphicsParam := codes.NonEmpty{First: graphicsParam, Second: codes.SharedListGraphicsParamNil}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:72:7
	escapeCode := codes.Graphics(nonEmptyListGraphicsParam)
	if a.Tag == PrintActionPrintTest && a.Field1.Tag == specresult.ResultSuccess {
		name := a.Field0
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		accumulator := codes.SharedListGraphicsParamNil
		var listGraphicsParamCells alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		for index := len(style.Green) - 1; index >= 0; index-- {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
			graphicsParam2 := style.Green[index]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
			accumulator = listGraphicsParamCells.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator})
		}
		var result string
		if accumulator == codes.SharedListGraphicsParamNil {
			result = "✓︎ "
		} else {
			x := accumulator.Cons0
			xs := accumulator.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = append(text, "✓︎ "...)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, escapeCode)
			result = string(text)
		}
		accumulator2 := codes.SharedListGraphicsParamNil
		var listGraphicsParamCells2 alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		for index := len(style.Dim) - 1; index >= 0; index-- {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
			graphicsParam2 := style.Dim[index]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
			accumulator2 = listGraphicsParamCells2.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator2})
		}
		var result2 string
		if accumulator2 == codes.SharedListGraphicsParamNil {
			result2 = name
		} else {
			x := accumulator2.Cons0
			xs := accumulator2.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = prim.AppendString(text, name)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, escapeCode)
			result2 = string(text)
		}
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		return base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: s_}, Second: prim.Concat(w, "  ", result, result2, "\n")}
	}
	if a.Tag == PrintActionPrintTest && a.Field1.Tag == specresult.ResultFailure {
		str := "✗ " + a.Field0 + ":"
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Reporter/Console.purs:101:42
		str2 := exception.Message(a.Field1.Failure)
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		accumulator := codes.SharedListGraphicsParamNil
		var listGraphicsParamCells3 alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		for index := len(style.Red) - 1; index >= 0; index-- {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
			graphicsParam2 := style.Red[index]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
			accumulator = listGraphicsParamCells3.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator})
		}
		var result string
		if accumulator == codes.SharedListGraphicsParamNil {
			result = str
		} else {
			x := accumulator.Cons0
			xs := accumulator.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = prim.AppendString(text, str)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, escapeCode)
			result = string(text)
		}
		accumulator2 := codes.SharedListGraphicsParamNil
		var listGraphicsParamCells4 alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		for index := len(style.Red) - 1; index >= 0; index-- {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
			graphicsParam2 := style.Red[index]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
			accumulator2 = listGraphicsParamCells4.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator2})
		}
		var result2 string
		if accumulator2 == codes.SharedListGraphicsParamNil {
			result2 = str2
		} else {
			x := accumulator2.Cons0
			xs := accumulator2.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = prim.AppendString(text, str2)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, escapeCode)
			result2 = string(text)
		}
		return base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: s_}, Second: w + "  " + result + "\n\n  " + result2 + "\n"}
	}
	str3 := "~ " + a.Field0
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
	accumulator3 := codes.SharedListGraphicsParamNil
	var listGraphicsParamCells5 alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
	for index := len(style.Cyan) - 1; index >= 0; index-- {
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		graphicsParam2 := style.Cyan[index]
//line ../../../../../../../p/transformers-6.1.0/src/Control/Monad/Writer/Trans.purs:74:33
		accumulator3 = listGraphicsParamCells5.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator3})
	}
	var result3 string
	if accumulator3 == codes.SharedListGraphicsParamNil {
		result3 = str3
	} else {
		x := accumulator3.Cons0
		xs := accumulator3.Cons1
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
		var text []byte
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
		text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
		text = prim.AppendString(text, str3)
//line ../../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
		text = codes.AppendEscapeCodeToString(text, escapeCode)
		result3 = string(text)
	}
	return base.TupleTupleUnitState{First: base.TupleUnit{First: struct{}{}, Second: s_}, Second: w + "  " + result3 + "\n"}
}
