package runner

import (
	"github.com/i-am-the-slime/iris-go/aff"
	"github.com/i-am-the-slime/iris-go/alloc"
	arraypkg "github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/avar"
	"github.com/i-am-the-slime/iris-go/direct"
	"github.com/i-am-the-slime/iris-go/exception"
	"github.com/i-am-the-slime/iris-go/now"
	"github.com/i-am-the-slime/iris-go/prim"
	specpkg "github.com/i-am-the-slime/iris-go/spec"
	"iris.local/draw-test/Ansi/Codes"
	"iris.local/draw-test/Control/Category"
	"iris.local/draw-test/Data/Either"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Tuple"
	"iris.local/draw-test/Effect/Aff"
	"iris.local/draw-test/Pipes/Core"
	"iris.local/draw-test/Pipes/Internal"
	specresult "iris.local/draw-test/Test/Spec/Result"
	"iris.local/draw-test/Test/Spec/Runner/Event"
	"iris.local/draw-test/Test/Spec/Speed"
	"iris.local/draw-test/Test/Spec/Style"
	"iris.local/draw-test/Test/Spec/Summary"
	"iris.local/draw-test/Test/Spec/Tree"
	"iris.local/draw-test/shared/Test.Spec.Runner.Event+Pipes.Internal"
	"iris.local/draw-test/shared/Test.Spec.Tree+Test.Spec.Result"
	"strconv"
)

type EitherArrayTreeStringVoidResult struct {
	Left  exception.Error
	Right []*tree_result.Tree
	Tag   uint8
}

const (
	EitherArrayTreeStringVoidResultLeft = iota
	EitherArrayTreeStringVoidResultRight
)

type ProxyEventArrayTreeStringVoidResult struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyEventArrayTreeStringVoidResult
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyEventArrayTreeStringVoidResult
	M        func() *ProxyEventArrayTreeStringVoidResult
	Pure     []*tree_result.Tree
	Tag      uint8
}

const (
	ProxyEventArrayTreeStringVoidResultRequest = iota
	ProxyEventArrayTreeStringVoidResultRespond
	ProxyEventArrayTreeStringVoidResultM
	ProxyEventArrayTreeStringVoidResultPure
)

type ProxyX struct {
	Field0 internal.X
	Field1 func(struct{}) *ProxyX
	M      func() *ProxyX
	Pure   []*tree_result.Tree
	Tag    uint8
}

const (
	ProxyXRequest = iota
	ProxyXRespond
	ProxyXM
	ProxyXPure
)

type ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult
	M        func() *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult
	Pure     func([]*tree_result.Tree) [][]*tree_result.Tree
	Tag      uint8
}

const (
	ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultRequest = iota
	ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultRespond
	ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultM
	ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultPure
)

type ProxyArrayArrayTreeStringVoidResult struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyArrayArrayTreeStringVoidResult
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyArrayArrayTreeStringVoidResult
	M        func() *ProxyArrayArrayTreeStringVoidResult
	Pure     [][]*tree_result.Tree
	Tag      uint8
}

const (
	ProxyArrayArrayTreeStringVoidResultRequest = iota
	ProxyArrayArrayTreeStringVoidResultRespond
	ProxyArrayArrayTreeStringVoidResultM
	ProxyArrayArrayTreeStringVoidResultPure
)

type Ref[T any] struct {
	Value T
}

type Element struct {
	Test             *tree.TreeTupleArrayPathItemString
	IsParallelizable bool
}

type ProxyBoolean struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyBoolean
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyBoolean
	M        func() *ProxyBoolean
	Tag      uint8
	Pure     bool
}

const (
	ProxyBooleanRequest = iota
	ProxyBooleanRespond
	ProxyBooleanM
	ProxyBooleanPure
)

type ProxyAVarEitherUnitEvent struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyAVarEitherUnitEvent
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyAVarEitherUnitEvent
	M        func() *ProxyAVarEitherUnitEvent
	Pure     avar.AVar[event.EitherUnit]
	Tag      uint8
}

const (
	ProxyAVarEitherUnitEventRequest = iota
	ProxyAVarEitherUnitEventRespond
	ProxyAVarEitherUnitEventM
	ProxyAVarEitherUnitEventPure
)

type ProxyUnit struct {
	Pure     struct{}
	Request0 internal.X
	Request1 func(struct{}) *ProxyUnit
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyUnit
	M        func() *ProxyUnit
	Tag      uint8
}

const (
	ProxyUnitRequest = iota
	ProxyUnitRespond
	ProxyUnitM
	ProxyUnitPure
)

func PureProxyUnit(pure struct{}) *ProxyUnit {
	return &ProxyUnit{Tag: ProxyUnitPure, Pure: pure}
}

type ProxyResult struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyResult
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyResult
	M        func() *ProxyResult
	Pure     specresult.Result
	Tag      uint8
}

const (
	ProxyResultRequest = iota
	ProxyResultRespond
	ProxyResultM
	ProxyResultPure
)

type ProxyFnUnitArrayTreeStringVoidResult struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyFnUnitArrayTreeStringVoidResult
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyFnUnitArrayTreeStringVoidResult
	M        func() *ProxyFnUnitArrayTreeStringVoidResult
	Pure     func(struct{}) []*tree_result.Tree
	Tag      uint8
}

const (
	ProxyFnUnitArrayTreeStringVoidResultRequest = iota
	ProxyFnUnitArrayTreeStringVoidResultRespond
	ProxyFnUnitArrayTreeStringVoidResultM
	ProxyFnUnitArrayTreeStringVoidResultPure
)

type EitherArrayArrayTreeStringVoidResult struct {
	Left  exception.Error
	Right [][]*tree_result.Tree
	Tag   uint8
}

const (
	EitherArrayArrayTreeStringVoidResultLeft = iota
	EitherArrayArrayTreeStringVoidResultRight
)

type OnComplete struct {
	Handler func(func(EitherArrayArrayTreeStringVoidResult))
	Rethrow bool
}

type Fiber struct {
	IsSuspended func() bool
	Join        func(func(EitherArrayArrayTreeStringVoidResult)) func()
	Kill        func(exception.Error) func(func(either.EitherUnit)) func()
	OnComplete  func(OnComplete) func()
	Run         func()
}

type ProxyFiber struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyFiber
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyFiber
	M        func() *ProxyFiber
	Pure     Fiber
	Tag      uint8
}

const (
	ProxyFiberRequest = iota
	ProxyFiberRespond
	ProxyFiberM
	ProxyFiberPure
)

type ProxyEitherUnitEvent struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyEitherUnitEvent
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyEitherUnitEvent
	M        func() *ProxyEitherUnitEvent
	Pure     event.EitherUnit
	Tag      uint8
}

const (
	ProxyEitherUnitEventRequest = iota
	ProxyEitherUnitEventRespond
	ProxyEitherUnitEventM
	ProxyEitherUnitEventPure
)

type ProxyRefBoolean struct {
	Request0 internal.X
	Request1 func(struct{}) *ProxyRefBoolean
	Respond0 *event.Event
	Respond1 func(struct{}) *ProxyRefBoolean
	M        func() *ProxyRefBoolean
	Pure     *Ref[bool]
	Tag      uint8
}

const (
	ProxyRefBooleanRequest = iota
	ProxyRefBooleanRespond
	ProxyRefBooleanM
	ProxyRefBooleanPure
)

func Convert(value either.EitherFnUnitAffUnit) *either.EitherFnUnit {
	switch value.Tag {
	case either.EitherFnUnitAffUnitLeft:
		return &either.EitherFnUnit{Tag: either.EitherFnUnitLeft, Left: value.Left, Origin: value}
	default: // either.EitherFnUnitAffUnitRight
		return &either.EitherFnUnit{Tag: either.EitherFnUnitRight, Right: func(coerced any) any {
			original := func() {
				value.Right(coerced.(struct{}))
			}
			return func() any {
				original()
				return struct{}{}
			}
		}, Origin: value}
	}
}

func Convert2(value tree.Item) maybe.ExampleIsFocusedIsParallelizable {
	return maybe.ExampleIsFocusedIsParallelizable{Example: func(coerced func(func(any) any) any) any {
		original3 := func() {
			value.Example(func(coerced2 func(struct{})) {
				original2 := coerced(func(coerced3 any) any {
					original := func() {
						coerced2(coerced3.(struct{}))
					}
					return func() any {
						original()
						return struct{}{}
					}
				}).(func() any)
				func() {
					original2()
				}()
			})
		}
		return func() any {
			original3()
			return struct{}{}
		}
	}, IsFocused: value.IsFocused, IsParallelizable: value.IsParallelizable}
}

func Convert3(value tree.MaybeItem) *maybe.MaybeExampleIsFocusedIsParallelizable {
	switch value.Tag {
	case maybe.TagNothing:
		return &maybe.MaybeExampleIsFocusedIsParallelizable{Tag: maybe.TagNothing, Origin: value}
	default: // maybe.TagJust
		return &maybe.MaybeExampleIsFocusedIsParallelizable{Tag: maybe.TagJust, Just: Convert2(value.Just), Origin: value}
	}
}

func Convert4(value *tree.TreeStringFnUnitAffUnit) *tree.TreeFnUnit {
	if value == nil {
		return nil
	}
	switch value.Tag {
	case tree.TreeStringFnUnitAffUnitNode:
		return &tree.TreeFnUnit{Tag: tree.TreeFnUnitNode, Node0: Convert(value.Node0), Node1: Convert5(value.Node1), Origin: value}
	default: // tree.TreeStringFnUnitAffUnitLeaf
		return &tree.TreeFnUnit{Tag: tree.TreeFnUnitLeaf, Leaf0: value.Leaf0, Leaf1: Convert3(value.Leaf1), Origin: value}
	}
}

func Convert5(value []*tree.TreeStringFnUnitAffUnit) []*tree.TreeFnUnit {
	converted := make([]*tree.TreeFnUnit, len(value))
	for index, element := range value {
		converted[index] = Convert4(element)
	}
	return converted
}

func Convert6(value *either.EitherFnUnit) either.EitherFnUnitAffUnit {
	if original, ok := value.Origin.(either.EitherFnUnitAffUnit); ok {
		return original
	}
	switch value.Tag {
	case either.EitherFnUnitLeft:
		return either.EitherFnUnitAffUnit{Tag: either.EitherFnUnitAffUnitLeft, Left: value.Left}
	default: // either.EitherFnUnitRight
		return either.EitherFnUnitAffUnit{Tag: either.EitherFnUnitAffUnitRight, Right: func(coerced struct{}) {
			original := value.Right(coerced).(func() any)
			func() {
				original()
			}()
		}}
	}
}

func Convert7(value maybe.ExampleIsFocusedIsParallelizable) tree.Item {
	return tree.Item{Example: func(coerced func(func(struct{}))) {
		original3 := value.Example(func(coerced2 func(any) any) any {
			original2 := func() {
				coerced(func(coerced3 struct{}) {
					original := coerced2(coerced3).(func() any)
					func() {
						original()
					}()
				})
			}
			return func() any {
				original2()
				return struct{}{}
			}
		}).(func() any)
		func() {
			original3()
		}()
	}, IsFocused: value.IsFocused, IsParallelizable: value.IsParallelizable}
}

func Convert8(value *maybe.MaybeExampleIsFocusedIsParallelizable) tree.MaybeItem {
	if original, ok := value.Origin.(tree.MaybeItem); ok {
		return original
	}
	switch value.Tag {
	case maybe.TagNothing:
		return tree.NoItem
	default: // maybe.TagJust
		return tree.JustItem(Convert7(value.Just))
	}
}

func Convert9(value *tree.TreeFnUnit) *tree.TreeStringFnUnitAffUnit {
	if value == nil {
		return nil
	}
	if original, ok := value.Origin.(*tree.TreeStringFnUnitAffUnit); ok {
		return original
	}
	switch value.Tag {
	case tree.TreeFnUnitNode:
		return &tree.TreeStringFnUnitAffUnit{Tag: tree.TreeStringFnUnitAffUnitNode, Node0: Convert6(value.Node0), Node1: Convert10(value.Node1)}
	default: // tree.TreeFnUnitLeaf
		return &tree.TreeStringFnUnitAffUnit{Tag: tree.TreeStringFnUnitAffUnitLeaf, Leaf0: value.Leaf0, Leaf1: Convert8(value.Leaf1)}
	}
}

func Convert10(value []*tree.TreeFnUnit) []*tree.TreeStringFnUnitAffUnit {
	converted := make([]*tree.TreeStringFnUnitAffUnit, len(value))
	for index, element := range value {
		converted[index] = Convert9(element)
	}
	return converted
}

func RightEitherArrayTreeStringVoidResult(right []*tree_result.Tree) EitherArrayTreeStringVoidResult {
	return EitherArrayTreeStringVoidResult{Tag: EitherArrayTreeStringVoidResultRight, Right: right}
}

func LeftEitherArrayTreeStringVoidResult(left exception.Error) EitherArrayTreeStringVoidResult {
	return EitherArrayTreeStringVoidResult{Tag: EitherArrayTreeStringVoidResultLeft, Left: left}
}

func RespondProxyEventArrayTreeStringVoidResult(respond0 *event.Event, respond1 func(struct{}) *ProxyEventArrayTreeStringVoidResult) *ProxyEventArrayTreeStringVoidResult {
	return &ProxyEventArrayTreeStringVoidResult{Tag: ProxyEventArrayTreeStringVoidResultRespond, Respond0: respond0, Respond1: respond1}
}

func MProxyEventArrayTreeStringVoidResult(m func() *ProxyEventArrayTreeStringVoidResult) *ProxyEventArrayTreeStringVoidResult {
	return &ProxyEventArrayTreeStringVoidResult{Tag: ProxyEventArrayTreeStringVoidResultM, M: m}
}

func PureProxyEventArrayTreeStringVoidResult(pure []*tree_result.Tree) *ProxyEventArrayTreeStringVoidResult {
	return &ProxyEventArrayTreeStringVoidResult{Tag: ProxyEventArrayTreeStringVoidResultPure, Pure: pure}
}

func RequestProxyEventArrayTreeStringVoidResult(request0 internal.X, request1 func(struct{}) *ProxyEventArrayTreeStringVoidResult) *ProxyEventArrayTreeStringVoidResult {
	return &ProxyEventArrayTreeStringVoidResult{Tag: ProxyEventArrayTreeStringVoidResultRequest, Request0: request0, Request1: request1}
}

func PureProxyBoolean(pure bool) *ProxyBoolean {
	return &ProxyBoolean{Tag: ProxyBooleanPure, Pure: pure}
}

func PureProxyArrayArrayTreeStringVoidResult(pure [][]*tree_result.Tree) *ProxyArrayArrayTreeStringVoidResult {
	return &ProxyArrayArrayTreeStringVoidResult{Tag: ProxyArrayArrayTreeStringVoidResultPure, Pure: pure}
}

func PureProxyAVarEitherUnitEvent(pure avar.AVar[event.EitherUnit]) *ProxyAVarEitherUnitEvent {
	return &ProxyAVarEitherUnitEvent{Tag: ProxyAVarEitherUnitEventPure, Pure: pure}
}

func Leaf(leaf0 string, leaf1 specresult.MaybeResult) *tree_result.Tree {
	return &tree_result.Tree{Tag: tree_result.TreeLeaf, Leaf0: leaf0, Leaf1: leaf1}
}

func RespondProxyUnit(respond0 *event.Event, respond1 func(struct{}) *ProxyUnit) *ProxyUnit {
	return &ProxyUnit{Tag: ProxyUnitRespond, Respond0: respond0, Respond1: respond1}
}

func MProxyUnit(m func() *ProxyUnit) *ProxyUnit {
	return &ProxyUnit{Tag: ProxyUnitM, M: m}
}

func PureProxyResult(pure specresult.Result) *ProxyResult {
	return &ProxyResult{Tag: ProxyResultPure, Pure: pure}
}

func RequestProxyFnUnitArrayTreeStringVoidResult(request0 internal.X, request1 func(struct{}) *ProxyFnUnitArrayTreeStringVoidResult) *ProxyFnUnitArrayTreeStringVoidResult {
	return &ProxyFnUnitArrayTreeStringVoidResult{Tag: ProxyFnUnitArrayTreeStringVoidResultRequest, Request0: request0, Request1: request1}
}

func RespondProxyFnUnitArrayTreeStringVoidResult(respond0 *event.Event, respond1 func(struct{}) *ProxyFnUnitArrayTreeStringVoidResult) *ProxyFnUnitArrayTreeStringVoidResult {
	return &ProxyFnUnitArrayTreeStringVoidResult{Tag: ProxyFnUnitArrayTreeStringVoidResultRespond, Respond0: respond0, Respond1: respond1}
}

func MProxyFnUnitArrayTreeStringVoidResult(m func() *ProxyFnUnitArrayTreeStringVoidResult) *ProxyFnUnitArrayTreeStringVoidResult {
	return &ProxyFnUnitArrayTreeStringVoidResult{Tag: ProxyFnUnitArrayTreeStringVoidResultM, M: m}
}

func PureProxyFnUnitArrayTreeStringVoidResult(pure func(struct{}) []*tree_result.Tree) *ProxyFnUnitArrayTreeStringVoidResult {
	return &ProxyFnUnitArrayTreeStringVoidResult{Tag: ProxyFnUnitArrayTreeStringVoidResultPure, Pure: pure}
}

func Node(node0 *either.EitherVoid, node1 []*tree_result.Tree) *tree_result.Tree {
	return &tree_result.Tree{Tag: tree_result.TreeNode, Node0: node0, Node1: node1}
}

func RequestProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult(request0 internal.X, request1 func(struct{}) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult {
	return &ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult{Tag: ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultRequest, Request0: request0, Request1: request1}
}

func RespondProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult(respond0 *event.Event, respond1 func(struct{}) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult {
	return &ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult{Tag: ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultRespond, Respond0: respond0, Respond1: respond1}
}

func MProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult(m func() *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult {
	return &ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult{Tag: ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultM, M: m}
}

func PureProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult(pure func([]*tree_result.Tree) [][]*tree_result.Tree) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult {
	return &ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult{Tag: ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultPure, Pure: pure}
}

func RequestProxyArrayArrayTreeStringVoidResult(request0 internal.X, request1 func(struct{}) *ProxyArrayArrayTreeStringVoidResult) *ProxyArrayArrayTreeStringVoidResult {
	return &ProxyArrayArrayTreeStringVoidResult{Tag: ProxyArrayArrayTreeStringVoidResultRequest, Request0: request0, Request1: request1}
}

func RespondProxyArrayArrayTreeStringVoidResult(respond0 *event.Event, respond1 func(struct{}) *ProxyArrayArrayTreeStringVoidResult) *ProxyArrayArrayTreeStringVoidResult {
	return &ProxyArrayArrayTreeStringVoidResult{Tag: ProxyArrayArrayTreeStringVoidResultRespond, Respond0: respond0, Respond1: respond1}
}

func MProxyArrayArrayTreeStringVoidResult(m func() *ProxyArrayArrayTreeStringVoidResult) *ProxyArrayArrayTreeStringVoidResult {
	return &ProxyArrayArrayTreeStringVoidResult{Tag: ProxyArrayArrayTreeStringVoidResultM, M: m}
}

func LeftEitherArrayArrayTreeStringVoidResult(left exception.Error) EitherArrayArrayTreeStringVoidResult {
	return EitherArrayArrayTreeStringVoidResult{Tag: EitherArrayArrayTreeStringVoidResultLeft, Left: left}
}

func RightEitherArrayArrayTreeStringVoidResult(right [][]*tree_result.Tree) EitherArrayArrayTreeStringVoidResult {
	return EitherArrayArrayTreeStringVoidResult{Tag: EitherArrayArrayTreeStringVoidResultRight, Right: right}
}

func PureProxyFiber(pure Fiber) *ProxyFiber {
	return &ProxyFiber{Tag: ProxyFiberPure, Pure: pure}
}

func PureProxyEitherUnitEvent(pure event.EitherUnit) *ProxyEitherUnitEvent {
	return &ProxyEitherUnitEvent{Tag: ProxyEitherUnitEventPure, Pure: pure}
}

func PureProxyRefBoolean(pure *Ref[bool]) *ProxyRefBoolean {
	return &ProxyRefBoolean{Tag: ProxyRefBooleanPure, Pure: pure}
}

func RequestProxyX(field0 internal.X, field1 func(struct{}) *ProxyX) *ProxyX {
	return &ProxyX{Tag: ProxyXRequest, Field0: field0, Field1: field1}
}

func MProxyX(m func() *ProxyX) *ProxyX {
	return &ProxyX{Tag: ProxyXM, M: m}
}

func PureProxyX(pure []*tree_result.Tree) *ProxyX {
	return &ProxyX{Tag: ProxyXPure, Pure: pure}
}

func RespondProxyX(field0 internal.X, field1 func(struct{}) *ProxyX) *ProxyX {
	return &ProxyX{Tag: ProxyXRespond, Field0: field0, Field1: field1}
}

func EvalSpecT(config tree.Config, reporters []*event_internal.Proxy, spec tree.TupleUnit) []*tree_result.Tree {
	value := spec.Second
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:165:1
	bound := make([]*tree.TreeStringFnUnitAffUnit, 0, len(value))
	for _, x := range value {
//line ../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
		scrutinee := tree.FilterTreeF(x)
		if scrutinee != nil {
//line ../../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
			bound = append(bound, scrutinee)
		}
	}
	var tests []*tree.TreeStringFnUnitAffUnit
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:165:1
	if len(bound) == 0 {
		tests = value
	} else {
		tests = bound
	}
//line ../../../../../../p/prelude-6.0.2/src/Data/Functor.purs:42:19
	reportedEvents := func() []*tree_result.Tree {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:190:22
		append_ := direct.Appender[struct{}]()
		accumulator := func(s int) tuple.TupleArrayUnit {
			return tuple.TupleArrayUnit{First: []struct{}{}, Second: s}
		}
		for _, element := range tests {
			accumulator2 := accumulator
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:190:22
			a := EvalSpecTGo(element)
			accumulator = func(s int) tuple.TupleArrayUnit {
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
				scrutinee := accumulator2(s)
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:30
				scrutinee2 := a(scrutinee.Second)
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:30
				return tuple.TupleArrayUnit{First: append_(scrutinee.First)(scrutinee2.First), Second: scrutinee2.Second}
			}
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:190:22
		scrutinee3 := accumulator(0)
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:190:22
		accumulator3 := EvalSpecTGoRespond(config, tests, event.Start(scrutinee3.Second))
		for _, element := range reporters {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:190:22
			accumulator3 = ComposePull_Fb_(accumulator3, element)
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:190:22
		return RunEffect(EvalSpecTGo2(accumulator3))
	}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:191:3
	graphicsParam := codes.GraphicsParam{Tag: codes.GraphicsParamReset}
	nonEmptyListGraphicsParam := codes.NonEmpty{First: graphicsParam, Second: codes.SharedListGraphicsParamNil}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:191:3
	escapeCode := codes.Graphics(nonEmptyListGraphicsParam)
	if config.Exit {
		array := []string{"WARNING: The use of `runSpec` or `runSpecT` under NodeJS is deprecated ", "and will be removed in the next major release. ", "Please migrate to `runSpecAndExitProcess` from the 'spec-node' package."}
		accumulator := ""
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:193:58
		for index := len(array) - 1; index >= 0; index-- {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:193:58
			x := array[index]
//line ../../../../../../p/foldable-traversable-6.0.0/src/Data/Foldable.purs:115:38
			accumulator = prim.Append(x, accumulator)
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
		accumulator2 := codes.SharedListGraphicsParamNil
		var listGraphicsParamCells alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
		for index := len(style.Yellow) - 1; index >= 0; index-- {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
			graphicsParam2 := style.Yellow[index]
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
			accumulator2 = listGraphicsParamCells.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator2})
		}
		var result string
		if accumulator2 == codes.SharedListGraphicsParamNil {
			result = accumulator
		} else {
			x := accumulator2.Cons0
			xs := accumulator2.Cons1
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			var text []byte
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = prim.AppendString(text, accumulator)
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
			text = codes.AppendEscapeCodeToString(text, escapeCode)
			result = string(text)
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
		specpkg.Write(result)
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:198:7
		either := direct.CatchError(func() EitherArrayTreeStringVoidResult {
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/Error/Class.purs:67:10
			array2 := reportedEvents()
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/Error/Class.purs:67:10
			return RightEitherArrayTreeStringVoidResult(array2)
		}, func(x exception.Error) EitherArrayTreeStringVoidResult {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:198:7
			return LeftEitherArrayTreeStringVoidResult(x)
		})
		switch either.Tag {
		case EitherArrayTreeStringVoidResultLeft:
			err := either.Left
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:200:58
			str := exception.ShowErrorImpl(err) + "\n"
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
			accumulator3 := codes.SharedListGraphicsParamNil
			var listGraphicsParamCells2 alloc.Cells[codes.ListGraphicsParam]
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
			for index := len(style.Red) - 1; index >= 0; index-- {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
				graphicsParam2 := style.Red[index]
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
				accumulator3 = listGraphicsParamCells2.New(codes.ListGraphicsParam{Cons0: graphicsParam2, Cons1: accumulator3})
			}
			var result2 string
			if accumulator3 == codes.SharedListGraphicsParamNil {
				result2 = str
			} else {
				x := accumulator3.Cons0
				xs := accumulator3.Cons1
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
				var text []byte
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
				text = codes.AppendEscapeCodeToString(text, codes.Graphics(codes.NonEmpty{First: x, Second: xs}))
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
				text = prim.AppendString(text, str)
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:20:7
				text = codes.AppendEscapeCodeToString(text, escapeCode)
				result2 = string(text)
			}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Style.purs:16:3
			specpkg.Write(result2)
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:201:11
			specpkg.Exit(1)
			panic(err)
		default: // EitherArrayTreeStringVoidResultRight
			results := either.Right
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:204:22
			var code int
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:204:22
			if summary.Summarize(results).Failed == 0 {
				code = 0
			} else {
				code = 1
			}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:205:11
			specpkg.Exit(code)
			return results
		}
	} else {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:208:7
		return reportedEvents()
	}
}

//line runner.go:723
var RunEffect func(*ProxyX) []*tree_result.Tree = RunEffectGo

const (
	ComposePull_Step = iota
	ComposePush_Step
)

func ComposePull_OrComposePush_UnitXEventUnitEventUnitAffArrayTreeStringVoidResult(step int, fb_ func(struct{}) *ProxyEventArrayTreeStringVoidResult, p *event_internal.Proxy, p2 *ProxyEventArrayTreeStringVoidResult, fb func(*event.Event) *event_internal.Proxy) *ProxyEventArrayTreeStringVoidResult {
	for {
		switch step {
		case ComposePull_Step:
			fb_Captured := fb_
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
			switch p.Tag {
			case event_internal.ProxyRequest:
				fb2 := p.Request1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:198:23
				step, p2, fb = ComposePush_Step, fb_Captured(struct{}{}), fb2
				continue
			case event_internal.ProxyRespond:
				c := p.Respond0
				fc_ := p.Respond1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
				return RespondProxyEventArrayTreeStringVoidResult(c, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
					return ComposePull_(fb_Captured, fc_(x))
				})
			case event_internal.ProxyM:
				m := p.M
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
				return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:200:26
					proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:200:26
					return ComposePull_(fb_Captured, proxy)
				})
			default: // event_internal.ProxyPure
				r := p.Pure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
				return PureProxyEventArrayTreeStringVoidResult(r)
			}
		default: // ComposePush_Step
			fbCaptured := fb
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:222:21
			switch p2.Tag {
			case ProxyEventArrayTreeStringVoidResultRequest:
				a_ := p2.Request0
				fa := p2.Request1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:222:21
				return RequestProxyEventArrayTreeStringVoidResult(a_, func(a struct{}) *ProxyEventArrayTreeStringVoidResult {
					return ComposePush_(fa(a), fbCaptured)
				})
			case ProxyEventArrayTreeStringVoidResultRespond:
				b := p2.Respond0
				fb_2 := p2.Respond1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:224:23
				step, fb_, p = ComposePull_Step, fb_2, fbCaptured(b)
				continue
			case ProxyEventArrayTreeStringVoidResultM:
				m := p2.M
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:222:21
				return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:225:26
					proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:225:39
					return ComposePush_(proxy, fbCaptured)
				})
			default: // ProxyEventArrayTreeStringVoidResultPure
				r := p2.Pure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:222:21
				return PureProxyEventArrayTreeStringVoidResult(r)
			}
		}
	}
}

//line runner.go:800
func ComposePull_(fb_ func(struct{}) *ProxyEventArrayTreeStringVoidResult, p *event_internal.Proxy) *ProxyEventArrayTreeStringVoidResult {
	return ComposePull_OrComposePush_UnitXEventUnitEventUnitAffArrayTreeStringVoidResult(ComposePull_Step, fb_, p, nil, nil)
}

func Closed(x internal.X) []*tree_result.Tree {
	for {
	}
}

var RunEffectRec func(*ProxyX) []*tree_result.Tree = RunEffectRecGo

func ComposePush_(p *ProxyEventArrayTreeStringVoidResult, fb func(*event.Event) *event_internal.Proxy) *ProxyEventArrayTreeStringVoidResult {
	return ComposePull_OrComposePush_UnitXEventUnitEventUnitAffArrayTreeStringVoidResult(ComposePush_Step, nil, nil, p, fb)
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:35:18
func Apply(pf0 *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult, px *ProxyEventArrayTreeStringVoidResult) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:35:18
	return ApplyGo(px, pf0)
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:135:25
func ComposeResponseFb(p0 *ProxyEventArrayTreeStringVoidResult, aVar avar.AVar[event.EitherUnit]) *ProxyX {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:135:25
	return ComposeResponseFbGo(aVar, p0)
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
func ComposePull_Fb_(proxy *ProxyEventArrayTreeStringVoidResult, p *event_internal.Proxy) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
	switch p.Tag {
	case event_internal.ProxyRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:198:23
		return ComposePush_(proxy, p.Request1)
	case event_internal.ProxyRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
			return ComposePull_Fb_(proxy, p.Respond1(x))
		})
	case event_internal.ProxyM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:200:26
			proxy2 := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:200:26
			return ComposePull_Fb_(proxy, proxy2)
		})
	default: // event_internal.ProxyPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:197:22
		return PureProxyEventArrayTreeStringVoidResult(p.Pure)
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
func EvalSpecTRunItem(config tree.Config, keepRunningVar *Ref[bool], element Element) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
	return EvalSpecTRunItemGoM(keepRunningVar, element.IsParallelizable, element.Test, config, func() *ProxyBoolean {
		boolean := keepRunningVar.Value
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return PureProxyBoolean(boolean)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
func EvalSpecTLoop(config tree.Config, keepRunningVar *Ref[bool], tests []*tree.TreeTupleArrayPathItemString) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
	append_ := direct.Appender[[]*tree_result.Tree]()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
	mapped := make([]Element, len(tests))
	for index, test := range tests {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:102:47
		argument := tree.IsAllParallelizable(test)
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:102:47
		mapped[index] = Element{IsParallelizable: argument, Test: test}
	}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
	array := arraypkg.GroupBy(func(a Element, b Element) bool {
		return a.IsParallelizable == b.IsParallelizable
	}, mapped)
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
	accumulator := PureProxyArrayArrayTreeStringVoidResult([][]*tree_result.Tree{})
	for _, element := range array {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:107:11
		var result2 Element
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:107:11
		if len(element) > 0 {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:107:11
			result2 = element[0]
		} else {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:107:11
			panic("Failed pattern match")
		}
		var result *ProxyEventArrayTreeStringVoidResult
		if result2.IsParallelizable {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:108:45
			mapped2 := make([]*ProxyEventArrayTreeStringVoidResult, len(element))
			for index, element2 := range element {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:108:45
				mapped2[index] = EvalSpecTRunItem(config, keepRunningVar, element2)
			}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
			result = EvalSpecTLoopGo(EvalSpecTLoopGoM(mapped2, func() *ProxyAVarEitherUnitEvent {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:29:9
				aVar := avar.Empty[event.EitherUnit]()
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:29:9
				return PureProxyAVarEitherUnitEvent(aVar)
			}))
		} else {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
			append_2 := direct.Appender[[]*tree_result.Tree]()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
			accumulator2 := PureProxyArrayArrayTreeStringVoidResult([][]*tree_result.Tree{})
			for _, element2 := range element {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
				accumulator2 = Apply(EvalSpecTLoopGo2(append_2, accumulator2), EvalSpecTRunItem(config, keepRunningVar, element2))
			}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
			result = EvalSpecTLoopGo(accumulator2)
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:107:11
		accumulator = Apply(EvalSpecTLoopGo2(append_, accumulator), result)
	}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
	return EvalSpecTLoopGo(accumulator)
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTRunItemGo(keepRunningVar *Ref[bool], isParallelizable bool, test *tree.TreeTupleArrayPathItemString, config tree.Config, p *ProxyBoolean) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyBooleanRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTRunItemGo(keepRunningVar, isParallelizable, test, config, p.Request1(x))
		})
	case ProxyBooleanRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTRunItemGo(keepRunningVar, isParallelizable, test, config, p.Respond1(x))
		})
	case ProxyBooleanM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTRunItemGo(keepRunningVar, isParallelizable, test, config, proxy)
		})
	default: // ProxyBooleanPure
		keepRunning := p.Pure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:51:25
		if test.Tag == tree.TreeTupleArrayPathItemStringLeaf && test.Leaf1.Tag == maybe.TagJust {
			path := test.Leaf0.First
			name := test.Leaf0.Second
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:117:11
			if keepRunning {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
				var result event.Execution
				if isParallelizable {
					result = event.ExecutionParallel
				} else {
					result = event.ExecutionSequential
				}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
				return EvalSpecTRunItemGoGoRespond(path, name, test.Leaf1.Just, keepRunningVar, config, event.Test(result, tree.TupleArrayPathItem{First: path, Second: name}))
			}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:117:11
			return PureProxyEventArrayTreeStringVoidResult([]*tree_result.Tree{Leaf(name, specresult.NoResult)})
		}
		if test.Tag == tree.TreeTupleArrayPathItemStringLeaf && test.Leaf1.Tag == maybe.TagNothing {
			name := test.Leaf0.Second
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
			var result *ProxyUnit
			if keepRunning {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
				result = RespondProxyUnit(event.Pending(tree.TupleArrayPathItem{First: test.Leaf0.First, Second: name}), PureProxyUnit)
			} else {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
				result = PureProxyUnit(struct{}{})
			}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
			return EvalSpecTRunItemGoGo(name, result)
		}
		if test.Tag == tree.TreeTupleArrayPathItemStringNode && test.Node0.Tag == tree.EitherRight {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:134:11
			b := MProxyUnit(func() *ProxyUnit {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:134:43
				test.Node0.Right(struct{}{})
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:134:43
				return PureProxyUnit(struct{}{})
			})
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:35:18
			return EvalSpecTRunItemGoGo3(b, EvalSpecTRunItemGoGo2(EvalSpecTLoop(config, keepRunningVar, test.Node1)))
		}
		path2 := test.Node0.Left.First
		name2 := test.Node0.Left.Second
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		var result2 *ProxyUnit
		if keepRunning {
			var result event.Execution
			if isParallelizable {
				result = event.ExecutionParallel
			} else {
				result = event.ExecutionSequential
			}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
			result2 = RespondProxyUnit(event.Suite(result, tree.TupleArrayPathItem{First: path2, Second: name2}), PureProxyUnit)
		} else {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
			result2 = PureProxyUnit(struct{}{})
		}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return EvalSpecTRunItemGoGo4(keepRunning, path2, name2, test.Node1, keepRunningVar, config, result2)
	}
}

//line runner.go:1019
func EvalSpecTRunItemGoM(keepRunningVar *Ref[bool], isParallelizable bool, test *tree.TreeTupleArrayPathItemString, config tree.Config, m func() *ProxyBoolean) *ProxyEventArrayTreeStringVoidResult {
	return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
		proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
		return EvalSpecTRunItemGo(keepRunningVar, isParallelizable, test, config, proxy)
	})
}

//line runner.go:1029
func EvalSpecTRunItemGoGoRespond(path []tree.PathItem, name string, item tree.Item, keepRunningVar *Ref[bool], config tree.Config, b *event.Event) *ProxyEventArrayTreeStringVoidResult {
	return RespondProxyEventArrayTreeStringVoidResult(b, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
		return EvalSpecTRunItemGoGoPure(path, name, item, keepRunningVar, config, x)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
func EvalSpecTRunItemGoGoPure(path []tree.PathItem, name string, item tree.Item, keepRunningVar *Ref[bool], config tree.Config, r struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
	return EvalSpecTRunItemGoGoPureGoM(path, name, keepRunningVar, config, func() *ProxyResult {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:145:16
		start := now.Now()
		var wrap func(func())
		switch config.Timeout.Tag {
		case maybe.TagNothing:
			wrap = category.Identity2
		default: // maybe.TagJust
			wrap = func(t func()) {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:70:3
				either2 := aff.ParAltDirect(func() either.EitherUnit {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:70:15
					return direct.CatchError(func() either.EitherUnit {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:61:3
						aff.DelayDirect(nil, config.Timeout.Just)
						aff.MakeAffDirect(func(value either.EitherUnit) bool {
							return value.Tag == either.EitherUnitLeft
						}, func(value either.EitherUnit) exception.Error {
							return value.Left
						}, func(value either.EitherUnit) struct{} {
							return value.Right
						}, func(cb func(either.EitherUnit)) func(exception.Error) {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:63:5
							cb(either.LeftEitherUnit(exception.NewError("test timed out after " + strconv.Itoa(prim.IntRound(config.Timeout.Just)) + "ms")))
							return effectaff.NonCanceler
						})
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:62:3
						return either.RightEitherUnit(struct{}{})
					}, func(x exception.Error) either.EitherUnit {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:70:15
						return either.LeftEitherUnit(x)
					})
				}, func() either.EitherUnit {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:70:53
					return direct.CatchError(func() either.EitherUnit {
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/Error/Class.purs:67:10
						t()
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/Error/Class.purs:67:10
						return either.RightEitherUnit(struct{}{})
					}, func(x exception.Error) either.EitherUnit {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:70:53
						return either.LeftEitherUnit(x)
					})
				})
//line ../../../../../../p/either-6.1.0/src/Data/Either.purs:209:1
				switch either2.Tag {
				case either.EitherUnitLeft:
//line ../../../../../../p/either-6.1.0/src/Data/Either.purs:209:23
					effectaff.ThrowError(either2.Left)
				default: // either.EitherUnitRight
				}
			}
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:147:12
		either3 := direct.CatchError(func() either.EitherUnit {
//line ../../../../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
			wrap(func() {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:147:29
				item.Example(func(a func(struct{})) {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:147:37
					a(struct{}{})
				})
			})
//line ../../../../../../p/prelude-6.0.2/src/Data/Function.purs:49:13
			return either.RightEitherUnit(struct{}{})
		}, func(x exception.Error) either.EitherUnit {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:147:12
			return either.LeftEitherUnit(x)
		})
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:148:14
		end := now.Now()
		duration := end + -start
		thresh := config.Slow
		var result specresult.Result
		switch either3.Tag {
		case either.EitherUnitLeft:
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:148:14
			result = specresult.Failure(either3.Left)
		default: // either.EitherUnitRight
			var result2 speed.Speed
			if !(duration <= thresh) {
				result2 = speed.SpeedSlow
			} else if !(duration <= thresh/2.0) {
				result2 = speed.SpeedMedium
			} else {
				result2 = speed.SpeedFast
			}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:148:14
			result = specresult.Success(result2, duration)
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:148:14
		return PureProxyResult(result)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTRunItemGoGoPureGo(path []tree.PathItem, name string, keepRunningVar *Ref[bool], config tree.Config, p *ProxyResult) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyResultRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTRunItemGoGoPureGo(path, name, keepRunningVar, config, p.Request1(x))
		})
	case ProxyResultRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTRunItemGoGoPureGo(path, name, keepRunningVar, config, p.Respond1(x))
		})
	case ProxyResultM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTRunItemGoGoPureGo(path, name, keepRunningVar, config, proxy)
		})
	default: // ProxyResultPure
		r := p.Pure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		var result *ProxyUnit
		if r.Tag == specresult.ResultFailure && config.FailFast {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
			result = MProxyUnit(func() *ProxyUnit {
				keepRunningVar.Value = false
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
				return PureProxyUnit(struct{}{})
			})
		} else {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
			result = PureProxyUnit(struct{}{})
		}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return EvalSpecTRunItemGoGoPureGoGo(r, path, name, result)
	}
}

//line runner.go:1178
func EvalSpecTRunItemGoGoPureGoM(path []tree.PathItem, name string, keepRunningVar *Ref[bool], config tree.Config, m func() *ProxyResult) *ProxyEventArrayTreeStringVoidResult {
	return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
		proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
		return EvalSpecTRunItemGoGoPureGo(path, name, keepRunningVar, config, proxy)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTRunItemGoGoPureGoGo(r specresult.Result, path []tree.PathItem, name string, p *ProxyUnit) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyUnitRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTRunItemGoGoPureGoGo(r, path, name, p.Request1(x))
		})
	case ProxyUnitRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTRunItemGoGoPureGoGo(r, path, name, p.Respond1(x))
		})
	case ProxyUnitM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTRunItemGoGoPureGoGo(r, path, name, proxy)
		})
	default: // ProxyUnitPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return EvalSpecTRunItemGoGoPureGoGoGoRespond(r, name, event.TestEnd(tree.TupleArrayPathItem{First: path, Second: name}, r))
	}
}

//line runner.go:1217
func EvalSpecTRunItemGoGoPureGoGoGoRespond(r specresult.Result, name string, b *event.Event) *ProxyEventArrayTreeStringVoidResult {
	return RespondProxyEventArrayTreeStringVoidResult(b, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
		return EvalSpecTRunItemGoGoPureGoGoGoPure(r, name, x)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTRunItemGoGoPureGoGoGoPure(r specresult.Result, name string, r2 struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	return PureProxyEventArrayTreeStringVoidResult([]*tree_result.Tree{Leaf(name, specresult.JustResult(r))})
}

//line runner.go:1231
func EvalSpecTRunItemGoGo(name string, p *ProxyUnit) *ProxyEventArrayTreeStringVoidResult {
	switch p.Tag {
	case ProxyUnitRequest:
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:48:42
			return EvalSpecTRunItemGoGo(name, p.Request1(x))
		})
	case ProxyUnitRespond:
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTRunItemGoGo(name, p.Respond1(x))
		})
	case ProxyUnitM:
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTRunItemGoGo(name, proxy)
		})
	default: // ProxyUnitPure
		return PureProxyEventArrayTreeStringVoidResult([]*tree_result.Tree{Leaf(name, specresult.NoResult)})
	}
}

//line runner.go:1256
func EvalSpecTRunItemGoGo2(p *ProxyEventArrayTreeStringVoidResult) *ProxyFnUnitArrayTreeStringVoidResult {
	switch p.Tag {
	case ProxyEventArrayTreeStringVoidResultRequest:
		return RequestProxyFnUnitArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyFnUnitArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:29:42
			return EvalSpecTRunItemGoGo2(p.Request1(x))
		})
	case ProxyEventArrayTreeStringVoidResultRespond:
		return RespondProxyFnUnitArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyFnUnitArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:30:42
			return EvalSpecTRunItemGoGo2(p.Respond1(x))
		})
	case ProxyEventArrayTreeStringVoidResultM:
		return MProxyFnUnitArrayTreeStringVoidResult(func() *ProxyFnUnitArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:49
			return EvalSpecTRunItemGoGo2(proxy)
		})
	default: // ProxyEventArrayTreeStringVoidResultPure
		return PureProxyFnUnitArrayTreeStringVoidResult(func(_ struct{}) []*tree_result.Tree {
			return p.Pure
		})
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
func EvalSpecTRunItemGoGo3(b *ProxyUnit, pf *ProxyFnUnitArrayTreeStringVoidResult) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
	switch pf.Tag {
	case ProxyFnUnitArrayTreeStringVoidResultRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return RequestProxyEventArrayTreeStringVoidResult(pf.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTRunItemGoGo3(b, pf.Request1(x))
		})
	case ProxyFnUnitArrayTreeStringVoidResultRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return RespondProxyEventArrayTreeStringVoidResult(pf.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:38:42
			return EvalSpecTRunItemGoGo3(b, pf.Respond1(x))
		})
	case ProxyFnUnitArrayTreeStringVoidResultM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:39:37
			proxy := pf.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:39:49
			return EvalSpecTRunItemGoGo3(b, proxy)
		})
	default: // ProxyFnUnitArrayTreeStringVoidResultPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
		return EvalSpecTRunItemGoGoGo(pf.Pure, b)
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
func EvalSpecTRunItemGoGoGo(f func(struct{}) []*tree_result.Tree, p *ProxyUnit) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
	switch p.Tag {
	case ProxyUnitRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTRunItemGoGoGo(f, p.Request1(x))
		})
	case ProxyUnitRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:30:42
			return EvalSpecTRunItemGoGoGo(f, p.Respond1(x))
		})
	case ProxyUnitM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:49
			return EvalSpecTRunItemGoGoGo(f, proxy)
		})
	default: // ProxyUnitPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:32:25
		return PureProxyEventArrayTreeStringVoidResult(f(p.Pure))
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTRunItemGoGo4(keepRunning bool, path []tree.PathItem, name string, xs []*tree.TreeTupleArrayPathItemString, keepRunningVar *Ref[bool], config tree.Config, p *ProxyUnit) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyUnitRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTRunItemGoGo4(keepRunning, path, name, xs, keepRunningVar, config, p.Request1(x))
		})
	case ProxyUnitRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTRunItemGoGo4(keepRunning, path, name, xs, keepRunningVar, config, p.Respond1(x))
		})
	case ProxyUnitM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTRunItemGoGo4(keepRunning, path, name, xs, keepRunningVar, config, proxy)
		})
	default: // ProxyUnitPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return EvalSpecTRunItemGoGoGo2(keepRunning, path, name, EvalSpecTLoop(config, keepRunningVar, xs))
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTRunItemGoGoGo2(keepRunning bool, path []tree.PathItem, name string, p *ProxyEventArrayTreeStringVoidResult) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyEventArrayTreeStringVoidResultRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTRunItemGoGoGo2(keepRunning, path, name, p.Request1(x))
		})
	case ProxyEventArrayTreeStringVoidResultRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTRunItemGoGoGo2(keepRunning, path, name, p.Respond1(x))
		})
	case ProxyEventArrayTreeStringVoidResultM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTRunItemGoGoGo2(keepRunning, path, name, proxy)
		})
	default: // ProxyEventArrayTreeStringVoidResultPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		var result *ProxyUnit
		if keepRunning {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
			result = RespondProxyUnit(event.SuiteEnd(tree.TupleArrayPathItem{First: path, Second: name}), PureProxyUnit)
		} else {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
			result = PureProxyUnit(struct{}{})
		}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return EvalSpecTRunItemGoGoGoGo(p.Pure, name, result)
	}
}

//line runner.go:1408
func EvalSpecTRunItemGoGoGoGo(res []*tree_result.Tree, name string, p *ProxyUnit) *ProxyEventArrayTreeStringVoidResult {
	switch p.Tag {
	case ProxyUnitRequest:
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:48:42
			return EvalSpecTRunItemGoGoGoGo(res, name, p.Request1(x))
		})
	case ProxyUnitRespond:
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTRunItemGoGoGoGo(res, name, p.Respond1(x))
		})
	case ProxyUnitM:
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTRunItemGoGoGoGo(res, name, proxy)
		})
	default: // ProxyUnitPure
		return PureProxyEventArrayTreeStringVoidResult([]*tree_result.Tree{Node(either.LeftEitherVoid(name), res)})
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
func EvalSpecTLoopGo(p *ProxyArrayArrayTreeStringVoidResult) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
	switch p.Tag {
	case ProxyArrayArrayTreeStringVoidResultRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTLoopGo(p.Request1(x))
		})
	case ProxyArrayArrayTreeStringVoidResultRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:30:42
			return EvalSpecTLoopGo(p.Respond1(x))
		})
	case ProxyArrayArrayTreeStringVoidResultM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:49
			return EvalSpecTLoopGo(proxy)
		})
	default: // ProxyArrayArrayTreeStringVoidResultPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:32:25
		return PureProxyEventArrayTreeStringVoidResult(arraypkg.Concat(p.Pure))
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
func EvalSpecTLoopGo2(append_ func([][]*tree_result.Tree) func([]*tree_result.Tree) [][]*tree_result.Tree, p *ProxyArrayArrayTreeStringVoidResult) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
	switch p.Tag {
	case ProxyArrayArrayTreeStringVoidResultRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return RequestProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult {
			return EvalSpecTLoopGo2(append_, p.Request1(x))
		})
	case ProxyArrayArrayTreeStringVoidResultRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return RespondProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:30:42
			return EvalSpecTLoopGo2(append_, p.Respond1(x))
		})
	case ProxyArrayArrayTreeStringVoidResultM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return MProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult(func() *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:49
			return EvalSpecTLoopGo2(append_, proxy)
		})
	default: // ProxyArrayArrayTreeStringVoidResultPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:32:25
		return PureProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult(append_(p.Pure))
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTLoopGo3(ps []*ProxyEventArrayTreeStringVoidResult, p *ProxyAVarEitherUnitEvent) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyAVarEitherUnitEventRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyArrayArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
			return EvalSpecTLoopGo3(ps, p.Request1(x))
		})
	case ProxyAVarEitherUnitEventRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyArrayArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTLoopGo3(ps, p.Respond1(x))
		})
	case ProxyAVarEitherUnitEventM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTLoopGo3(ps, proxy)
		})
	default: // ProxyAVarEitherUnitEventPure
		var_ := p.Pure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return EvalSpecTLoopGoGoM(var_, func() *ProxyFiber {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:157:17
			fiber := aff.ForkFiber(true, func() [][]*tree_result.Tree {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:159:10
				array := aff.TraverseParArrayDirect(func(x *ProxyEventArrayTreeStringVoidResult) []*tree_result.Tree {
//line ../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
					return RunEffectRec(ComposeResponseFb(x, var_))
				}, ps)
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:160:5
				aff.MakeAffDirect(func(value either.EitherUnit) bool {
					return value.Tag == either.EitherUnitLeft
				}, func(value either.EitherUnit) exception.Error {
					return value.Left
				}, func(value either.EitherUnit) struct{} {
					return value.Right
				}, func(k func(either.EitherUnit)) func(exception.Error) {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:53:8
					effect := avar.PutVar(func(payload exception.Error) either.EitherUnit {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:53:8
						return either.LeftEitherUnit(payload)
					}, func(payload struct{}) either.EitherUnit {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:53:8
						return either.RightEitherUnit(payload)
					}, event.LeftEitherUnit(struct{}{}), var_, k)
					return func(_ exception.Error) {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:53:8
						effect()
					}
				})
				return array
			})
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:157:17
			isSuspended, join, kill, onComplete, run := aff.FiberFields(fiber, func(payload exception.Error) EitherArrayArrayTreeStringVoidResult {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:157:17
				return LeftEitherArrayArrayTreeStringVoidResult(payload)
			}, func(payload [][]*tree_result.Tree) EitherArrayArrayTreeStringVoidResult {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:157:17
				return RightEitherArrayArrayTreeStringVoidResult(payload)
			}, either.RightEitherUnit(struct{}{}), func(value OnComplete) func(func(EitherArrayArrayTreeStringVoidResult)) {
				return value.Handler
			})
			fiber2 := Fiber{IsSuspended: isSuspended, Join: join, Kill: kill, OnComplete: onComplete, Run: run}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:157:17
			return PureProxyFiber(fiber2)
		})
	}
}

//line runner.go:1565
func EvalSpecTLoopGoM(ps []*ProxyEventArrayTreeStringVoidResult, m func() *ProxyAVarEitherUnitEvent) *ProxyArrayArrayTreeStringVoidResult {
	return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
		proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
		return EvalSpecTLoopGo3(ps, proxy)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTLoopGoGo(var_ avar.AVar[event.EitherUnit], p *ProxyFiber) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyFiberRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyArrayArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
			return EvalSpecTLoopGoGo(var_, p.Request1(x))
		})
	case ProxyFiberRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyArrayArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTLoopGoGo(var_, p.Respond1(x))
		})
	case ProxyFiberM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTLoopGoGo(var_, proxy)
		})
	default: // ProxyFiberPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:51:25
		var loop *ProxyArrayArrayTreeStringVoidResult
		var go_ func(*ProxyEitherUnitEvent) *ProxyArrayArrayTreeStringVoidResult
		var goM func(func() *ProxyEitherUnitEvent) *ProxyArrayArrayTreeStringVoidResult
		go_ = func(p2 *ProxyEitherUnitEvent) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
			switch p2.Tag {
			case ProxyEitherUnitEventRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
				return RequestProxyArrayArrayTreeStringVoidResult(p2.Request0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
					return go_(p2.Request1(x))
				})
			case ProxyEitherUnitEventRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
				return RespondProxyArrayArrayTreeStringVoidResult(p2.Respond0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
					return go_(p2.Respond1(x))
				})
			case ProxyEitherUnitEventM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
				return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
					proxy := p2.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
					return go_(proxy)
				})
			default: // ProxyEitherUnitEventPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:51:25
				switch p2.Pure.Tag {
				case event.EitherUnitLeft:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:51:25
					return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:167:26
						array := aff.MakeAffDirect(func(value EitherArrayArrayTreeStringVoidResult) bool {
							return value.Tag == EitherArrayArrayTreeStringVoidResultLeft
						}, func(value EitherArrayArrayTreeStringVoidResult) exception.Error {
							return value.Left
						}, func(value EitherArrayArrayTreeStringVoidResult) [][]*tree_result.Tree {
							return value.Right
						}, func(k func(EitherArrayArrayTreeStringVoidResult)) func(exception.Error) {
//line ../../../../../../p/aff-8.0.0/src/Effect/Aff.purs:197:56
							effect := p.Pure.Join(k)
							return func(_ exception.Error) {
//line ../../../../../../p/aff-8.0.0/src/Effect/Aff.purs:197:56
								effect()
							}
						})
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:167:26
						return PureProxyArrayArrayTreeStringVoidResult(array)
					})
				default: // event.EitherUnitRight
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:169:11
					var goRespond func(*event.Event) *ProxyArrayArrayTreeStringVoidResult
					var goPure func(struct{}) *ProxyArrayArrayTreeStringVoidResult
					goRespond = func(b *event.Event) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:169:11
						return RespondProxyArrayArrayTreeStringVoidResult(b, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
							return goPure(x)
						})
					}
					goPure = func(_ struct{}) *ProxyArrayArrayTreeStringVoidResult {
						return loop
					}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
					return goRespond(p2.Pure.Right)
				}
			}
		}
		goM = func(m func() *ProxyEitherUnitEvent) *ProxyArrayArrayTreeStringVoidResult {
			return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
				proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
				return go_(proxy)
			})
		}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:51:25
		loop = goM(func() *ProxyEitherUnitEvent {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:165:20
			either := aff.MakeAffDirect(func(value event.EitherError) bool {
				return value.Tag == event.EitherErrorLeft
			}, func(value event.EitherError) exception.Error {
				return value.Left
			}, func(value event.EitherError) event.EitherUnit {
				return value.Right
			}, func(k func(event.EitherError)) func(exception.Error) {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:40:8
				effect := avar.TakeVar(func(payload exception.Error) event.EitherError {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:40:8
					return event.LeftEitherError(payload)
				}, func(payload event.EitherUnit) event.EitherError {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:40:8
					return event.RightEitherError(payload)
				}, var_, k)
				return func(_ exception.Error) {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:40:8
					effect()
				}
			})
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:165:20
			return PureProxyEitherUnitEvent(either)
		})
		return loop
	}
}

//line runner.go:1706
func EvalSpecTLoopGoGoM(var_ avar.AVar[event.EitherUnit], m func() *ProxyFiber) *ProxyArrayArrayTreeStringVoidResult {
	return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
		proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
		return EvalSpecTLoopGoGo(var_, proxy)
	})
}

//line runner.go:1716
func EvalSpecTGo(tree2 *tree.TreeStringFnUnitAffUnit) func(int) tuple.TupleUnit {
	switch tree2.Tag {
	case tree.TreeStringFnUnitAffUnitNode:
		return func(i int) tuple.TupleUnit {
			function := func(x *tree.TreeStringFnUnitAffUnit) func(func(int) tuple.TupleUnit) func(int) tuple.TupleUnit {
//line ../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:22
				a := EvalSpecTGo(x)
				return func(b func(int) tuple.TupleUnit) func(int) tuple.TupleUnit {
					return func(s int) tuple.TupleUnit {
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:5
						scrutinee := a(s)
//line ../../../../../../p/transformers-6.1.0/src/Control/Monad/State/Trans.purs:78:30
						scrutinee2 := b(scrutinee.Second)
						return tuple.TupleUnit{First: scrutinee2.First, Second: scrutinee2.Second}
					}
				}
			}
			return arraypkg.Foldr(func(tree *tree.TreeStringFnUnitAffUnit, f func(int) tuple.TupleUnit) func(int) tuple.TupleUnit {
				return function(tree)(f)
			}, func(s int) tuple.TupleUnit {
				return tuple.TupleUnit{First: struct{}{}, Second: s}
			}, tree2.Node1)(i)
		}
	default: // tree.TreeStringFnUnitAffUnitLeaf
		return func(x int) tuple.TupleUnit {
			return tuple.TupleUnit{First: struct{}{}, Second: x + 1}
		}
	}
}

//line runner.go:1747
func EvalSpecTGoRespond(config tree.Config, tests []*tree.TreeStringFnUnitAffUnit, b *event.Event) *ProxyEventArrayTreeStringVoidResult {
	return RespondProxyEventArrayTreeStringVoidResult(b, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
		return EvalSpecTGoPure(config, tests, x)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
func EvalSpecTGoPure(config tree.Config, tests []*tree.TreeStringFnUnitAffUnit, r struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
	return EvalSpecTGoPureGoM(config, tests, func() *ProxyRefBoolean {
		ref := &Ref[bool]{Value: true}
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return PureProxyRefBoolean(ref)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTGoPureGo(config tree.Config, tests []*tree.TreeStringFnUnitAffUnit, p *ProxyRefBoolean) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyRefBooleanRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTGoPureGo(config, tests, p.Request1(x))
		})
	case ProxyRefBooleanRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTGoPureGo(config, tests, p.Respond1(x))
		})
	case ProxyRefBooleanM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTGoPureGo(config, tests, proxy)
		})
	default: // ProxyRefBooleanPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return EvalSpecTGoPureGoGo(EvalSpecTLoop(config, p.Pure, tree.AnnotatedWithPaths(Convert10(config.FilterTree(Convert5(tests))))))
	}
}

//line runner.go:1794
func EvalSpecTGoPureGoM(config tree.Config, tests []*tree.TreeStringFnUnitAffUnit, m func() *ProxyRefBoolean) *ProxyEventArrayTreeStringVoidResult {
	return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
		proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
		return EvalSpecTGoPureGo(config, tests, proxy)
	})
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
func EvalSpecTGoPureGoGo(p *ProxyEventArrayTreeStringVoidResult) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
	switch p.Tag {
	case ProxyEventArrayTreeStringVoidResultRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RequestProxyEventArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
			return EvalSpecTGoPureGoGo(p.Request1(x))
		})
	case ProxyEventArrayTreeStringVoidResultRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return RespondProxyEventArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
			return EvalSpecTGoPureGoGo(p.Respond1(x))
		})
	case ProxyEventArrayTreeStringVoidResultM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
		return MProxyEventArrayTreeStringVoidResult(func() *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
			return EvalSpecTGoPureGoGo(proxy)
		})
	default: // ProxyEventArrayTreeStringVoidResultPure
		r := p.Pure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
		return EvalSpecTGoPureGoGoGoRespond(r, event.End(r))
	}
}

//line runner.go:1834
func EvalSpecTGoPureGoGoGoRespond(r []*tree_result.Tree, b *event.Event) *ProxyEventArrayTreeStringVoidResult {
	return RespondProxyEventArrayTreeStringVoidResult(b, func(x struct{}) *ProxyEventArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
		return EvalSpecTGoPureGoGoGoPure(r, x)
	})
}

//line runner.go:1842
func EvalSpecTGoPureGoGoGoPure(r []*tree_result.Tree, r2 struct{}) *ProxyEventArrayTreeStringVoidResult {
	return PureProxyEventArrayTreeStringVoidResult(r)
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
func EvalSpecTGo2(p *ProxyEventArrayTreeStringVoidResult) *ProxyX {
	for {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
		switch p.Tag {
		case ProxyEventArrayTreeStringVoidResultRequest:
			x_ := p.Request0
			fx := p.Request1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
			return RequestProxyX(x_, func(x struct{}) *ProxyX {
//line ../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
				return EvalSpecTGo2(fx(x))
			})
		case ProxyEventArrayTreeStringVoidResultRespond:
			fb_ := p.Respond1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:139:27
			p = fb_(struct{}{})
			continue
		case ProxyEventArrayTreeStringVoidResultM:
			m := p.M
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
			return MProxyX(func() *ProxyX {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:140:30
				proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:140:30
				return EvalSpecTGo2(proxy)
			})
		default: // ProxyEventArrayTreeStringVoidResultPure
			a := p.Pure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
			return PureProxyX(a)
		}
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:95:12
func RunEffectGo(p *ProxyX) []*tree_result.Tree {
	for {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:95:12
		switch p.Tag {
		case ProxyXRequest, ProxyXRespond:
			v := p.Field0
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:96:24
			return Closed(v)
		case ProxyXM:
			m := p.M
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:98:24
			p = m()
			continue
		default: // ProxyXPure
			return p.Pure
		}
	}
}

//line ../../../../../../p/aff-8.0.0/src/Effect/Aff.purs:103:7
func RunEffectRecGo(a *ProxyX) []*tree_result.Tree {
	for {
//line ../../../../../../p/aff-8.0.0/src/Effect/Aff.purs:103:7
		switch a.Tag {
		case ProxyXRequest, ProxyXRespond:
			v := a.Field0
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:104:31
			return Closed(v)
		case ProxyXPure:
			return a.Pure
		default: // ProxyXM
			mr := a.M
//line ../../../../../../p/aff-8.0.0/src/Effect/Aff.purs:103:7
			a = mr()
			continue
		}
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
func ApplyGo(px *ProxyEventArrayTreeStringVoidResult, pf *ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResult) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
	switch pf.Tag {
	case ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return RequestProxyArrayArrayTreeStringVoidResult(pf.Request0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
			return ApplyGo(px, pf.Request1(x))
		})
	case ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return RespondProxyArrayArrayTreeStringVoidResult(pf.Respond0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:38:42
			return ApplyGo(px, pf.Respond1(x))
		})
	case ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:39:37
			proxy := pf.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:39:49
			return ApplyGo(px, proxy)
		})
	default: // ProxyFnArrayTreeStringVoidResultArrayArrayTreeStringVoidResultPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:27:14
		return ApplyGoGo(pf.Pure, px)
	}
}

//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
func ApplyGoGo(f func([]*tree_result.Tree) [][]*tree_result.Tree, p *ProxyEventArrayTreeStringVoidResult) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
	switch p.Tag {
	case ProxyEventArrayTreeStringVoidResultRequest:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return RequestProxyArrayArrayTreeStringVoidResult(p.Request0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
			return ApplyGoGo(f, p.Request1(x))
		})
	case ProxyEventArrayTreeStringVoidResultRespond:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return RespondProxyArrayArrayTreeStringVoidResult(p.Respond0, func(x struct{}) *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:30:42
			return ApplyGoGo(f, p.Respond1(x))
		})
	case ProxyEventArrayTreeStringVoidResultM:
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:28:12
		return MProxyArrayArrayTreeStringVoidResult(func() *ProxyArrayArrayTreeStringVoidResult {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:37
			proxy := p.M()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:49
			return ApplyGoGo(f, proxy)
		})
	default: // ProxyEventArrayTreeStringVoidResultPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:32:25
		return PureProxyArrayArrayTreeStringVoidResult(f(p.Pure))
	}
}

//line runner.go:1980
const (
	ComposeResponseFbGoStep = iota
	ComposeResponseFbGoGoStep
)

func ComposeResponseFbGoOrComposeResponseAffUnitXArrayTreeStringVoidResultEventUnitXUnitFbGoGo(step int, aVar avar.AVar[event.EitherUnit], p *ProxyEventArrayTreeStringVoidResult, fb_ func(struct{}) *ProxyEventArrayTreeStringVoidResult, p2 *internal.ProxyUnit) *ProxyX {
	for {
		switch step {
		case ComposeResponseFbGoStep:
			aVarCaptured := aVar
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
			switch p.Tag {
			case ProxyEventArrayTreeStringVoidResultRequest:
				x_ := p.Request0
				fx := p.Request1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
				return RequestProxyX(x_, func(x struct{}) *ProxyX {
//line ../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
					return ComposeResponseFbGo(aVarCaptured, fx(x))
				})
			case ProxyEventArrayTreeStringVoidResultRespond:
				b := p.Respond0
				fb_2 := p.Respond1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:46:15
				step, fb_, aVar, p2 = ComposeResponseFbGoGoStep, fb_2, aVarCaptured, core.ComposeResponseFbGoGo2(core.ComposeResponseFbGoGoM(func() *internal.ProxyUnit {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:158:28
					aff.MakeAffDirect(func(value either.EitherUnit) bool {
						return value.Tag == either.EitherUnitLeft
					}, func(value either.EitherUnit) exception.Error {
						return value.Left
					}, func(value either.EitherUnit) struct{} {
						return value.Right
					}, func(k func(either.EitherUnit)) func(exception.Error) {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:53:8
						effect := avar.PutVar(func(payload exception.Error) either.EitherUnit {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:53:8
							return either.LeftEitherUnit(payload)
						}, func(payload struct{}) either.EitherUnit {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:53:8
							return either.RightEitherUnit(payload)
						}, event.RightEitherUnit(b), aVarCaptured, k)
						return func(_ exception.Error) {
//line ../../../../../../p/avar-5.0.1/src/Effect/Aff/AVar.purs:53:8
							effect()
						}
					})
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Runner.purs:158:28
					return internal.PureProxyUnit(struct{}{})
				}))
				continue
			case ProxyEventArrayTreeStringVoidResultM:
				m := p.M
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
				return MProxyX(func() *ProxyX {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:140:30
					proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:140:30
					return ComposeResponseFbGo(aVarCaptured, proxy)
				})
			default: // ProxyEventArrayTreeStringVoidResultPure
				a := p.Pure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Core.purs:137:12
				return PureProxyX(a)
			}
		default: // ComposeResponseFbGoGoStep
			fb_Captured := fb_
			aVarCaptured := aVar
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
			switch p2.Tag {
			case internal.ProxyUnitRequest:
				a_ := p2.Field0
				fa := p2.Field1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
				return RequestProxyX(a_, func(x struct{}) *ProxyX {
					return ComposeResponseFbGoGo(fb_Captured, aVarCaptured, fa(x))
				})
			case internal.ProxyUnitRespond:
				b := p2.Field0
				fb_2 := p2.Field1
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
				return RespondProxyX(b, func(x struct{}) *ProxyX {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:49:42
					return ComposeResponseFbGoGo(fb_Captured, aVarCaptured, fb_2(x))
				})
			case internal.ProxyUnitM:
				m := p2.M
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:47:12
				return MProxyX(func() *ProxyX {
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:37
					proxy := m()
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:50:49
					return ComposeResponseFbGoGo(fb_Captured, aVarCaptured, proxy)
				})
			default: // internal.ProxyUnitPure
//line ../../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:51:25
				step, aVar, p = ComposeResponseFbGoStep, aVarCaptured, fb_Captured(struct{}{})
				continue
			}
		}
	}
}

//line runner.go:2083
func ComposeResponseFbGo(aVar avar.AVar[event.EitherUnit], p *ProxyEventArrayTreeStringVoidResult) *ProxyX {
	return ComposeResponseFbGoOrComposeResponseAffUnitXArrayTreeStringVoidResultEventUnitXUnitFbGoGo(ComposeResponseFbGoStep, aVar, p, nil, nil)
}

func ComposeResponseFbGoGo(fb_ func(struct{}) *ProxyEventArrayTreeStringVoidResult, aVar avar.AVar[event.EitherUnit], p *internal.ProxyUnit) *ProxyX {
	return ComposeResponseFbGoOrComposeResponseAffUnitXArrayTreeStringVoidResultEventUnitXUnitFbGoGo(ComposeResponseFbGoGoStep, aVar, nil, fb_, p)
}
