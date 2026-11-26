package set

import (
	"github.com/i-am-the-slime/iris-go/direct"
	list "iris.local/draw-test/Data/List/Types"
	m "iris.local/draw-test/Data/Map/Internal"
	"iris.local/draw-test/Data/Tuple"
)

func Foldr(f func(tuple.TupleInt) func(*direct.ConsList[tuple.TupleInt]) *direct.ConsList[tuple.TupleInt], x *direct.ConsList[tuple.TupleInt]) func(*m.MapTupleIntInt) *direct.ConsList[tuple.TupleInt] {
	return func(x2 *m.MapTupleIntInt) *direct.ConsList[tuple.TupleInt] {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
		return FoldrGoSet(f, x, FoldrGoSetReusing(list.SharedListTupleIntIntNil, DrawGo(x2, list.SharedListTupleIntIntNil)))
	}
}

//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:189:26
func DrawGo(m_ *m.MapTupleIntInt, z_ *list.ListTupleIntInt) *list.ListTupleIntInt {
	for {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:189:26
		if m_ == m.SharedMapTupleIntIntLeaf {
			return z_
		}
		k := m_.Node2
		l := m_.Node4
		r := m_.Node5
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:192:9
		m_, z_ = l, list.ConsListTupleIntInt(k, DrawGo(r, z_))
	}
}

//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
func InsertGo(a tuple.TupleInt, value *m.MapTupleIntInt) *m.MapTupleIntInt {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
	if value == m.SharedMapTupleIntIntLeaf {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:468:3
		return m.NodeMapTupleIntInt(1, 1, a, struct{}{}, m.SharedMapTupleIntIntLeaf, m.SharedMapTupleIntIntLeaf)
	}
	mk := value.Node2
	ml := value.Node4
	mr := value.Node5
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
	left0 := a.First
	left1 := a.Second
	right0 := mk.First
	right1 := mk.Second
	if left0 < right0 {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:472:15
		return m.UnsafeBalancedNodeTupleIntInt(mk, struct{}{}, InsertGo(a, ml), mr)
	}
	if left0 == right0 {
		if left1 < right1 {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:472:15
			return m.UnsafeBalancedNodeTupleIntInt(mk, struct{}{}, InsertGo(a, ml), mr)
		}
		if left1 == right1 {
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:471:7
			return m.NodeMapTupleIntInt(value.Node0, value.Node1, a, struct{}{}, ml, mr)
		}
		return m.UnsafeBalancedNodeTupleIntInt(mk, struct{}{}, ml, InsertGo(a, mr))
	}
//line ../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:473:15
	return m.UnsafeBalancedNodeTupleIntInt(mk, struct{}{}, ml, InsertGo(a, mr))
}

//line set.go:67
func FoldrGoSet(f func(tuple.TupleInt) func(*direct.ConsList[tuple.TupleInt]) *direct.ConsList[tuple.TupleInt], b *direct.ConsList[tuple.TupleInt], value *list.ListTupleIntInt) *direct.ConsList[tuple.TupleInt] {
	for {
		if value == list.SharedListTupleIntIntNil {
			return b
		}
		a := value.Cons0
		as := value.Cons1
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:113:17
		b, value = f(a)(b), as
	}
}

//line set.go:80
func FoldrGoSetReusing(acc *list.ListTupleIntInt, list2 *list.ListTupleIntInt) *list.ListTupleIntInt {
	for {
		if list2 == list.SharedListTupleIntIntNil {
			return acc
		}
		xs := list2.Cons1
		list2.Cons1 = acc
		acc, list2 = list2, xs
	}
}
