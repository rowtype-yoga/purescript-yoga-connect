package ffi

import "time"

var started = time.Now()

func NowInMillisecondsImpl() float64 {
	return float64(time.Since(started).Nanoseconds()) / 1e6
}
