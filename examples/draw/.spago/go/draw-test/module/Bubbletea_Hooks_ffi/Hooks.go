package Bubbletea_Hooks_ffi

import "sync"

// Outbox collects the commands that actions run until the program drains them, appending in place.
type Outbox[A any] = *outbox[A]

type outbox[A any] struct {
	mutex sync.Mutex
	items []A
}

func (box *outbox[A]) Push(item A) {
	box.mutex.Lock()
	box.items = append(box.items, item)
	box.mutex.Unlock()
}

func (box *outbox[A]) Drain() []A {
	box.mutex.Lock()
	items := box.items
	box.items = nil
	box.mutex.Unlock()
	if items == nil {
		return []A{}
	}
	return items
}

func NewOutbox[A any]() *outbox[A] { return &outbox[A]{} }

func NewOutboxImpl[A any]() *outbox[A] { return NewOutbox[A]() }

func PushImpl[A any](box *outbox[A], item A) { box.Push(item) }

func DrainImpl[A any](box *outbox[A]) []A { return box.Drain() }

// Registry holds what a scope's hooks registered, in order, appending in place. Only Bubble Tea's
// goroutine registers and reads, so it needs no lock.
type Registry[A any] = *registry[A]

type registry[A any] struct {
	items []A
}

func (registry *registry[A]) Register(item A) { registry.items = append(registry.items, item) }

// Registered is what the registry holds now; registering more later leaves it as it is.
func (registry *registry[A]) Registered() []A {
	return registry.items[:len(registry.items):len(registry.items)]
}

func NewRegistry[A any]() *registry[A] { return &registry[A]{} }

func NewRegistryImpl[A any]() *registry[A] { return NewRegistry[A]() }

func RegisterImpl[A any](registry *registry[A], item A) { registry.Register(item) }

func RegisteredImpl[A any](registry *registry[A]) []A { return registry.Registered() }
