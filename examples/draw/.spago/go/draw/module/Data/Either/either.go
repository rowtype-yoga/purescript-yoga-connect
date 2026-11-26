package either

import "github.com/i-am-the-slime/iris-go/exception"

type Either struct {
	Right struct{}
	Left  string
	Tag   uint8
}

const (
	EitherLeft = iota
	EitherRight
)

func Left(left string) Either {
	return Either{Tag: EitherLeft, Left: left}
}

type ErrorFailure struct {
	Failed bool
	Left   exception.Error
}

func Right(right struct{}) Either {
	return Either{Tag: EitherRight, Right: right}
}
