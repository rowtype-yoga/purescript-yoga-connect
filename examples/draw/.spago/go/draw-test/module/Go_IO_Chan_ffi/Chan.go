package Go_IO_Chan_ffi

// Chan is a Go channel.
type Chan[A any] = chan A

// ReceiveChan is a Go channel that can only be received from, such as a context's Done.
type ReceiveChan[A any] = <-chan A

func NewChanImpl[A any]() chan A { return make(chan A) }

func NewBufferedChanImpl[A any](capacity int) chan A { return make(chan A, capacity) }

func ReceiveEndImpl[A any](channel chan A) <-chan A { return channel }

func SendImpl[A any](channel chan A, value A) { channel <- value }

func ReceiveImpl[A, R any](closed R, received func(A) R, channel <-chan A) R {
	value, open := <-channel
	if !open {
		return closed
	}
	return received(value)
}

func CloseImpl[A any](channel chan A) { close(channel) }

func TrySendImpl[A any](channel chan A, value A) bool {
	select {
	case channel <- value:
		return true
	default:
		return false
	}
}

func TryReceiveImpl[A, R any](empty R, closed R, received func(A) R, channel <-chan A) R {
	select {
	case value, open := <-channel:
		if !open {
			return closed
		}
		return received(value)
	default:
		return empty
	}
}

func CapacityImpl[A any](channel <-chan A) int { return cap(channel) }

func LengthImpl[A any](channel <-chan A) int { return len(channel) }
