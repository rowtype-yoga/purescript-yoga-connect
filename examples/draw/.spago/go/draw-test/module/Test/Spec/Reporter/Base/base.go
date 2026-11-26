package base

import (
	"github.com/i-am-the-slime/iris-go/alloc"
	"iris.local/draw-test/Data/Maybe"
	specresult "iris.local/draw-test/Test/Spec/Result"
	"iris.local/draw-test/Test/Spec/Runner/Event"
	s "iris.local/draw-test/Test/Spec/Tree"
)

type RunningItem struct {
	RunningTest  specresult.MaybeResult
	Tag          uint8
	RunningSuite bool
}

const (
	RunningItemRunningTest = iota
	RunningItemRunningPending
	RunningItemRunningSuite
)

type Map struct {
	Node2 s.TupleArrayPathItem
	Node3 RunningItem
	Node4 *Map
	Node5 *Map
	Node0 int
	Node1 int
}

var SharedMapLeaf = &Map{}

type State struct {
	LastPrintedSuitePath s.MaybeArrayPathItem
	RunningItems         *Map
}

type Proxy struct {
	Request0 struct{}
	Request1 func(*event.Event) Proxy
	Respond0 *event.Event
	Respond1 func(struct{}) Proxy
	M        func() Proxy
	Pure     State
	Tag      uint8
}

const (
	ProxyRequest = iota
	ProxyRespond
	ProxyM
	ProxyPure
)

type TupleUnit struct {
	First  struct{}
	Second State
}

type TupleTupleUnitState struct {
	First  TupleUnit
	Second string
}

type MapIter struct {
	IterEmit0 s.TupleArrayPathItem
	IterEmit1 RunningItem
	IterEmit2 *MapIter
	IterNode0 *Map
	IterNode1 *MapIter
	Tag       uint8
}

const (
	MapIterIterLeaf = iota
	MapIterIterEmit
	MapIterIterNode
)

type TupleTupleArrayPathItemString struct {
	First  s.TupleArrayPathItem
	Second RunningItem
}

type TupleTupleTupleArrayPathItemStringRunningItem struct {
	First  TupleTupleArrayPathItemString
	Second *MapIter
}

type MaybeTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem struct {
	Just TupleTupleTupleArrayPathItemStringRunningItem
	Tag  uint8
}

type SplitLast struct {
	First  s.TupleArrayPathItem
	Second RunningItem
	Third  *Map
}

type MaybeRunningItem struct {
	Just RunningItem
	Tag  uint8
}

var NoRunningItem = MaybeRunningItem{Tag: maybe.TagNothing}

func JustRunningItem(just RunningItem) MaybeRunningItem {
	return MaybeRunningItem{Tag: maybe.TagJust, Just: just}
}

var NoTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem = MaybeTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem{Tag: maybe.TagNothing}

func JustTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem(just TupleTupleTupleArrayPathItemStringRunningItem) MaybeTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem {
	return MaybeTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem{Tag: maybe.TagJust, Just: just}
}

func Pure(pure State) Proxy {
	return Proxy{Tag: ProxyPure, Pure: pure}
}

func Node(node0 int, node1 int, node2 s.TupleArrayPathItem, node3 RunningItem, node4 *Map, node5 *Map) *Map {
	return &Map{Node0: node0, Node1: node1, Node2: node2, Node3: node3, Node4: node4, Node5: node5}
}

func RunningSuite(runningSuite bool) RunningItem {
	return RunningItem{Tag: RunningItemRunningSuite, RunningSuite: runningSuite}
}

func RunningTest(runningTest specresult.MaybeResult) RunningItem {
	return RunningItem{Tag: RunningItemRunningTest, RunningTest: runningTest}
}

func IterNode(iterNode0 *Map, iterNode1 *MapIter) *MapIter {
	return &MapIter{Tag: MapIterIterNode, IterNode0: iterNode0, IterNode1: iterNode1}
}

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:717:40
func UnsafeBalancedNode(k s.TupleArrayPathItem, v RunningItem, l *Map, r *Map) *Map {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:717:40
	if l == SharedMapLeaf {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:719:5
		if r == SharedMapLeaf {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:719:5
			return Node(1, 1, k, v, SharedMapLeaf, SharedMapLeaf)
		}
		if r.Node0 > 1 {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:724:13
			return UnsafeBalancedNodeRotateLeft(k, v, l, r.Node2, r.Node3, r.Node4, r.Node5)
		}
		return UnsafeNode(k, v, l, r)
	}
	lh := l.Node0
	lk := l.Node2
	lv := l.Node3
	ll := l.Node4
	lr := l.Node5
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:728:5
	if r != SharedMapLeaf {
		rh := r.Node0
		if rh > lh+1 {
			return UnsafeBalancedNodeRotateLeft(k, v, l, r.Node2, r.Node3, r.Node4, r.Node5)
		}
		if lh > rh+1 {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:733:13
			return UnsafeBalancedNodeRotateRight(k, v, lk, lv, ll, lr, r)
		}
	}
	if r == SharedMapLeaf {
		if lh > 1 {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:736:13
			return UnsafeBalancedNodeRotateRight(k, v, lk, lv, ll, lr, r)
		}
	}
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:738:9
	return UnsafeNode(k, v, l, r)
}

//line base.go:181
func UnsafeNode(k s.TupleArrayPathItem, v RunningItem, l *Map, r *Map) *Map {
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

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:776:37
func UnsafeSplitLast(k s.TupleArrayPathItem, v RunningItem, l *Map, r *Map) SplitLast {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:776:37
	if r == SharedMapLeaf {
		return SplitLast{First: k, Second: v, Third: l}
	}
	scrutinee := UnsafeSplitLast(r.Node2, r.Node3, r.Node4, r.Node5)
	return SplitLast{First: scrutinee.First, Second: scrutinee.Second, Third: UnsafeBalancedNode(k, v, l, scrutinee.Third)}
}

//line base.go:209
var IterMapL func(*MapIter, *Map) *MapIter = IterMapLGo

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:940:3
func StepUnfoldrGo(value *MapIter) MaybeTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem {
	for {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:940:3
		switch value.Tag {
		case MapIterIterLeaf:
			return NoTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem
		case MapIterIterEmit:
			k := value.IterEmit0
			v := value.IterEmit1
			iter := value.IterEmit2
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:944:7
			return JustTupleTupleTupleArrayPathItemStringRunningItemMapIterTupleArrayPathItemStringRunningItem(TupleTupleTupleArrayPathItemStringRunningItem{First: TupleTupleArrayPathItemString{First: k, Second: v}, Second: iter})
		default: // MapIterIterNode
			m := value.IterNode0
			iter := value.IterNode1
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:946:7
			value = IterMapL(iter, m)
			continue
		}
	}
}

//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:741:44
func UnsafeBalancedNodeRotateLeft(k s.TupleArrayPathItem, v RunningItem, l *Map, rk s.TupleArrayPathItem, rv RunningItem, rl *Map, rr *Map) *Map {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:741:44
	if rl != SharedMapLeaf {
		var result int
		if rr == SharedMapLeaf {
			result = 0
		} else {
			result = rr.Node0
		}
		if rl.Node0 > result {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:744:11
			return UnsafeNode(rl.Node2, rl.Node3, UnsafeNode(k, v, l, rl.Node4), UnsafeNode(rk, rv, rl.Node5, rr))
		}
	}
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:746:7
	return UnsafeNode(rk, rv, UnsafeNode(k, v, l, rl), rr)
}

func UnsafeBalancedNodeRotateRight(k s.TupleArrayPathItem, v RunningItem, lk s.TupleArrayPathItem, lv RunningItem, ll *Map, lr *Map, r *Map) *Map {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:749:45
	if lr != SharedMapLeaf {
		var result int
		if ll == SharedMapLeaf {
			result = 0
		} else {
			result = ll.Node0
		}
		if result <= lr.Node0 {
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:752:11
			return UnsafeNode(lr.Node2, lr.Node3, UnsafeNode(lk, lv, ll, lr.Node4), UnsafeNode(k, v, lr.Node5, r))
		}
	}
//line ../../../../../../../p/ordered-collections-3.2.0/src/Data/Map/Internal.purs:754:7
	return UnsafeNode(lk, lv, ll, UnsafeNode(k, v, lr, r))
}

//line base.go:272
func IterMapLGo(iter *MapIter, value *Map) *MapIter {
	var mapIterTupleArrayRecordIndexNameStringRunningItemCells alloc.Cells[MapIter]
	for {
		if value == SharedMapLeaf {
			return iter
		}
		k := value.Node2
		v := value.Node3
		l := value.Node4
		r := value.Node5
		if r == SharedMapLeaf {
			iter, value = mapIterTupleArrayRecordIndexNameStringRunningItemCells.New(MapIter{Tag: MapIterIterEmit, IterEmit0: k, IterEmit1: v, IterEmit2: iter}), l
			continue
		}
		iter, value = mapIterTupleArrayRecordIndexNameStringRunningItemCells.New(MapIter{Tag: MapIterIterEmit, IterEmit0: k, IterEmit1: v, IterEmit2: mapIterTupleArrayRecordIndexNameStringRunningItemCells.New(MapIter{Tag: MapIterIterNode, IterNode0: r, IterNode1: iter})}), l
	}
}
