package internal

import (
	"iris.local/draw-test/Data/Tuple"
	chanffi "iris.local/draw-test/Go_IO_Chan_ffi"
	ioref "iris.local/draw-test/Go_IO_IORef_ffi"
)

type MapTupleIntInt struct {
	Node3 struct{}
	Node4 *MapTupleIntInt
	Node5 *MapTupleIntInt
	Node0 int
	Node1 int
	Node2 tuple.TupleInt
}

var SharedMapTupleIntIntLeaf = &MapTupleIntInt{}

type MapInt struct {
	Node3 chanffi.Chan[tuple.TupleInt]
	Node4 *MapInt
	Node5 *MapInt
	Node0 int
	Node1 int
	Node2 int
}

var SharedMapIntLeaf = &MapInt{}

type Canvas struct {
	Cells    *MapTupleIntInt
	Outboxes *MapInt
	NextId   int
}

type Room struct {
	Room ioref.IORef[Canvas]
}

type TupleYou struct {
	First  Canvas
	Second tuple.You
}

type TupleUnit struct {
	Second struct{}
	First  Canvas
}

type TupleCanvas struct {
	First  Canvas
	Second Canvas
}

type SplitLast struct {
	Second chanffi.Chan[tuple.TupleInt]
	Third  *MapInt
	First  int
}

func NodeMapTupleIntInt(node0 int, node1 int, node2 tuple.TupleInt, node3 struct{}, node4 *MapTupleIntInt, node5 *MapTupleIntInt) *MapTupleIntInt {
	return &MapTupleIntInt{Node0: node0, Node1: node1, Node2: node2, Node3: node3, Node4: node4, Node5: node5}
}

func NodeMapInt(node0 int, node1 int, node2 int, node3 chanffi.Chan[tuple.TupleInt], node4 *MapInt, node5 *MapInt) *MapInt {
	return &MapInt{Node0: node0, Node1: node1, Node2: node2, Node3: node3, Node4: node4, Node5: node5}
}

//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:717:40
func UnsafeBalancedNodeTupleIntInt(k tuple.TupleInt, v struct{}, l *MapTupleIntInt, r *MapTupleIntInt) *MapTupleIntInt {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:717:40
	if l == SharedMapTupleIntIntLeaf {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:719:5
		if r == SharedMapTupleIntIntLeaf {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:719:5
			return NodeMapTupleIntInt(1, 1, k, v, SharedMapTupleIntIntLeaf, SharedMapTupleIntIntLeaf)
		}
		if r.Node0 > 1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:724:13
			return UnsafeBalancedNodeRotateLeftTupleIntInt(k, v, l, r.Node2, r.Node3, r.Node4, r.Node5)
		}
		return UnsafeNodeTupleIntInt(k, v, l, r)
	}
	lh := l.Node0
	lk := l.Node2
	ll := l.Node4
	lr := l.Node5
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:728:5
	if r != SharedMapTupleIntIntLeaf {
		rh := r.Node0
		if rh > lh+1 {
			return UnsafeBalancedNodeRotateLeftTupleIntInt(k, v, l, r.Node2, r.Node3, r.Node4, r.Node5)
		}
		if lh > rh+1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:733:13
			return UnsafeBalancedNodeRotateRightTupleIntInt(k, v, lk, struct{}{}, ll, lr, r)
		}
	}
	if r == SharedMapTupleIntIntLeaf {
		if lh > 1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:736:13
			return UnsafeBalancedNodeRotateRightTupleIntInt(k, v, lk, struct{}{}, ll, lr, r)
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:738:9
	return UnsafeNodeTupleIntInt(k, v, l, r)
}

//line internal.go:111
func FstCanvas(tuple TupleCanvas) Canvas {
	return tuple.First
}

func SndCanvas(tuple TupleCanvas) Canvas {
	return tuple.Second
}

func FstYou(tuple TupleYou) Canvas {
	return tuple.First
}

func SndYou(tuple TupleYou) tuple.You {
	return tuple.Second
}

//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:717:40
func UnsafeBalancedNodeInt(k int, v chanffi.Chan[tuple.TupleInt], l *MapInt, r *MapInt) *MapInt {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:717:40
	if l == SharedMapIntLeaf {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:719:5
		if r == SharedMapIntLeaf {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:719:5
			return NodeMapInt(1, 1, k, v, SharedMapIntLeaf, SharedMapIntLeaf)
		}
		if r.Node0 > 1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:724:13
			return UnsafeBalancedNodeRotateLeftInt(k, v, l, r.Node2, r.Node3, r.Node4, r.Node5)
		}
		return UnsafeNodeInt(k, v, l, r)
	}
	lh := l.Node0
	lk := l.Node2
	lv := l.Node3
	ll := l.Node4
	lr := l.Node5
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:728:5
	if r != SharedMapIntLeaf {
		rh := r.Node0
		if rh > lh+1 {
			return UnsafeBalancedNodeRotateLeftInt(k, v, l, r.Node2, r.Node3, r.Node4, r.Node5)
		}
		if lh > rh+1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:733:13
			return UnsafeBalancedNodeRotateRightInt(k, v, lk, lv, ll, lr, r)
		}
	}
	if r == SharedMapIntLeaf {
		if lh > 1 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:736:13
			return UnsafeBalancedNodeRotateRightInt(k, v, lk, lv, ll, lr, r)
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:738:9
	return UnsafeNodeInt(k, v, l, r)
}

//line internal.go:169
func FstUnit(tuple TupleUnit) Canvas {
	return tuple.First
}

func SndUnit(tuple TupleUnit) {
}

func UnsafeNodeTupleIntInt(k tuple.TupleInt, v struct{}, l *MapTupleIntInt, r *MapTupleIntInt) *MapTupleIntInt {
	if l == SharedMapTupleIntIntLeaf {
		if r == SharedMapTupleIntIntLeaf {
			return NodeMapTupleIntInt(1, 1, k, v, l, r)
		}
		return NodeMapTupleIntInt(1+r.Node0, 1+r.Node1, k, v, l, r)
	}
	h1 := l.Node0
	s1 := l.Node1
	if r == SharedMapTupleIntIntLeaf {
		return NodeMapTupleIntInt(1+h1, 1+s1, k, v, l, r)
	}
	h2 := r.Node0
	result := max(h1, h2)
	return NodeMapTupleIntInt(1+result, 1+s1+r.Node1, k, v, l, r)
}

func UnsafeNodeInt(k int, v chanffi.Chan[tuple.TupleInt], l *MapInt, r *MapInt) *MapInt {
	if l == SharedMapIntLeaf {
		if r == SharedMapIntLeaf {
			return NodeMapInt(1, 1, k, v, l, r)
		}
		return NodeMapInt(1+r.Node0, 1+r.Node1, k, v, l, r)
	}
	h1 := l.Node0
	s1 := l.Node1
	if r == SharedMapIntLeaf {
		return NodeMapInt(1+h1, 1+s1, k, v, l, r)
	}
	h2 := r.Node0
	result := max(h1, h2)
	return NodeMapInt(1+result, 1+s1+r.Node1, k, v, l, r)
}

//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:776:37
func UnsafeSplitLast(k int, v chanffi.Chan[tuple.TupleInt], l *MapInt, r *MapInt) SplitLast {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:776:37
	if r == SharedMapIntLeaf {
		return SplitLast{First: k, Second: v, Third: l}
	}
	scrutinee := UnsafeSplitLast(r.Node2, r.Node3, r.Node4, r.Node5)
	return SplitLast{First: scrutinee.First, Second: scrutinee.Second, Third: UnsafeBalancedNodeInt(k, v, l, scrutinee.Third)}
}

//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:741:44
func UnsafeBalancedNodeRotateLeftTupleIntInt(k tuple.TupleInt, v struct{}, l *MapTupleIntInt, rk tuple.TupleInt, rv struct{}, rl *MapTupleIntInt, rr *MapTupleIntInt) *MapTupleIntInt {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:741:44
	if rl != SharedMapTupleIntIntLeaf {
		var result int
		if rr == SharedMapTupleIntIntLeaf {
			result = 0
		} else {
			result = rr.Node0
		}
		if rl.Node0 > result {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:744:11
			return UnsafeNodeTupleIntInt(rl.Node2, rl.Node3, UnsafeNodeTupleIntInt(k, v, l, rl.Node4), UnsafeNodeTupleIntInt(rk, rv, rl.Node5, rr))
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:746:7
	return UnsafeNodeTupleIntInt(rk, rv, UnsafeNodeTupleIntInt(k, v, l, rl), rr)
}

func UnsafeBalancedNodeRotateRightTupleIntInt(k tuple.TupleInt, v struct{}, lk tuple.TupleInt, lv struct{}, ll *MapTupleIntInt, lr *MapTupleIntInt, r *MapTupleIntInt) *MapTupleIntInt {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:749:45
	if lr != SharedMapTupleIntIntLeaf {
		var result int
		if ll == SharedMapTupleIntIntLeaf {
			result = 0
		} else {
			result = ll.Node0
		}
		if result <= lr.Node0 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:752:11
			return UnsafeNodeTupleIntInt(lr.Node2, lr.Node3, UnsafeNodeTupleIntInt(lk, lv, ll, lr.Node4), UnsafeNodeTupleIntInt(k, v, lr.Node5, r))
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:754:7
	return UnsafeNodeTupleIntInt(lk, lv, ll, UnsafeNodeTupleIntInt(k, v, lr, r))
}

//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:741:44
func UnsafeBalancedNodeRotateLeftInt(k int, v chanffi.Chan[tuple.TupleInt], l *MapInt, rk int, rv chanffi.Chan[tuple.TupleInt], rl *MapInt, rr *MapInt) *MapInt {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:741:44
	if rl != SharedMapIntLeaf {
		var result int
		if rr == SharedMapIntLeaf {
			result = 0
		} else {
			result = rr.Node0
		}
		if rl.Node0 > result {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:744:11
			return UnsafeNodeInt(rl.Node2, rl.Node3, UnsafeNodeInt(k, v, l, rl.Node4), UnsafeNodeInt(rk, rv, rl.Node5, rr))
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:746:7
	return UnsafeNodeInt(rk, rv, UnsafeNodeInt(k, v, l, rl), rr)
}

func UnsafeBalancedNodeRotateRightInt(k int, v chanffi.Chan[tuple.TupleInt], lk int, lv chanffi.Chan[tuple.TupleInt], ll *MapInt, lr *MapInt, r *MapInt) *MapInt {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:749:45
	if lr != SharedMapIntLeaf {
		var result int
		if ll == SharedMapIntLeaf {
			result = 0
		} else {
			result = ll.Node0
		}
		if result <= lr.Node0 {
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:752:11
			return UnsafeNodeInt(lr.Node2, lr.Node3, UnsafeNodeInt(lk, lv, ll, lr.Node4), UnsafeNodeInt(k, v, lr.Node5, r))
		}
	}
//line ../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:754:7
	return UnsafeNodeInt(lk, lv, ll, UnsafeNodeInt(k, v, lr, r))
}
