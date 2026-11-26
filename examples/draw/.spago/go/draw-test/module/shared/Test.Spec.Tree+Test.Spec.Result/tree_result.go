package tree_result

import (
	"iris.local/draw-test/Data/Either"
	"iris.local/draw-test/Test/Spec/Result"
)

type Tree struct {
	Node0 *either.EitherVoid
	Node1 []*Tree
	Leaf0 string
	Leaf1 result.MaybeResult
	Tag   uint8
}

const (
	TreeNode = iota
	TreeLeaf
)
