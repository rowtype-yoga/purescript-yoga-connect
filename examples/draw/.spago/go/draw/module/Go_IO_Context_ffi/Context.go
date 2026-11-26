package Go_IO_Context_ffi

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

func WithDeadlineImpl[R any](made func(context.Context) func(context.CancelFunc) R, millisecondsSinceEpoch float64, parent context.Context) R {
	deadline := time.UnixMicro(int64(millisecondsSinceEpoch * 1000))
	child, cancel := context.WithDeadline(parent, deadline)
	return made(child)(cancel)
}

func CancelImpl(cancel context.CancelFunc) { cancel() }

func DoneImpl(ctx context.Context) <-chan struct{} { return ctx.Done() }

func ErrImpl[R any](running R, cancelled R, deadlineExceeded R, ctx context.Context) R {
	switch err := ctx.Err(); {
	case err == nil:
		return running
	case errors.Is(err, context.DeadlineExceeded):
		return deadlineExceeded
	default:
		return cancelled
	}
}

func IsCancelledImpl(ctx context.Context) bool { return ctx.Err() != nil }

func duration(milliseconds float64) time.Duration {
	return time.Duration(milliseconds * float64(time.Millisecond))
}
