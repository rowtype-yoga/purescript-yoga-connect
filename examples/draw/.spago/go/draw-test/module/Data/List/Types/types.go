package types

import (
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Tuple"
)

type ListTupleIntInt struct {
	Cons1 *ListTupleIntInt
	Cons0 tuple.TupleInt
}

var SharedListTupleIntIntNil = &ListTupleIntInt{}

type Tuple struct {
	Second *ListTupleIntInt
	First  tuple.TupleInt
}

type MaybeTupleTupleIntIntListTupleIntInt struct {
	Just Tuple
	Tag  uint8
}

var NoTupleTupleIntIntListTupleIntInt = MaybeTupleTupleIntIntListTupleIntInt{Tag: maybe.TagNothing}

func JustTupleTupleIntIntListTupleIntInt(just Tuple) MaybeTupleTupleIntIntListTupleIntInt {
	return MaybeTupleTupleIntIntListTupleIntInt{Tag: maybe.TagJust, Just: just}
}

type ListString struct {
	Cons0 string
	Cons1 *ListString
}

func ConsListTupleIntInt(cons0 tuple.TupleInt, cons1 *ListTupleIntInt) *ListTupleIntInt {
	return &ListTupleIntInt{Cons0: cons0, Cons1: cons1}
}

func ConsListString(cons0 string, cons1 *ListString) *ListString {
	return &ListString{Cons0: cons0, Cons1: cons1}
}
