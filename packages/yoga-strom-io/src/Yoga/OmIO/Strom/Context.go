package ffi

import (
	"context"
	"errors"
	"time"
)

// Context is Go's context.Context, so that other foreign code takes it as it is.
type Context = context.Context

// CancelContext is Go's context.CancelFunc.
type CancelContext = context.CancelFunc

var Background context.Context = context.Background()

func WithCancelImpl[R any](made func(context.Context) func(context.CancelFunc) R, parent context.Context) R {
	child, cancel := context.WithCancel(parent)
	return made(child)(cancel)
}

func WithTimeoutImpl[R any](made func(context.Context) func(context.CancelFunc) R, milliseconds float64, parent context.Context) R {
	child, cancel := context.WithTimeout(parent, duration(milliseconds))
	return made(child)(cancel)
}

func CancelImpl(cancel context.CancelFunc) { cancel() }

func IsDoneImpl(ctx context.Context) bool { return ctx.Err() != nil }

func DeadlinePassedImpl(ctx context.Context) bool {
	return errors.Is(ctx.Err(), context.DeadlineExceeded)
}

func SleepUnlessDoneImpl(ctx context.Context, milliseconds float64) bool {
	if ctx.Err() != nil {
		return false
	}
	timer := time.NewTimer(duration(milliseconds))
	defer timer.Stop()
	select {
	case <-timer.C:
		return true
	case <-ctx.Done():
		return false
	}
}

func duration(milliseconds float64) time.Duration {
	return time.Duration(milliseconds * float64(time.Millisecond))
}
