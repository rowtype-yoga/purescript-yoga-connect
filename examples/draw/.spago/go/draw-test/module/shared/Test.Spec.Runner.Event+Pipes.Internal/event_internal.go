package event_internal

import (
	"iris.local/draw-test/Test/Spec/Runner/Event"
	"iris.local/draw-test/shared/Test.Spec.Tree+Test.Spec.Result"
)

type Proxy struct {
	Request0 struct{}
	Request1 func(*event.Event) *Proxy
	Respond0 *event.Event
	Respond1 func(struct{}) *Proxy
	M        func() *Proxy
	Pure     []*tree_result.Tree
	Tag      uint8
}

const (
	ProxyRequest = iota
	ProxyRespond
	ProxyM
	ProxyPure
)
