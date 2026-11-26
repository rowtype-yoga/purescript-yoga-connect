package Go_IO_ffi

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/i-am-the-slime/iris-go/eventloop"
	"github.com/i-am-the-slime/iris-go/exception"
)

var threads atomic.Int32

// ForkIOImpl runs the action on a new goroutine. An exception it doesn't catch goes to standard
// error, as GHC reports one, and ends only that goroutine.
func ForkIOImpl(action func()) int {
	id := threads.Add(1)
	go func() {
		_, failure, failed := attempt(func() struct{} {
			action()
			return struct{}{}
		})
		if failed {
			os.Stderr.WriteString(filepath.Base(os.Args[0]) + ": thread " + strconv.Itoa(int(id)) + ": " + exception.ShowErrorImpl(failure) + "\n")
		}
	}()
	return int(id)
}

func ThreadDelayImpl(milliseconds float64) {
	time.Sleep(time.Duration(milliseconds * float64(time.Millisecond)))
}

// ThreadDelayWithImpl sleeps for the duration, or until the context is done, and says whether it
// slept all of it.
func ThreadDelayWithImpl(ctx context.Context, milliseconds float64) bool {
	if ctx.Err() != nil {
		return false
	}
	timer := time.NewTimer(time.Duration(milliseconds * float64(time.Millisecond)))
	defer timer.Stop()
	select {
	case <-timer.C:
		return true
	case <-ctx.Done():
		return false
	}
}

func YieldImpl() { runtime.Gosched() }

func NumCapabilitiesImpl() int { return runtime.GOMAXPROCS(0) }

// ToAffImpl is LaunchIOImpl, for toAff to resume a fiber with.
func ToAffImpl[A any](action func() A, succeeded func(A), failed func(exception.Error)) {
	LaunchIOImpl(action, succeeded, failed)
}

// LaunchIOImpl runs the action on a goroutine of its own and hands its outcome to succeeded or
// failed, IO actions on that goroutine. Until then it keeps the program running. An exception
// either throws ends the program as an uncaught Go panic: `panic: <message>`, Go's stack, exit
// status 2.
func LaunchIOImpl[A any](action func() A, succeeded func(A), failed func(exception.Error)) {
	retire := eventloop.Hold()
	go func() {
		value, failure, didFail := attempt(action)
		if didFail {
			failed(failure)
		} else {
			succeeded(value)
		}
		retire()
	}()
}

// LaunchImpl runs the action on a goroutine of its own and returns right away. Until it's done it
// keeps the program running. An exception it throws is rethrown on the event loop, where nothing
// catches it: it ends the program as an uncaught Go panic.
func LaunchImpl(action func()) {
	retire := eventloop.Hold()
	go func() {
		_, failure, failed := attempt(func() struct{} {
			action()
			return struct{}{}
		})
		if failed {
			eventloop.Turn(func() { panic(failure) })
		}
		retire()
	}()
}

func attempt[A any](action func() A) (result A, failure exception.Error, failed bool) {
	defer recoverInto(&failure, &failed)
	return action(), exception.Error{}, false
}

// recoverInto, deferred, turns a panic into the failure. It never sees the action's result, so
// attempt passes that on as it is, unboxed.
func recoverInto(failure *exception.Error, failed *bool) {
	if recovered := recover(); recovered != nil {
		*failure, *failed = asError(recovered), true
	}
}

func asError(recovered any) exception.Error {
	if failure, ok := exception.AsError(recovered); ok {
		return failure
	}
	return exception.NewError(panicMessage(recovered))
}

func panicMessage(recovered any) string {
	switch value := recovered.(type) {
	case error:
		return value.Error()
	case interface{ String() string }:
		return value.String()
	case string:
		return value
	case bool:
		return strconv.FormatBool(value)
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case uintptr:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'g', -1, 64)
	default:
		return "panic with a value that is not an error, a Stringer, a string or a number"
	}
}
