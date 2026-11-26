package internal

type X any

type ProxyUnit struct {
	Pure   struct{}
	Field0 X
	Field1 func(struct{}) *ProxyUnit
	M      func() *ProxyUnit
	Tag    uint8
}

const (
	ProxyUnitRequest = iota
	ProxyUnitRespond
	ProxyUnitM
	ProxyUnitPure
)

type ProxyFnUnitUnit struct {
	Field0 X
	Field1 func(struct{}) *ProxyFnUnitUnit
	M      func() *ProxyFnUnitUnit
	Pure   func(struct{})
	Tag    uint8
}

const (
	ProxyFnUnitUnitRequest = iota
	ProxyFnUnitUnitRespond
	ProxyFnUnitUnitM
	ProxyFnUnitUnitPure
)

func PureProxyUnit(pure struct{}) *ProxyUnit {
	return &ProxyUnit{Tag: ProxyUnitPure, Pure: pure}
}

func RequestProxyFnUnitUnit(field0 X, field1 func(struct{}) *ProxyFnUnitUnit) *ProxyFnUnitUnit {
	return &ProxyFnUnitUnit{Tag: ProxyFnUnitUnitRequest, Field0: field0, Field1: field1}
}

func RespondProxyFnUnitUnit(field0 X, field1 func(struct{}) *ProxyFnUnitUnit) *ProxyFnUnitUnit {
	return &ProxyFnUnitUnit{Tag: ProxyFnUnitUnitRespond, Field0: field0, Field1: field1}
}

func MProxyFnUnitUnit(m func() *ProxyFnUnitUnit) *ProxyFnUnitUnit {
	return &ProxyFnUnitUnit{Tag: ProxyFnUnitUnitM, M: m}
}

func PureProxyFnUnitUnit(pure func(struct{})) *ProxyFnUnitUnit {
	return &ProxyFnUnitUnit{Tag: ProxyFnUnitUnitPure, Pure: pure}
}

func RequestProxyUnit(field0 X, field1 func(struct{}) *ProxyUnit) *ProxyUnit {
	return &ProxyUnit{Tag: ProxyUnitRequest, Field0: field0, Field1: field1}
}

func RespondProxyUnit(field0 X, field1 func(struct{}) *ProxyUnit) *ProxyUnit {
	return &ProxyUnit{Tag: ProxyUnitRespond, Field0: field0, Field1: field1}
}

func MProxyUnit(m func() *ProxyUnit) *ProxyUnit {
	return &ProxyUnit{Tag: ProxyUnitM, M: m}
}
