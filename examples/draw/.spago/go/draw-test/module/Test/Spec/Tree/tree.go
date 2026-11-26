package tree

import (
	"github.com/i-am-the-slime/iris-go/array"
	"github.com/i-am-the-slime/iris-go/direct"
	"github.com/i-am-the-slime/iris-go/prim"
	"iris.local/draw-test/Data/Either"
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Data/Ordering"
)

type PathItem struct {
	Name  maybe.String
	Index int
}

type TupleArrayPathItem struct {
	First  []PathItem
	Second string
}

type Item struct {
	Example          func(func(func(struct{})))
	IsFocused        bool
	IsParallelizable maybe.Boolean
}

type MaybeItem struct {
	Just Item
	Tag  uint8
}

type TreeStringFnUnitAffUnit struct {
	Node0 either.EitherFnUnitAffUnit
	Node1 []*TreeStringFnUnitAffUnit
	Leaf0 string
	Leaf1 MaybeItem
	Tag   uint8
}

const (
	TreeStringFnUnitAffUnitNode = iota
	TreeStringFnUnitAffUnitLeaf
)

func JustItem(just Item) MaybeItem {
	return MaybeItem{Tag: maybe.TagJust, Just: just}
}

type TupleUnit struct {
	First  struct{}
	Second []*TreeStringFnUnitAffUnit
}

type TreeFnUnit struct {
	Node0  *either.EitherFnUnit
	Node1  []*TreeFnUnit
	Leaf0  string
	Leaf1  *maybe.MaybeExampleIsFocusedIsParallelizable
	Origin any
	Tag    uint8
}

const (
	TreeFnUnitNode = iota
	TreeFnUnitLeaf
)

type Config struct {
	FilterTree func([]*TreeFnUnit) []*TreeFnUnit
	Slow       float64
	Timeout    maybe.Number
	Exit       bool
	FailFast   bool
}

type MaybeArrayPathItem struct {
	Just []PathItem
	Tag  uint8
}

var NoArrayPathItem = MaybeArrayPathItem{Tag: maybe.TagNothing}

func JustArrayPathItem(just []PathItem) MaybeArrayPathItem {
	return MaybeArrayPathItem{Tag: maybe.TagJust, Just: just}
}

type Either struct {
	Left  TupleArrayPathItem
	Right func(struct{})
	Tag   uint8
}

const (
	EitherLeft = iota
	EitherRight
)

type TreeTupleArrayPathItemString struct {
	Node0 Either
	Node1 []*TreeTupleArrayPathItemString
	Leaf0 TupleArrayPathItem
	Leaf1 MaybeItem
	Tag   uint8
}

const (
	TreeTupleArrayPathItemStringNode = iota
	TreeTupleArrayPathItemStringLeaf
)

var NoItem = MaybeItem{Tag: maybe.TagNothing}

func LeafTreeStringFnUnitAffUnit(leaf0 string, leaf1 MaybeItem) *TreeStringFnUnitAffUnit {
	return &TreeStringFnUnitAffUnit{Tag: TreeStringFnUnitAffUnitLeaf, Leaf0: leaf0, Leaf1: leaf1}
}

func NodeTreeStringFnUnitAffUnit(node0 either.EitherFnUnitAffUnit, node1 []*TreeStringFnUnitAffUnit) *TreeStringFnUnitAffUnit {
	return &TreeStringFnUnitAffUnit{Tag: TreeStringFnUnitAffUnitNode, Node0: node0, Node1: node1}
}

func Left(left TupleArrayPathItem) Either {
	return Either{Tag: EitherLeft, Left: left}
}

func Right(right func(struct{})) Either {
	return Either{Tag: EitherRight, Right: right}
}

func NodeTreeTupleArrayPathItemString(node0 Either, node1 []*TreeTupleArrayPathItemString) *TreeTupleArrayPathItemString {
	return &TreeTupleArrayPathItemString{Tag: TreeTupleArrayPathItemStringNode, Node0: node0, Node1: node1}
}

func LeafTreeTupleArrayPathItemString(leaf0 TupleArrayPathItem, leaf1 MaybeItem) *TreeTupleArrayPathItemString {
	return &TreeTupleArrayPathItemString{Tag: TreeTupleArrayPathItemStringLeaf, Leaf0: leaf0, Leaf1: leaf1}
}

func Identity(x []*TreeFnUnit) []*TreeFnUnit {
	return x
}

func Compare(left TupleArrayPathItem, right TupleArrayPathItem) ordering.Ordering {
	left1 := left.Second
	right1 := right.Second
//line ../../../../../../p/prelude-6.0.2/src/Data/Ord.purs:76:34
	i3 := direct.CompareArray(func(x PathItem, y PathItem) int {
//line ../../../../../../p/prelude-6.0.2/src/Data/Ord.purs:79:7
		i := x.Index
		i2 := y.Index
		var left2 ordering.Ordering
		if i < i2 {
			left2 = ordering.OrderingLT
		} else if i == i2 {
			left2 = ordering.OrderingEQ
		} else {
			left2 = ordering.OrderingGT
		}
//line ../../../../../../p/prelude-6.0.2/src/Data/Ord.purs:252:5
		var result ordering.Ordering
		if left2 != ordering.OrderingEQ {
			result = left2
		} else {
//line ../../../../../../p/prelude-6.0.2/src/Data/Ord.purs:257:12
			var left3 ordering.Ordering
			if x.Name.Tag == maybe.TagNothing && y.Name.Tag == maybe.TagNothing {
				left3 = ordering.OrderingEQ
			} else if x.Name.Tag == maybe.TagNothing {
				left3 = ordering.OrderingLT
			} else if y.Name.Tag == maybe.TagNothing {
				left3 = ordering.OrderingGT
			} else {
				left0 := x.Name.Just
				right0 := y.Name.Just
//line ../../../../../../p/prelude-6.0.2/src/Data/Ord.purs:257:12
				if prim.CompareString(left0, right0) < 0 {
					left3 = ordering.OrderingLT
				} else if left0 == right0 {
					left3 = ordering.OrderingEQ
				} else {
					left3 = ordering.OrderingGT
				}
			}
			if left3 != ordering.OrderingEQ {
				result = left3
			} else {
				result = ordering.OrderingEQ
			}
		}
		switch result {
		case ordering.OrderingEQ:
			return 0
		case ordering.OrderingLT:
			return 1
		default: // ordering.OrderingGT
			return -1
		}
	}, left.First, right.First)
	if 0 < i3 {
		return ordering.OrderingLT
	}
	if 0 == i3 {
//line ../../../../../../p/prelude-6.0.2/src/Data/Ord.purs:76:34
		if prim.CompareString(left1, right1) < 0 {
			return ordering.OrderingLT
		}
		if left1 == right1 {
			return ordering.OrderingEQ
		}
		return ordering.OrderingGT
	}
	return ordering.OrderingGT
}

//line tree.go:215
func Eq(ra PathItem, rb PathItem) bool {
	var tail bool
	if ra.Name.Tag == maybe.TagNothing && rb.Name.Tag == maybe.TagNothing {
		tail = true
	} else if ra.Name.Tag == maybe.TagJust && rb.Name.Tag == maybe.TagJust {
		tail = ra.Name.Just == rb.Name.Just
	} else {
		tail = false
	}
	return ra.Index == rb.Index && tail
}

//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:157:1
func IsAllParallelizable(value *TreeTupleArrayPathItemString) bool {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:157:1
	switch value.Tag {
	case TreeTupleArrayPathItemStringNode:
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:158:16
		every := true
		for _, x := range value.Node1 {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:158:16
			if !IsAllParallelizable(x) {
				every = false
				break
			}
		}
		return every
	default: // TreeTupleArrayPathItemStringLeaf
		x := value.Leaf1
		var result bool
		switch x.Tag {
		case maybe.TagNothing:
			result = true
		default:
			result = false
		}
		var result2 maybe.Boolean
		switch x.Tag {
		case maybe.TagJust:
			result2 = x.Just.IsParallelizable
		default: // maybe.TagNothing
			result2 = maybe.NoBoolean
		}
		var result3 bool
		if result2 != maybe.NoBoolean {
			result3 = result2 == maybe.True
		} else {
			result3 = false
		}
		return result || result3
	}
}

//line tree.go:269
var AnnotatedWithPaths = InitializeAnnotatedWithPaths()

func InitializeAnnotatedWithPaths() func([]*TreeStringFnUnitAffUnit) []*TreeTupleArrayPathItemString {
	f := func(index int, value *TreeStringFnUnitAffUnit) *TreeTupleArrayPathItemString {
		return AnnotatedWithPathsGo([]PathItem{}, index, value)
	}
	return func(array []*TreeStringFnUnitAffUnit) []*TreeTupleArrayPathItemString {
//line ../../../../../../p/foldable-traversable-6.0.0/src/Data/FunctorWithIndex.purs:41:3
		return direct.MapWithIndex(f, array)
	}
}

//line tree.go:282
func FilterTreeF(value *TreeStringFnUnitAffUnit) *TreeStringFnUnitAffUnit {
	if value.Tag == TreeStringFnUnitAffUnitNode {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:140:5
		bound := make([]*TreeStringFnUnitAffUnit, 0, len(value.Node1))
		for _, x := range value.Node1 {
//line ../../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
			scrutinee := FilterTreeF(x)
			if scrutinee != nil {
//line ../../../../../../p/maybe-6.0.0/src/Data/Maybe.purs:239:22
				bound = append(bound, scrutinee)
			}
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:140:5
		if len(bound) == 0 {
			return nil
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:140:5
		return NodeTreeStringFnUnitAffUnit(value.Node0, bound)
	}
	a := value.Leaf1
	var result bool
	switch a.Tag {
	case maybe.TagJust:
		result = a.Just.IsFocused
	default: // maybe.TagNothing
		result = false
	}
	if result {
		return LeafTreeStringFnUnitAffUnit(value.Leaf0, a)
	}
	return nil
}

//line tree.go:316
func AnnotatedWithPathsGo(path []PathItem, index int, value *TreeStringFnUnitAffUnit) *TreeTupleArrayPathItemString {
	switch value.Tag {
	case TreeStringFnUnitAffUnitNode:
		c := value.Node0
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:121:20
		var name maybe.String
		switch c.Tag {
		case either.EitherFnUnitAffUnitLeft:
//line ../../../../../../p/either-6.1.0/src/Data/Either.purs:209:23
			name = maybe.JustString(c.Left)
		default: // either.EitherFnUnitAffUnitRight
			name = maybe.NoString
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:124:11
		var result Either
		switch c.Tag {
		case either.EitherFnUnitAffUnitLeft:
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:124:11
			result = Left(TupleArrayPathItem{First: path, Second: c.Left})
		default: // either.EitherFnUnitAffUnitRight
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:124:11
			result = Right(c.Right)
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:124:11
		array2 := array.Append(path, []PathItem{{Index: index, Name: name}})
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:124:11
		mapped := make([]*TreeTupleArrayPathItemString, len(value.Node1))
		for slot, element := range value.Node1 {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:124:11
			mapped[slot] = AnnotatedWithPathsGo(array2, slot, element)
		}
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Tree.purs:124:11
		return NodeTreeTupleArrayPathItemString(result, mapped)
	default: // TreeStringFnUnitAffUnitLeaf
		return LeafTreeTupleArrayPathItemString(TupleArrayPathItem{First: path, Second: value.Leaf0}, value.Leaf1)
	}
}
