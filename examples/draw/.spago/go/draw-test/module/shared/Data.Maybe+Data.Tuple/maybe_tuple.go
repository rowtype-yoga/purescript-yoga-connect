package maybe_tuple

import (
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Tuple"
)

type MaybeTupleIntInt struct {
	Just tuple.TupleInt
	Tag  uint8
}

type MaybeMaybeTupleIntInt struct {
	Just MaybeTupleIntInt
	Tag  uint8
}

func JustTupleIntInt(just tuple.TupleInt) MaybeTupleIntInt {
	return MaybeTupleIntInt{Tag: maybe.TagJust, Just: just}
}

var NoTupleIntInt = MaybeTupleIntInt{Tag: maybe.TagNothing}

func JustMaybeTupleIntInt(just MaybeTupleIntInt) MaybeMaybeTupleIntInt {
	return MaybeMaybeTupleIntInt{Tag: maybe.TagJust, Just: just}
}

var NoMaybeTupleIntInt = MaybeMaybeTupleIntInt{Tag: maybe.TagNothing}
