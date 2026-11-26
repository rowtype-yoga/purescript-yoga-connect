package set

import (
	"github.com/i-am-the-slime/iris-go/direct"
	list "iris.local/draw/Data/List/Types"
	m "iris.local/draw/Data/Map/Internal"
	"iris.local/draw/Data/Tuple"
)

var SharedListNil = &list.List{}

func Foldr(f func(tuple.TupleInt) func(*direct.ConsList[tuple.TupleInt]) *direct.ConsList[tuple.TupleInt], x *direct.ConsList[tuple.TupleInt]) func(*m.Map) *direct.ConsList[tuple.TupleInt] {
	return func(x2 *m.Map) *direct.ConsList[tuple.TupleInt] {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
		return FoldrGoSet(f, x, FoldrGoSetReusing(SharedListNil, FoldrGoSet2(x2, SharedListNil)))
	}
}

//line set.go:20
func FoldrGoSet(f func(tuple.TupleInt) func(*direct.ConsList[tuple.TupleInt]) *direct.ConsList[tuple.TupleInt], b *direct.ConsList[tuple.TupleInt], value *list.List) *direct.ConsList[tuple.TupleInt] {
	for {
		if value == SharedListNil {
			return b
		}
		a := value.Cons0
		as := value.Cons1
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:113:17
		b, value = f(a)(b), as
	}
}

//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:189:26
func FoldrGoSet2(m_ *m.Map, z_ *list.List) *list.List {
	for {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:189:26
		if m_ == m.SharedMapLeaf {
			return z_
		}
		k := m_.Node2
		l := m_.Node4
		r := m_.Node5
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:192:9
		m_, z_ = l, list.Cons(k, FoldrGoSet2(r, z_))
	}
}

//line set.go:48
func FoldrGoSetReusing(acc *list.List, list *list.List) *list.List {
	for {
		if list == SharedListNil {
			return acc
		}
		xs := list.Cons1
		list.Cons1 = acc
		acc, list = list, xs
	}
}
