package key

import (
	"github.com/i-am-the-slime/iris-go/direct"
	"github.com/i-am-the-slime/iris-go/prim"
	key "iris.local/draw-test/Bubbletea_Key_ffi"
	datamaybe "iris.local/draw-test/Data/Maybe"
)

type KeyCode struct {
	Char   rune
	Field0 int
	Tag    uint8
}

const (
	KeyCodeChar = iota
	KeyCodeEnter
	KeyCodeTab
	KeyCodeEsc
	KeyCodeSpace
	KeyCodeBackspace
	KeyCodeDelete
	KeyCodeUp
	KeyCodeDown
	KeyCodeLeft
	KeyCodeRight
	KeyCodeHome
	KeyCodeEnd
	KeyCodePgUp
	KeyCodePgDown
	KeyCodeInsert
	KeyCodeF
	KeyCodeOther
)

type Key struct {
	Keystroke string
	Text      string
	Code      KeyCode
	Alt       bool
	Ctrl      bool
	Hyper     bool
	Meta      bool
	Repeat    bool
	Shift     bool
	Super     bool
}

type MaybeKeyCode struct {
	Just KeyCode
	Tag  uint8
}

var NoKeyCode = MaybeKeyCode{Tag: datamaybe.TagNothing}

type Tuple struct {
	First  KeyCode
	Second int
}

func JustKeyCode(just KeyCode) MaybeKeyCode {
	return MaybeKeyCode{Tag: datamaybe.TagJust, Just: just}
}

func Char(char rune) KeyCode {
	return KeyCode{Tag: KeyCodeChar, Char: char}
}

func Other(field0 int) KeyCode {
	return KeyCode{Tag: KeyCodeOther, Field0: field0}
}

func F(field0 int) KeyCode {
	return KeyCode{Tag: KeyCodeF, Field0: field0}
}

//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:146:1
func ToRune(value KeyCode) int {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:146:1
	switch value.Tag {
	case KeyCodeChar:
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:147:13
		return prim.ToCharCode(value.Char)
	case KeyCodeOther:
		return value.Field0
	default:
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:149:11
		findMap := direct.FindMap(datamaybe.NoInt, func(maybe datamaybe.Int) bool {
			return maybe.Tag == datamaybe.TagJust
		}, func(tuple Tuple) datamaybe.Int {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:149:53
			var result bool
			if tuple.First.Tag == KeyCodeChar && value.Tag == KeyCodeChar {
				result = tuple.First.Char == value.Char
			} else if tuple.First.Tag == KeyCodeEnter && value.Tag == KeyCodeEnter {
				result = true
			} else if tuple.First.Tag == KeyCodeTab && value.Tag == KeyCodeTab {
				result = true
			} else if tuple.First.Tag == KeyCodeEsc && value.Tag == KeyCodeEsc {
				result = true
			} else if tuple.First.Tag == KeyCodeSpace && value.Tag == KeyCodeSpace {
				result = true
			} else if tuple.First.Tag == KeyCodeBackspace && value.Tag == KeyCodeBackspace {
				result = true
			} else if tuple.First.Tag == KeyCodeDelete && value.Tag == KeyCodeDelete {
				result = true
			} else if tuple.First.Tag == KeyCodeUp && value.Tag == KeyCodeUp {
				result = true
			} else if tuple.First.Tag == KeyCodeDown && value.Tag == KeyCodeDown {
				result = true
			} else if tuple.First.Tag == KeyCodeLeft && value.Tag == KeyCodeLeft {
				result = true
			} else if tuple.First.Tag == KeyCodeRight && value.Tag == KeyCodeRight {
				result = true
			} else if tuple.First.Tag == KeyCodeHome && value.Tag == KeyCodeHome {
				result = true
			} else if tuple.First.Tag == KeyCodeEnd && value.Tag == KeyCodeEnd {
				result = true
			} else if tuple.First.Tag == KeyCodePgUp && value.Tag == KeyCodePgUp {
				result = true
			} else if tuple.First.Tag == KeyCodePgDown && value.Tag == KeyCodePgDown {
				result = true
			} else if tuple.First.Tag == KeyCodeInsert && value.Tag == KeyCodeInsert {
				result = true
			} else if tuple.First.Tag == KeyCodeF && value.Tag == KeyCodeF || tuple.First.Tag == KeyCodeOther && value.Tag == KeyCodeOther {
				result = tuple.First.Field0 == value.Field0
			} else {
				result = false
			}
			if result {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:149:75
				return datamaybe.JustInt(tuple.Second)
			}
			return datamaybe.NoInt
		}, SpecialKeys)
		switch findMap.Tag {
		case datamaybe.TagNothing:
			return 0
		default: // datamaybe.TagJust
			return findMap.Just
		}
	}
}

//line key.go:147
var SpecialKeys = InitializeSpecialKeys()

//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
func InitializeSpecialKeys() []Tuple {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built := make([]Tuple, len(array)+15)
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[0] = Tuple{First: KeyCode{Tag: KeyCodeEnter}, Second: key.KeyEnter}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[1] = Tuple{First: KeyCode{Tag: KeyCodeTab}, Second: key.KeyTab}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[2] = Tuple{First: KeyCode{Tag: KeyCodeEsc}, Second: key.KeyEsc}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[3] = Tuple{First: KeyCode{Tag: KeyCodeSpace}, Second: key.KeySpace}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[4] = Tuple{First: KeyCode{Tag: KeyCodeBackspace}, Second: key.KeyBackspace}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[5] = Tuple{First: KeyCode{Tag: KeyCodeDelete}, Second: key.KeyDelete}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[6] = Tuple{First: KeyCode{Tag: KeyCodeUp}, Second: key.KeyUp}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[7] = Tuple{First: KeyCode{Tag: KeyCodeDown}, Second: key.KeyDown}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[8] = Tuple{First: KeyCode{Tag: KeyCodeLeft}, Second: key.KeyLeft}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[9] = Tuple{First: KeyCode{Tag: KeyCodeRight}, Second: key.KeyRight}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[10] = Tuple{First: KeyCode{Tag: KeyCodeHome}, Second: key.KeyHome}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[11] = Tuple{First: KeyCode{Tag: KeyCodeEnd}, Second: key.KeyEnd}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[12] = Tuple{First: KeyCode{Tag: KeyCodePgUp}, Second: key.KeyPgUp}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[13] = Tuple{First: KeyCode{Tag: KeyCodePgDown}, Second: key.KeyPgDown}
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
	built[14] = Tuple{First: KeyCode{Tag: KeyCodeInsert}, Second: key.KeyInsert}
	for index, n := range array {
//line ../../../../../p/bubbletea/d52564695a184378df38f387d44a54d0439a43a9/src/Bubbletea/Key.purs:182:3
		built[index+15] = Tuple{First: F(n), Second: key.KeyF1 + n - 1}
	}
	return built
}
