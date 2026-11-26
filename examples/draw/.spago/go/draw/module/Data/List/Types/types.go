package types

import "iris.local/draw/Data/Tuple"

type List struct {
	Cons1 *List
	Cons0 tuple.TupleInt
}

func Cons(cons0 tuple.TupleInt, cons1 *List) *List {
	return &List{Cons0: cons0, Cons1: cons1}
}
