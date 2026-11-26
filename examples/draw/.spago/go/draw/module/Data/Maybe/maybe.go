package maybe

const TagNothing = 0

const TagJust = 1

type Int struct {
	Just int
	Tag  uint8
}

func JustInt(just int) Int {
	return Int{Tag: TagJust, Just: just}
}

var NoInt = Int{Tag: TagNothing}

type Maybe struct {
	Just any
	Tag  uint8
}

type String struct {
	Just string
	Tag  uint8
}
