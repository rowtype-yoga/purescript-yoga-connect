package core

import (
	"iris.local/draw-test/Control/Category"
	i "iris.local/draw-test/Pipes/Internal"
)

func ComposeResponseFbGoGo(p *i.ProxyUnit) *i.ProxyFnUnitUnit {
	switch p.Tag {
	case i.ProxyUnitRequest:
		return i.RequestProxyFnUnitUnit(p.Field0, func(x struct{}) *i.ProxyFnUnitUnit {
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:29:42
			return ComposeResponseFbGoGo(p.Field1(x))
		})
	case i.ProxyUnitRespond:
		return i.RespondProxyFnUnitUnit(p.Field0, func(x struct{}) *i.ProxyFnUnitUnit {
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:30:42
			return ComposeResponseFbGoGo(p.Field1(x))
		})
	case i.ProxyUnitM:
		return i.MProxyFnUnitUnit(func() *i.ProxyFnUnitUnit {
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:37
			proxy := p.M()
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:49
			return ComposeResponseFbGoGo(proxy)
		})
	default: // i.ProxyUnitPure
		return i.PureProxyFnUnitUnit(category.Identity)
	}
}

//line core.go:33
func ComposeResponseFbGoGoM(m func() *i.ProxyUnit) *i.ProxyFnUnitUnit {
	return i.MProxyFnUnitUnit(func() *i.ProxyFnUnitUnit {
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:37
		proxy := m()
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:31:49
		return ComposeResponseFbGoGo(proxy)
	})
}

//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
func ComposeResponseFbGoGo2(pf *i.ProxyFnUnitUnit) *i.ProxyUnit {
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
	switch pf.Tag {
	case i.ProxyFnUnitUnitRequest:
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return i.RequestProxyUnit(pf.Field0, func(x struct{}) *i.ProxyUnit {
			return ComposeResponseFbGoGo2(pf.Field1(x))
		})
	case i.ProxyFnUnitUnitRespond:
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return i.RespondProxyUnit(pf.Field0, func(x struct{}) *i.ProxyUnit {
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:38:42
			return ComposeResponseFbGoGo2(pf.Field1(x))
		})
	case i.ProxyFnUnitUnitM:
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:36:13
		return i.MProxyUnit(func() *i.ProxyUnit {
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:39:37
			proxy := pf.M()
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:39:49
			return ComposeResponseFbGoGo2(proxy)
		})
	default: // i.ProxyFnUnitUnitPure
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:40:25
		pf.Pure(struct{}{})
//line ../../../../../p/pipes-8.0.0/src/Pipes/Internal.purs:40:25
		return i.PureProxyUnit(struct{}{})
	}
}
