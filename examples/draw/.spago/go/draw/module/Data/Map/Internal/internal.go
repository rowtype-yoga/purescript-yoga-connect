package internal

import "iris.local/draw/Data/Tuple"

type Map struct {
	Node3 struct{}
	Node4 *Map
	Node5 *Map
	Node0 int
	Node1 int
	Node2 tuple.TupleInt
}

var SharedMapLeaf = &Map{}

func Node(node0 int, node1 int, node2 tuple.TupleInt, node3 struct{}, node4 *Map, node5 *Map) *Map {
	return &Map{Node0: node0, Node1: node1, Node2: node2, Node3: node3, Node4: node4, Node5: node5}
}

//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:717:40
func UnsafeBalancedNode(k tuple.TupleInt, v struct{}, l *Map, r *Map) *Map {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:717:40
	if l == SharedMapLeaf {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:719:5
		if r == SharedMapLeaf {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:719:5
			return Node(1, 1, k, v, SharedMapLeaf, SharedMapLeaf)
		}
		if r.Node0 > 1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:724:13
			return UnsafeBalancedNodeRotateLeft(k, v, l, r.Node2, r.Node3, r.Node4, r.Node5)
		}
		return UnsafeNode(k, v, l, r)
	}
	lh := l.Node0
	lk := l.Node2
	ll := l.Node4
	lr := l.Node5
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:728:5
	if r != SharedMapLeaf {
		rh := r.Node0
		if rh > lh+1 {
			return UnsafeBalancedNodeRotateLeft(k, v, l, r.Node2, r.Node3, r.Node4, r.Node5)
		}
		if lh > rh+1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:733:13
			return UnsafeBalancedNodeRotateRight(k, v, lk, struct{}{}, ll, lr, r)
		}
	}
	if r == SharedMapLeaf {
		if lh > 1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:736:13
			return UnsafeBalancedNodeRotateRight(k, v, lk, struct{}{}, ll, lr, r)
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:738:9
	return UnsafeNode(k, v, l, r)
}

//line internal.go:61
func UnsafeNode(k tuple.TupleInt, v struct{}, l *Map, r *Map) *Map {
	if l == SharedMapLeaf {
		if r == SharedMapLeaf {
			return Node(1, 1, k, v, l, r)
		}
		return Node(1+r.Node0, 1+r.Node1, k, v, l, r)
	}
	h1 := l.Node0
	s1 := l.Node1
	if r == SharedMapLeaf {
		return Node(1+h1, 1+s1, k, v, l, r)
	}
	h2 := r.Node0
	result := max(h1, h2)
	return Node(1+result, 1+s1+r.Node1, k, v, l, r)
}

//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:741:44
func UnsafeBalancedNodeRotateLeft(k tuple.TupleInt, v struct{}, l *Map, rk tuple.TupleInt, rv struct{}, rl *Map, rr *Map) *Map {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:741:44
	if rl != SharedMapLeaf {
		var result int
		if rr == SharedMapLeaf {
			result = 0
		} else {
			result = rr.Node0
		}
		if rl.Node0 > result {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:744:11
			return UnsafeNode(rl.Node2, rl.Node3, UnsafeNode(k, v, l, rl.Node4), UnsafeNode(rk, rv, rl.Node5, rr))
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:746:7
	return UnsafeNode(rk, rv, UnsafeNode(k, v, l, rl), rr)
}

func UnsafeBalancedNodeRotateRight(k tuple.TupleInt, v struct{}, lk tuple.TupleInt, lv struct{}, ll *Map, lr *Map, r *Map) *Map {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:749:45
	if lr != SharedMapLeaf {
		var result int
		if ll == SharedMapLeaf {
			result = 0
		} else {
			result = ll.Node0
		}
		if result <= lr.Node0 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:752:11
			return UnsafeNode(lr.Node2, lr.Node3, UnsafeNode(lk, lv, ll, lr.Node4), UnsafeNode(k, v, lr.Node5, r))
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:754:7
	return UnsafeNode(lk, lv, ll, UnsafeNode(k, v, lr, r))
}
