package Go_IO_IORef_ffi

import "sync/atomic"

// IORef is a cell whose value every goroutine sees whole: reads and writes swap a pointer.
type IORef[A any] = *ioref[A]

type ioref[A any] struct {
	value atomic.Pointer[A]
}

func (ref *ioref[A]) Load() A { return *ref.value.Load() }

func (ref *ioref[A]) Store(value A) { ref.value.Store(&value) }

// Modify applies f to the current value and stores its result, unless another goroutine stored
// one first, in which case it tries again with that one.
func (ref *ioref[A]) Modify(f func(A) A) {
	for {
		old := ref.value.Load()
		updated := f(*old)
		if ref.value.CompareAndSwap(old, &updated) {
			return
		}
	}
}

// New makes a cell holding value.
func New[A any](value A) *ioref[A] {
	ref := &ioref[A]{}
	ref.value.Store(&value)
	return ref
}

func NewIORefImpl[A any](value A) *ioref[A] { return New(value) }

func ReadIORefImpl[A any](ref *ioref[A]) A { return ref.Load() }

func WriteIORefImpl[A any](ref *ioref[A], value A) { ref.Store(value) }

// AtomicModifyIORefImpl applies f to the current value and stores the first of its results, unless
// another goroutine stored one first, in which case it tries again with that one.
func AtomicModifyIORefImpl[A, B, T any](first func(T) A, second func(T) B, f func(A) T, ref *ioref[A]) B {
	for {
		old := ref.value.Load()
		result := f(*old)
		updated := first(result)
		if ref.value.CompareAndSwap(old, &updated) {
			return second(result)
		}
	}
}
