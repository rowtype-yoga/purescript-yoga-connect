package ffi

import "context"

// Handoff is a channel between the goroutines of one stream.
type Handoff[A any] = chan A

func NewHandoffImpl[A any](capacity int32) chan A {
	return make(chan A, capacity)
}

func OfferImpl[A any](ctx context.Context, handoff chan A, value A) bool {
	if ctx.Err() != nil {
		return false
	}
	select {
	case handoff <- value:
		return true
	case <-ctx.Done():
		return false
	}
}

func ReceiveImpl[A, R any](interrupted R, received func(A) R, ctx context.Context, handoff chan A) R {
	select {
	case value := <-handoff:
		return received(value)
	case <-ctx.Done():
		return interrupted
	}
}

// Permits holds a token for each permit taken.
type Permits = chan struct{}

func NewPermitsImpl(count int32) chan struct{} {
	return make(chan struct{}, count)
}

func AcquireImpl(ctx context.Context, permits chan struct{}) bool {
	if ctx.Err() != nil {
		return false
	}
	select {
	case permits <- struct{}{}:
		return true
	case <-ctx.Done():
		return false
	}
}

func ReleaseImpl(permits chan struct{}) { <-permits }

func AcquireManyImpl(ctx context.Context, permits chan struct{}, count int32) bool {
	for range count {
		if !AcquireImpl(ctx, permits) {
			return false
		}
	}
	return true
}

func ReleaseManyImpl(permits chan struct{}, count int32) {
	for range count {
		<-permits
	}
}

// AwaitAllReleasedImpl takes every permit, then gives them back.
func AwaitAllReleasedImpl(permits chan struct{}) {
	for range cap(permits) {
		permits <- struct{}{}
	}
	for range cap(permits) {
		<-permits
	}
}
