package Go_IO_Select_ffi

// Case is one case of a select: a receive, or a send of value. Each select below picks the Go
// select statement with exactly its cases, so the runtime never walks a case that can't go.
type Case[A, R any] struct {
	sends     bool
	receiving <-chan A
	received  func(A, bool) R
	sending   chan<- A
	value     A
	sent      func() R
}

func OnReceiveImpl[A, R any](channel <-chan A, received func(A, bool) R) Case[A, R] {
	return Case[A, R]{receiving: channel, received: received}
}

func OnSendImpl[A, R any](channel chan A, value A, sent func() R) Case[A, R] {
	return Case[A, R]{sends: true, sending: channel, value: value, sent: sent}
}

func Select2Impl[A, B, R any](first Case[A, R], second Case[B, R]) R {
	switch {
	case !first.sends && !second.sends:
		select {
		case value, open := <-first.receiving:
			return first.received(value, open)
		case value, open := <-second.receiving:
			return second.received(value, open)
		}
	case first.sends && !second.sends:
		select {
		case first.sending <- first.value:
			return first.sent()
		case value, open := <-second.receiving:
			return second.received(value, open)
		}
	case !first.sends && second.sends:
		select {
		case value, open := <-first.receiving:
			return first.received(value, open)
		case second.sending <- second.value:
			return second.sent()
		}
	default:
		select {
		case first.sending <- first.value:
			return first.sent()
		case second.sending <- second.value:
			return second.sent()
		}
	}
}

func Select3Impl[A, B, C, R any](first Case[A, R], second Case[B, R], third Case[C, R]) R {
	switch {
	case !first.sends && !second.sends && !third.sends:
		select {
		case value, open := <-first.receiving:
			return first.received(value, open)
		case value, open := <-second.receiving:
			return second.received(value, open)
		case value, open := <-third.receiving:
			return third.received(value, open)
		}
	case first.sends && !second.sends && !third.sends:
		select {
		case first.sending <- first.value:
			return first.sent()
		case value, open := <-second.receiving:
			return second.received(value, open)
		case value, open := <-third.receiving:
			return third.received(value, open)
		}
	case !first.sends && second.sends && !third.sends:
		select {
		case value, open := <-first.receiving:
			return first.received(value, open)
		case second.sending <- second.value:
			return second.sent()
		case value, open := <-third.receiving:
			return third.received(value, open)
		}
	case !first.sends && !second.sends && third.sends:
		select {
		case value, open := <-first.receiving:
			return first.received(value, open)
		case value, open := <-second.receiving:
			return second.received(value, open)
		case third.sending <- third.value:
			return third.sent()
		}
	case first.sends && second.sends && !third.sends:
		select {
		case first.sending <- first.value:
			return first.sent()
		case second.sending <- second.value:
			return second.sent()
		case value, open := <-third.receiving:
			return third.received(value, open)
		}
	case first.sends && !second.sends && third.sends:
		select {
		case first.sending <- first.value:
			return first.sent()
		case value, open := <-second.receiving:
			return second.received(value, open)
		case third.sending <- third.value:
			return third.sent()
		}
	case !first.sends && second.sends && third.sends:
		select {
		case value, open := <-first.receiving:
			return first.received(value, open)
		case second.sending <- second.value:
			return second.sent()
		case third.sending <- third.value:
			return third.sent()
		}
	default:
		select {
		case first.sending <- first.value:
			return first.sent()
		case second.sending <- second.value:
			return second.sent()
		case third.sending <- third.value:
			return third.sent()
		}
	}
}

func TrySelect1Impl[A, R any](first Case[A, R], otherwise func() R) R {
	if first.sends {
		select {
		case first.sending <- first.value:
			return first.sent()
		default:
			return otherwise()
		}
	}
	select {
	case value, open := <-first.receiving:
		return first.received(value, open)
	default:
		return otherwise()
	}
}

func TrySelect2Impl[A, B, R any](first Case[A, R], second Case[B, R], otherwise func() R) R {
	switch {
	case !first.sends && !second.sends:
		select {
		case value, open := <-first.receiving:
			return first.received(value, open)
		case value, open := <-second.receiving:
			return second.received(value, open)
		default:
			return otherwise()
		}
	case first.sends && !second.sends:
		select {
		case first.sending <- first.value:
			return first.sent()
		case value, open := <-second.receiving:
			return second.received(value, open)
		default:
			return otherwise()
		}
	case !first.sends && second.sends:
		select {
		case value, open := <-first.receiving:
			return first.received(value, open)
		case second.sending <- second.value:
			return second.sent()
		default:
			return otherwise()
		}
	default:
		select {
		case first.sending <- first.value:
			return first.sent()
		case second.sending <- second.value:
			return second.sent()
		default:
			return otherwise()
		}
	}
}
