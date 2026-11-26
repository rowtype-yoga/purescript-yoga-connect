//go:build !(goexperiment.simd && (arm64 || wasm))

package main

func foldLanes(xs []int, initial int) int {
	accumulator := initial
	for _, x := range xs {
		accumulator = max(accumulator, x)
	}
	return accumulator
}
