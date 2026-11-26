package maybe

import bytes "iris.local/draw-test/Yoga_Protobuf_Bytes_ffi"

const TagNothing = 0

const TagJust = 1

type String struct {
	Just string
	Tag  uint8
}

type Boolean uint8

const (
	False Boolean = iota
	True
	NoBoolean
)

type ExampleIsFocusedIsParallelizable struct {
	Example          func(func(func(any) any) any) any
	IsFocused        bool
	IsParallelizable Boolean
}

type MaybeExampleIsFocusedIsParallelizable struct {
	Just   ExampleIsFocusedIsParallelizable
	Origin any
	Tag    uint8
}

type Number struct {
	Just float64
	Tag  uint8
}

func JustNumber(just float64) Number {
	return Number{Tag: TagJust, Just: just}
}

type Int struct {
	Just int
	Tag  uint8
}

var NoInt = Int{Tag: TagNothing}

func JustInt(just int) Int {
	return Int{Tag: TagJust, Just: just}
}

type Maybe struct {
	Just any
	Tag  uint8
}

type BytesName struct {
	Bytes bytes.Bytes
	Name  string
}

type MaybeBytesName struct {
	Just BytesName
	Tag  uint8
}

func JustString(just string) String {
	return String{Tag: TagJust, Just: just}
}

var NoString = String{Tag: TagNothing}
