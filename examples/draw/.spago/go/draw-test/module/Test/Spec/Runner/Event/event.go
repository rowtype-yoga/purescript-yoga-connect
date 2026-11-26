package event

import (
	"github.com/i-am-the-slime/iris-go/exception"
	"iris.local/draw-test/Test/Spec/Result"
	"iris.local/draw-test/Test/Spec/Tree"
	"iris.local/draw-test/shared/Test.Spec.Tree+Test.Spec.Result"
)

type Execution uint8

const (
	ExecutionParallel Execution = iota
	ExecutionSequential
)

type Event struct {
	FieldA1 tree.TupleArrayPathItem
	FieldB0 tree.TupleArrayPathItem
	FieldB1 result.Result
	End     []*tree_result.Tree
	Start   int
	Tag     uint8
	FieldA0 Execution
}

const (
	EventStart = iota
	EventSuite
	EventSuiteEnd
	EventTest
	EventTestEnd
	EventPending
	EventEnd
)

type EitherUnit struct {
	Left  struct{}
	Right *Event
	Tag   uint8
}

const (
	EitherUnitLeft = iota
	EitherUnitRight
)

type EitherError struct {
	Left  exception.Error
	Right EitherUnit
	Tag   uint8
}

const (
	EitherErrorLeft = iota
	EitherErrorRight
)

func Start(start int) *Event {
	return &Event{Tag: EventStart, Start: start}
}

func Test(fieldA0 Execution, fieldA1 tree.TupleArrayPathItem) *Event {
	return &Event{Tag: EventTest, FieldA0: fieldA0, FieldA1: fieldA1}
}

func Pending(fieldB0 tree.TupleArrayPathItem) *Event {
	return &Event{Tag: EventPending, FieldB0: fieldB0}
}

func Suite(fieldA0 Execution, fieldA1 tree.TupleArrayPathItem) *Event {
	return &Event{Tag: EventSuite, FieldA0: fieldA0, FieldA1: fieldA1}
}

func TestEnd(fieldB0 tree.TupleArrayPathItem, fieldB1 result.Result) *Event {
	return &Event{Tag: EventTestEnd, FieldB0: fieldB0, FieldB1: fieldB1}
}

func SuiteEnd(fieldB0 tree.TupleArrayPathItem) *Event {
	return &Event{Tag: EventSuiteEnd, FieldB0: fieldB0}
}

func LeftEitherUnit(left struct{}) EitherUnit {
	return EitherUnit{Tag: EitherUnitLeft, Left: left}
}

func LeftEitherError(left exception.Error) EitherError {
	return EitherError{Tag: EitherErrorLeft, Left: left}
}

func RightEitherError(right EitherUnit) EitherError {
	return EitherError{Tag: EitherErrorRight, Right: right}
}

func End(end []*tree_result.Tree) *Event {
	return &Event{Tag: EventEnd, End: end}
}

func RightEitherUnit(right *Event) EitherUnit {
	return EitherUnit{Tag: EitherUnitRight, Right: right}
}
