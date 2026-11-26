package ordering

type Ordering uint8

const (
	OrderingLT Ordering = iota
	OrderingGT
	OrderingEQ
)
