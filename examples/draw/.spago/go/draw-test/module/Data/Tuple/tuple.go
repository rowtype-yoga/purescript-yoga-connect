package tuple

import (
	chanffi "iris.local/draw-test/Go_IO_Chan_ffi"
	context "iris.local/draw-test/Go_IO_Context_ffi"
)

type TupleInt struct {
	First  int
	Second int
}

type TupleArrayUnit struct {
	First  []struct{}
	Second int
}

type TupleUnit struct {
	First  struct{}
	Second int
}

type You struct {
	Before []TupleInt
	Outbox chanffi.Chan[TupleInt]
	Id     int
}

type Tuple struct {
	First  any
	Second any
}

type TupleContext struct {
	First  context.Context
	Second context.CancelContext
}

func Snd(tuple TupleInt) int {
	return tuple.Second
}

func Fst(tuple TupleInt) int {
	return tuple.First
}
