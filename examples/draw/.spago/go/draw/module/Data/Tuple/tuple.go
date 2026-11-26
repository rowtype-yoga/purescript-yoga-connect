package tuple

import context "iris.local/draw/Go_IO_Context_ffi"

type TupleInt struct {
	First  int
	Second int
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
