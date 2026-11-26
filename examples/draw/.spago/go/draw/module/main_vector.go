//go:build goexperiment.simd && wasm

package main

import (
	"simd"
	"unsafe"
)

func foldLanesLanes(xs []int64, initial int64) int64 {
	acc0 := simd.BroadcastInt64s(initial)
	acc1, acc2, acc3 := acc0, acc0, acc0
	step := acc0.Len()
	i := 0
	for ; i < len(xs)-4*step+1; i += 4 * step {
		x0 := simd.LoadInt64s(xs[i : i+step])
		acc0 = acc0.IfElse(acc0.Greater(x0), x0)
		x1 := simd.LoadInt64s(xs[i+step : i+2*step])
		acc1 = acc1.IfElse(acc1.Greater(x1), x1)
		x2 := simd.LoadInt64s(xs[i+2*step : i+3*step])
		acc2 = acc2.IfElse(acc2.Greater(x2), x2)
		x3 := simd.LoadInt64s(xs[i+3*step : i+4*step])
		acc3 = acc3.IfElse(acc3.Greater(x3), x3)
	}
	acc0 = acc3.IfElse(acc2.Less(acc3), acc2).IfElse(acc1.IfElse(acc0.Less(acc1), acc0).Less(acc3.IfElse(acc2.Less(acc3), acc2)), acc1.IfElse(acc0.Less(acc1), acc0))
	var lanes [8]int64
	acc0.Store(lanes[:])
	accumulator := initial
	for _, lane := range lanes[:step] {
		if accumulator < lane {
			accumulator = lane
		}
	}
	for ; i < len(xs); i++ {
		x := xs[i]
		accumulator = max(accumulator, x)
	}
	return accumulator
}
func foldLanes(xs []int, initial int) int {
	return int(foldLanesLanes(unsafe.Slice((*int64)(unsafe.Pointer(unsafe.SliceData(xs))), len(xs)), int64(initial)))
}
