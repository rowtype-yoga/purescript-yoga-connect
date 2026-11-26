package Go_IO_MVar_ffi

import (
	"context"
	"sync"
)

// MVar is a box that is empty or holds one value.
type MVar[A any] = *mvar[A]

type mvar[A any] struct {
	lock    sync.Mutex
	changed sync.Cond
	full    bool
	value   A
}

func NewEmptyMVarImpl[A any]() *mvar[A] {
	box := &mvar[A]{}
	box.changed.L = &box.lock
	return box
}

func TakeMVarImpl[A any](box *mvar[A]) A {
	box.lock.Lock()
	defer box.lock.Unlock()
	for !box.full {
		box.changed.Wait()
	}
	return take(box)
}

func PutMVarImpl[A any](box *mvar[A], value A) {
	box.lock.Lock()
	defer box.lock.Unlock()
	for box.full {
		box.changed.Wait()
	}
	put(box, value)
}

func ReadMVarImpl[A any](box *mvar[A]) A {
	box.lock.Lock()
	defer box.lock.Unlock()
	for !box.full {
		box.changed.Wait()
	}
	return box.value
}

func TakeMVarWithImpl[A, R any](interrupted R, taken func(A) R, ctx context.Context, box *mvar[A]) R {
	box.lock.Lock()
	defer box.lock.Unlock()
	if !awaitFullness(ctx, box, true) {
		return interrupted
	}
	return taken(take(box))
}

func PutMVarWithImpl[A any](ctx context.Context, box *mvar[A], value A) bool {
	box.lock.Lock()
	defer box.lock.Unlock()
	if !awaitFullness(ctx, box, false) {
		return false
	}
	put(box, value)
	return true
}

func ReadMVarWithImpl[A, R any](interrupted R, read func(A) R, ctx context.Context, box *mvar[A]) R {
	box.lock.Lock()
	defer box.lock.Unlock()
	if !awaitFullness(ctx, box, true) {
		return interrupted
	}
	return read(box.value)
}

// awaitFullness waits, holding the box's lock, until the box is full or empty as asked, and says
// whether it is. It isn't once the context is done: a waiter that is woken then sees it.
func awaitFullness[A any](ctx context.Context, box *mvar[A], full bool) bool {
	if box.full == full {
		return true
	}
	if ctx.Err() != nil {
		return false
	}
	stopWaking := context.AfterFunc(ctx, func() {
		box.lock.Lock()
		box.changed.Broadcast()
		box.lock.Unlock()
	})
	defer stopWaking()
	for box.full != full {
		if ctx.Err() != nil {
			return false
		}
		box.changed.Wait()
	}
	return true
}

func TryTakeMVarImpl[A, R any](empty R, found func(A) R, box *mvar[A]) R {
	box.lock.Lock()
	defer box.lock.Unlock()
	if !box.full {
		return empty
	}
	return found(take(box))
}

func TryPutMVarImpl[A any](box *mvar[A], value A) bool {
	box.lock.Lock()
	defer box.lock.Unlock()
	if box.full {
		return false
	}
	put(box, value)
	return true
}

func TryReadMVarImpl[A, R any](empty R, found func(A) R, box *mvar[A]) R {
	box.lock.Lock()
	defer box.lock.Unlock()
	if !box.full {
		return empty
	}
	return found(box.value)
}

func IsEmptyMVarImpl[A any](box *mvar[A]) bool {
	box.lock.Lock()
	defer box.lock.Unlock()
	return !box.full
}

// take empties the box, which holds its lock, and wakes its waiters.
func take[A any](box *mvar[A]) A {
	value := box.value
	var zero A
	box.value, box.full = zero, false
	box.changed.Broadcast()
	return value
}

// put fills the box, which holds its lock, and wakes its waiters.
func put[A any](box *mvar[A], value A) {
	box.value, box.full = value, true
	box.changed.Broadcast()
}
