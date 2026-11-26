package strom

import (
	"iris.local/draw/Data/Maybe"
	"iris.local/draw/shared"
)

type StepXY struct {
	Emit []shared.XY
	Tag  uint8
}

const (
	StepXYEmit = iota
	StepXYEnd
)

type StepBoolean struct {
	Emit []bool
	Tag  uint8
}

const (
	StepBooleanEmit = iota
	StepBooleanEnd
)

type Step struct {
	Emit []any
	Tag  uint8
}

const (
	StepEmit = iota
	StepEnd
)

func EmitStepXY(emit []shared.XY) StepXY {
	return StepXY{Tag: StepXYEmit, Emit: emit}
}

func EmitStepBoolean(emit []bool) StepBoolean {
	return StepBoolean{Tag: StepBooleanEmit, Emit: emit}
}

//line ../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1032:5
func OpenClosingAtEndStreamGo(chunk []bool, i int) maybe.Int {
	for {
//line ../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1032:5
		if i < len(chunk) {
//line ../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1036:17
			if chunk[i] {
				i += 1
				continue
			}
//line ../../../../../../p/arrays-7.3.0/src/Data/Array.purs:1036:45
			return maybe.JustInt(i)
		}
		return maybe.NoInt
	}
}
