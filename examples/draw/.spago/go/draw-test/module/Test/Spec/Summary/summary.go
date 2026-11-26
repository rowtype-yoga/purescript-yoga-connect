package summary

import (
	"iris.local/draw-test/Data/Maybe"
	"iris.local/draw-test/Test/Spec/Result"
	"iris.local/draw-test/shared/Test.Spec.Tree+Test.Spec.Result"
)

type Summary struct {
	Failed  int
	Passed  int
	Pending int
}

//line ../../../../../../p/foldable-traversable-6.0.0/src/Data/Foldable.purs:131:3
func Summarize(array []*tree_result.Tree) Summary {
//line ../../../../../../p/foldable-traversable-6.0.0/src/Data/Foldable.purs:131:3
	accumulator := Zero
//line ../../../../../../p/foldable-traversable-6.0.0/src/Data/Foldable.purs:131:3
	for index := len(array) - 1; index >= 0; index-- {
//line ../../../../../../p/foldable-traversable-6.0.0/src/Data/Foldable.purs:131:3
		x := array[index]
//line ../../../../../../p/foldable-traversable-6.0.0/src/Data/Foldable.purs:115:38
		var summary Summary
		if x.Tag == tree_result.TreeLeaf && x.Leaf1.Tag == maybe.TagJust && x.Leaf1.Just.Tag == result.ResultSuccess {
			summary = Summary{Passed: 1, Failed: 0, Pending: 0}
		} else if x.Tag == tree_result.TreeLeaf && x.Leaf1.Tag == maybe.TagJust && x.Leaf1.Just.Tag == result.ResultFailure {
			summary = Summary{Passed: 0, Failed: 1, Pending: 0}
		} else if x.Tag == tree_result.TreeLeaf && x.Leaf1.Tag == maybe.TagNothing {
			summary = Summary{Passed: 0, Failed: 0, Pending: 1}
		} else {
//line ../../../../../../p/spec-8.1.2/src/Test/Spec/Summary.purs:29:19
			summary = Summarize(x.Node1)
		}
		accumulator = Summary{Pending: summary.Pending + accumulator.Pending, Passed: summary.Passed + accumulator.Passed, Failed: summary.Failed + accumulator.Failed}
	}
	return accumulator
}

//line summary.go:41
var Zero Summary = Summary{Pending: 0, Passed: 0, Failed: 0}
