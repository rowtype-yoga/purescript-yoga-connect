package Go_IO_Internal_ffi

import (
	"strconv"

	"github.com/i-am-the-slime/iris-go/eventloop"
	"github.com/i-am-the-slime/iris-go/exception"
)

// LiftEffectImpl waits for the event loop, runs the effect there and returns its result; an
// exception it throws is rethrown on this goroutine.
func LiftEffectImpl[A any](effect func() A) A {
	var value A
	eventloop.Turn(func() { value = effect() })
	return value
}

// CatchImpl runs the action and hands whatever it panics with to the handler as an Error.
func CatchImpl[A any](handler func(exception.Error) A, action func() A) A {
	result, failure, failed := attempt(action)
	if failed {
		return handler(failure)
	}
	return result
}

// CatchResultsImpl is CatchImpl for an action that returns its result as a value and a failure,
// as Iris passes an Either. Iris calls it in place of CatchImpl at such actions.
func CatchResultsImpl[A, F any](handler func(exception.Error) (A, F), action func() (A, F)) (A, F) {
	value, err, failure, failed := attemptResults(action)
	if failed {
		return handler(failure)
	}
	return value, err
}

// BothImpl runs both effects on goroutines of their own and combines their results. The first
// exception either throws is rethrown here, without waiting for the other.
func BothImpl[A, B, C any](combine func(A) func(B) C, left func() A, right func() B) C {
	failures := make(chan *exception.Error, 2)
	var leftValue A
	var rightValue B
	go func() {
		leftResult, failure, failed := attempt(left)
		leftValue = leftResult
		failures <- failedWith(failure, failed)
	}()
	go func() {
		rightResult, failure, failed := attempt(right)
		rightValue = rightResult
		failures <- failedWith(failure, failed)
	}()
	for range 2 {
		rethrowIfFailed(<-failures)
	}
	return combine(leftValue)(rightValue)
}

// BothUnitImpl runs both actions on goroutines of their own and waits for both. The first
// exception either throws is rethrown here, without waiting for the other.
func BothUnitImpl(left func(), right func()) {
	failures := make(chan *exception.Error, 2)
	for _, action := range []func(){left, right} {
		go func() {
			_, failure, failed := attempt(func() struct{} {
				action()
				return struct{}{}
			})
			failures <- failedWith(failure, failed)
		}()
	}
	for range 2 {
		if failure := <-failures; failure != nil {
			panic(*failure)
		}
	}
}

// MapArrayConcurrentlyImpl applies f to every element on a goroutine of its own and returns the
// results in order. The first exception any throws is rethrown here, without waiting for the
// others.
func MapArrayConcurrentlyImpl[A, B any](f func(A) B, elements []A) []B {
	results := make([]B, len(elements))
	failures := make(chan *exception.Error, len(elements))
	for index, element := range elements {
		go func() {
			result, failure, failed := attempt(func() B { return f(element) })
			results[index] = result
			failures <- failedWith(failure, failed)
		}()
	}
	awaitAll(failures, len(elements))
	return results
}

// ForArrayConcurrentlyImpl runs f on every element on a goroutine of its own and waits for all
// of them. The first exception any throws is rethrown here, without waiting for the others.
func ForArrayConcurrentlyImpl[A any](elements []A, f func(A)) {
	failures := make(chan *exception.Error, len(elements))
	for _, element := range elements {
		go func() {
			_, failure, failed := attempt(func() struct{} {
				f(element)
				return struct{}{}
			})
			failures <- failedWith(failure, failed)
		}()
	}
	awaitAll(failures, len(elements))
}

// awaitAll waits for count goroutines to report and panics with the first failure.
func awaitAll(failures chan *exception.Error, count int) {
	for range count {
		if failure := <-failures; failure != nil {
			panic(*failure)
		}
	}
}

// FirstImpl runs both effects on goroutines of their own and returns the outcome of the first to
// finish.
func FirstImpl[A any](left func() A, right func() A) A {
	first := make(chan outcome[A], 2)
	for _, effect := range []func() A{left, right} {
		go func() {
			value, failure, failed := attempt(effect)
			first <- outcome[A]{value: value, failure: failedWith(failure, failed)}
		}()
	}
	winner := <-first
	rethrowIfFailed(winner.failure)
	return winner.value
}

func rethrowIfFailed(failure *exception.Error) {
	if failure != nil {
		panic(*failure)
	}
}

// outcome is how one of FirstImpl's effects ended.
type outcome[A any] struct {
	value   A
	failure *exception.Error
}

// NeverImpl waits forever.
func NeverImpl[A any]() A {
	select {}
}

func failedWith(failure exception.Error, failed bool) *exception.Error {
	if !failed {
		return nil
	}
	return &failure
}

func attempt[A any](action func() A) (result A, failure exception.Error, failed bool) {
	defer recoverInto(&failure, &failed)
	return action(), exception.Error{}, false
}

func attemptResults[A, F any](action func() (A, F)) (value A, err F, failure exception.Error, failed bool) {
	defer recoverInto(&failure, &failed)
	value, err = action()
	return value, err, exception.Error{}, false
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
