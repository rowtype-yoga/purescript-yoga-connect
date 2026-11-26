package codes

import (
	"github.com/i-am-the-slime/iris-go/alloc"
	"github.com/i-am-the-slime/iris-go/prim"
	"iris.local/draw-test/Data/List/Types"
	"strconv"
)

type RenderingMode uint8

const (
	RenderingModeBold RenderingMode = iota
	RenderingModeDim
	RenderingModeItalic
	RenderingModeUnderline
	RenderingModeInverse
	RenderingModeStrikethrough
)

type Color uint8

const (
	ColorBlack Color = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	ColorBrightBlack
	ColorBrightRed
	ColorBrightGreen
	ColorBrightYellow
	ColorBrightBlue
	ColorBrightMagenta
	ColorBrightCyan
	ColorBrightWhite
)

type GraphicsParam struct {
	Tag    uint8
	PMode  RenderingMode
	Field0 Color
}

const (
	GraphicsParamReset = iota
	GraphicsParamPMode
	GraphicsParamPForeground
	GraphicsParamPBackground
)

type ListGraphicsParam struct {
	Cons1 *ListGraphicsParam
	Cons0 GraphicsParam
}

var SharedListGraphicsParamNil = &ListGraphicsParam{}

type NonEmpty struct {
	Second *ListGraphicsParam
	First  GraphicsParam
}

type EraseParam uint8

const (
	EraseParamToEnd EraseParam = iota
	EraseParamFromBeginning
	EraseParamEntire
)

type EscapeCode struct {
	Graphics NonEmpty
	FieldA0  int
	FieldA1  int
	Tag      uint8
	FieldB0  EraseParam
}

const (
	EscapeCodeUp = iota
	EscapeCodeDown
	EscapeCodeForward
	EscapeCodeBack
	EscapeCodeNextLine
	EscapeCodePreviousLine
	EscapeCodeHorizontalAbsolute
	EscapeCodePosition
	EscapeCodeEraseData
	EscapeCodeEraseLine
	EscapeCodeScrollUp
	EscapeCodeScrollDown
	EscapeCodeGraphics
	EscapeCodeSavePosition
	EscapeCodeRestorePosition
	EscapeCodeQueryPosition
	EscapeCodeHideCursor
	EscapeCodeShowCursor
)

type AccInit struct {
	Acc  string
	Init bool
}

type ListListGraphicsParam struct {
	Cons0 *ListGraphicsParam
	Cons1 *ListListGraphicsParam
}

var SharedListListGraphicsParamNil = &ListListGraphicsParam{}

var SharedListStringNil = &types.ListString{}

func Graphics(graphics NonEmpty) EscapeCode {
	return EscapeCode{Tag: EscapeCodeGraphics, Graphics: graphics}
}

func PMode(pMode RenderingMode) GraphicsParam {
	return GraphicsParam{Tag: GraphicsParamPMode, PMode: pMode}
}

func PForeground(field0 Color) GraphicsParam {
	return GraphicsParam{Tag: GraphicsParamPForeground, Field0: field0}
}

func EscapeCodeToStringGo(accInit AccInit, x string) AccInit {
	if accInit.Init {
		return AccInit{Init: false, Acc: x}
	}
	return AccInit{Init: false, Acc: accInit.Acc + ";" + x}
}

//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:84:3
func EscapeCodeToStringChunkedRevMap(chunksAcc *ListListGraphicsParam, chunk *ListGraphicsParam) *types.ListString {
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:84:3
	var listListGraphicsParamCells alloc.Cells[ListListGraphicsParam]
	for {
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:84:3
		if chunk != SharedListGraphicsParamNil && chunk.Cons1 != SharedListGraphicsParamNil && chunk.Cons1.Cons1 != SharedListGraphicsParamNil {
			xs := chunk.Cons1.Cons1.Cons1
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:84:3
			chunksAcc, chunk = listListGraphicsParamCells.New(ListListGraphicsParam{Cons0: chunk, Cons1: chunksAcc}), xs
			continue
		}
		var result *types.ListString
		if chunk != SharedListGraphicsParamNil && chunk.Cons1 != SharedListGraphicsParamNil && chunk.Cons1.Cons1 == SharedListGraphicsParamNil {
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:90:35
			var result2 string
			switch chunk.Cons0.Tag {
			case GraphicsParamReset:
				result2 = "0"
			case GraphicsParamPMode:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				var result3 int
				switch chunk.Cons0.PMode {
				case RenderingModeBold:
					result3 = 1
				case RenderingModeDim:
					result3 = 2
				case RenderingModeItalic:
					result3 = 3
				case RenderingModeUnderline:
					result3 = 4
				case RenderingModeInverse:
					result3 = 7
				default: // RenderingModeStrikethrough
					result3 = 9
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				result2 = strconv.Itoa(result3)
			case GraphicsParamPForeground:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				var result3 int
				switch chunk.Cons0.Field0 {
				case ColorBlack:
					result3 = 30
				case ColorRed:
					result3 = 31
				case ColorGreen:
					result3 = 32
				case ColorYellow:
					result3 = 33
				case ColorBlue:
					result3 = 34
				case ColorMagenta:
					result3 = 35
				case ColorCyan:
					result3 = 36
				case ColorWhite:
					result3 = 37
				case ColorBrightBlack:
					result3 = 90
				case ColorBrightRed:
					result3 = 91
				case ColorBrightGreen:
					result3 = 92
				case ColorBrightYellow:
					result3 = 93
				case ColorBrightBlue:
					result3 = 94
				case ColorBrightMagenta:
					result3 = 95
				case ColorBrightCyan:
					result3 = 96
				default: // ColorBrightWhite
					result3 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				result2 = strconv.Itoa(result3)
			default: // GraphicsParamPBackground
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				var result3 int
				switch chunk.Cons0.Field0 {
				case ColorBlack:
					result3 = 30
				case ColorRed:
					result3 = 31
				case ColorGreen:
					result3 = 32
				case ColorYellow:
					result3 = 33
				case ColorBlue:
					result3 = 34
				case ColorMagenta:
					result3 = 35
				case ColorCyan:
					result3 = 36
				case ColorWhite:
					result3 = 37
				case ColorBrightBlack:
					result3 = 90
				case ColorBrightRed:
					result3 = 91
				case ColorBrightGreen:
					result3 = 92
				case ColorBrightYellow:
					result3 = 93
				case ColorBrightBlue:
					result3 = 94
				case ColorBrightMagenta:
					result3 = 95
				case ColorBrightCyan:
					result3 = 96
				default: // ColorBrightWhite
					result3 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				result2 = strconv.Itoa(result3 + 10)
			}
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:90:42
			var result4 string
			switch chunk.Cons1.Cons0.Tag {
			case GraphicsParamReset:
				result4 = "0"
			case GraphicsParamPMode:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				var result3 int
				switch chunk.Cons1.Cons0.PMode {
				case RenderingModeBold:
					result3 = 1
				case RenderingModeDim:
					result3 = 2
				case RenderingModeItalic:
					result3 = 3
				case RenderingModeUnderline:
					result3 = 4
				case RenderingModeInverse:
					result3 = 7
				default: // RenderingModeStrikethrough
					result3 = 9
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				result4 = strconv.Itoa(result3)
			case GraphicsParamPForeground:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				var result3 int
				switch chunk.Cons1.Cons0.Field0 {
				case ColorBlack:
					result3 = 30
				case ColorRed:
					result3 = 31
				case ColorGreen:
					result3 = 32
				case ColorYellow:
					result3 = 33
				case ColorBlue:
					result3 = 34
				case ColorMagenta:
					result3 = 35
				case ColorCyan:
					result3 = 36
				case ColorWhite:
					result3 = 37
				case ColorBrightBlack:
					result3 = 90
				case ColorBrightRed:
					result3 = 91
				case ColorBrightGreen:
					result3 = 92
				case ColorBrightYellow:
					result3 = 93
				case ColorBrightBlue:
					result3 = 94
				case ColorBrightMagenta:
					result3 = 95
				case ColorBrightCyan:
					result3 = 96
				default: // ColorBrightWhite
					result3 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				result4 = strconv.Itoa(result3)
			default: // GraphicsParamPBackground
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				var result3 int
				switch chunk.Cons1.Cons0.Field0 {
				case ColorBlack:
					result3 = 30
				case ColorRed:
					result3 = 31
				case ColorGreen:
					result3 = 32
				case ColorYellow:
					result3 = 33
				case ColorBlue:
					result3 = 34
				case ColorMagenta:
					result3 = 35
				case ColorCyan:
					result3 = 36
				case ColorWhite:
					result3 = 37
				case ColorBrightBlack:
					result3 = 90
				case ColorBrightRed:
					result3 = 91
				case ColorBrightGreen:
					result3 = 92
				case ColorBrightYellow:
					result3 = 93
				case ColorBrightBlue:
					result3 = 94
				case ColorBrightMagenta:
					result3 = 95
				case ColorBrightCyan:
					result3 = 96
				default: // ColorBrightWhite
					result3 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				result4 = strconv.Itoa(result3 + 10)
			}
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:90:42
			result = types.ConsListString(result2, types.ConsListString(result4, SharedListStringNil))
		} else if chunk != SharedListGraphicsParamNil && chunk.Cons1 == SharedListGraphicsParamNil {
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:91:30
			var result2 string
			switch chunk.Cons0.Tag {
			case GraphicsParamReset:
				result2 = "0"
			case GraphicsParamPMode:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				var result3 int
				switch chunk.Cons0.PMode {
				case RenderingModeBold:
					result3 = 1
				case RenderingModeDim:
					result3 = 2
				case RenderingModeItalic:
					result3 = 3
				case RenderingModeUnderline:
					result3 = 4
				case RenderingModeInverse:
					result3 = 7
				default: // RenderingModeStrikethrough
					result3 = 9
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				result2 = strconv.Itoa(result3)
			case GraphicsParamPForeground:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				var result3 int
				switch chunk.Cons0.Field0 {
				case ColorBlack:
					result3 = 30
				case ColorRed:
					result3 = 31
				case ColorGreen:
					result3 = 32
				case ColorYellow:
					result3 = 33
				case ColorBlue:
					result3 = 34
				case ColorMagenta:
					result3 = 35
				case ColorCyan:
					result3 = 36
				case ColorWhite:
					result3 = 37
				case ColorBrightBlack:
					result3 = 90
				case ColorBrightRed:
					result3 = 91
				case ColorBrightGreen:
					result3 = 92
				case ColorBrightYellow:
					result3 = 93
				case ColorBrightBlue:
					result3 = 94
				case ColorBrightMagenta:
					result3 = 95
				case ColorBrightCyan:
					result3 = 96
				default: // ColorBrightWhite
					result3 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				result2 = strconv.Itoa(result3)
			default: // GraphicsParamPBackground
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				var result3 int
				switch chunk.Cons0.Field0 {
				case ColorBlack:
					result3 = 30
				case ColorRed:
					result3 = 31
				case ColorGreen:
					result3 = 32
				case ColorYellow:
					result3 = 33
				case ColorBlue:
					result3 = 34
				case ColorMagenta:
					result3 = 35
				case ColorCyan:
					result3 = 36
				case ColorWhite:
					result3 = 37
				case ColorBrightBlack:
					result3 = 90
				case ColorBrightRed:
					result3 = 91
				case ColorBrightGreen:
					result3 = 92
				case ColorBrightYellow:
					result3 = 93
				case ColorBrightBlue:
					result3 = 94
				case ColorBrightMagenta:
					result3 = 95
				case ColorBrightCyan:
					result3 = 96
				default: // ColorBrightWhite
					result3 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				result2 = strconv.Itoa(result3 + 10)
			}
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:91:30
			result = types.ConsListString(result2, SharedListStringNil)
		} else {
			result = SharedListStringNil
		}
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:87:36
		return EscapeCodeToStringChunkedRevMapReverseUnrolledMap(chunksAcc, result)
	}
}

//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:95:5
func EscapeCodeToStringChunkedRevMapReverseUnrolledMap(list *ListListGraphicsParam, acc *types.ListString) *types.ListString {
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:95:5
	var listStringCells alloc.Cells[types.ListString]
	for {
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:95:5
		if list != SharedListListGraphicsParamNil && list.Cons0 != SharedListGraphicsParamNil && list.Cons0.Cons1 != SharedListGraphicsParamNil && list.Cons0.Cons1.Cons1 != SharedListGraphicsParamNil {
			x1 := list.Cons0.Cons0
			x2 := list.Cons0.Cons1.Cons0
			x3 := list.Cons0.Cons1.Cons1.Cons0
			cs := list.Cons1
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:96:30
			var result string
			switch x1.Tag {
			case GraphicsParamReset:
				result = "0"
			case GraphicsParamPMode:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				var result2 int
				switch x1.PMode {
				case RenderingModeBold:
					result2 = 1
				case RenderingModeDim:
					result2 = 2
				case RenderingModeItalic:
					result2 = 3
				case RenderingModeUnderline:
					result2 = 4
				case RenderingModeInverse:
					result2 = 7
				default: // RenderingModeStrikethrough
					result2 = 9
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				result = strconv.Itoa(result2)
			case GraphicsParamPForeground:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				var result2 int
				switch x1.Field0 {
				case ColorBlack:
					result2 = 30
				case ColorRed:
					result2 = 31
				case ColorGreen:
					result2 = 32
				case ColorYellow:
					result2 = 33
				case ColorBlue:
					result2 = 34
				case ColorMagenta:
					result2 = 35
				case ColorCyan:
					result2 = 36
				case ColorWhite:
					result2 = 37
				case ColorBrightBlack:
					result2 = 90
				case ColorBrightRed:
					result2 = 91
				case ColorBrightGreen:
					result2 = 92
				case ColorBrightYellow:
					result2 = 93
				case ColorBrightBlue:
					result2 = 94
				case ColorBrightMagenta:
					result2 = 95
				case ColorBrightCyan:
					result2 = 96
				default: // ColorBrightWhite
					result2 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				result = strconv.Itoa(result2)
			default: // GraphicsParamPBackground
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				var result2 int
				switch x1.Field0 {
				case ColorBlack:
					result2 = 30
				case ColorRed:
					result2 = 31
				case ColorGreen:
					result2 = 32
				case ColorYellow:
					result2 = 33
				case ColorBlue:
					result2 = 34
				case ColorMagenta:
					result2 = 35
				case ColorCyan:
					result2 = 36
				case ColorWhite:
					result2 = 37
				case ColorBrightBlack:
					result2 = 90
				case ColorBrightRed:
					result2 = 91
				case ColorBrightGreen:
					result2 = 92
				case ColorBrightYellow:
					result2 = 93
				case ColorBrightBlue:
					result2 = 94
				case ColorBrightMagenta:
					result2 = 95
				case ColorBrightCyan:
					result2 = 96
				default: // ColorBrightWhite
					result2 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				result = strconv.Itoa(result2 + 10)
			}
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:96:37
			var result3 string
			switch x2.Tag {
			case GraphicsParamReset:
				result3 = "0"
			case GraphicsParamPMode:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				var result2 int
				switch x2.PMode {
				case RenderingModeBold:
					result2 = 1
				case RenderingModeDim:
					result2 = 2
				case RenderingModeItalic:
					result2 = 3
				case RenderingModeUnderline:
					result2 = 4
				case RenderingModeInverse:
					result2 = 7
				default: // RenderingModeStrikethrough
					result2 = 9
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				result3 = strconv.Itoa(result2)
			case GraphicsParamPForeground:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				var result2 int
				switch x2.Field0 {
				case ColorBlack:
					result2 = 30
				case ColorRed:
					result2 = 31
				case ColorGreen:
					result2 = 32
				case ColorYellow:
					result2 = 33
				case ColorBlue:
					result2 = 34
				case ColorMagenta:
					result2 = 35
				case ColorCyan:
					result2 = 36
				case ColorWhite:
					result2 = 37
				case ColorBrightBlack:
					result2 = 90
				case ColorBrightRed:
					result2 = 91
				case ColorBrightGreen:
					result2 = 92
				case ColorBrightYellow:
					result2 = 93
				case ColorBrightBlue:
					result2 = 94
				case ColorBrightMagenta:
					result2 = 95
				case ColorBrightCyan:
					result2 = 96
				default: // ColorBrightWhite
					result2 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				result3 = strconv.Itoa(result2)
			default: // GraphicsParamPBackground
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				var result2 int
				switch x2.Field0 {
				case ColorBlack:
					result2 = 30
				case ColorRed:
					result2 = 31
				case ColorGreen:
					result2 = 32
				case ColorYellow:
					result2 = 33
				case ColorBlue:
					result2 = 34
				case ColorMagenta:
					result2 = 35
				case ColorCyan:
					result2 = 36
				case ColorWhite:
					result2 = 37
				case ColorBrightBlack:
					result2 = 90
				case ColorBrightRed:
					result2 = 91
				case ColorBrightGreen:
					result2 = 92
				case ColorBrightYellow:
					result2 = 93
				case ColorBrightBlue:
					result2 = 94
				case ColorBrightMagenta:
					result2 = 95
				case ColorBrightCyan:
					result2 = 96
				default: // ColorBrightWhite
					result2 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				result3 = strconv.Itoa(result2 + 10)
			}
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:96:44
			var result4 string
			switch x3.Tag {
			case GraphicsParamReset:
				result4 = "0"
			case GraphicsParamPMode:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				var result2 int
				switch x3.PMode {
				case RenderingModeBold:
					result2 = 1
				case RenderingModeDim:
					result2 = 2
				case RenderingModeItalic:
					result2 = 3
				case RenderingModeUnderline:
					result2 = 4
				case RenderingModeInverse:
					result2 = 7
				default: // RenderingModeStrikethrough
					result2 = 9
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
				result4 = strconv.Itoa(result2)
			case GraphicsParamPForeground:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				var result2 int
				switch x3.Field0 {
				case ColorBlack:
					result2 = 30
				case ColorRed:
					result2 = 31
				case ColorGreen:
					result2 = 32
				case ColorYellow:
					result2 = 33
				case ColorBlue:
					result2 = 34
				case ColorMagenta:
					result2 = 35
				case ColorCyan:
					result2 = 36
				case ColorWhite:
					result2 = 37
				case ColorBrightBlack:
					result2 = 90
				case ColorBrightRed:
					result2 = 91
				case ColorBrightGreen:
					result2 = 92
				case ColorBrightYellow:
					result2 = 93
				case ColorBrightBlue:
					result2 = 94
				case ColorBrightMagenta:
					result2 = 95
				case ColorBrightCyan:
					result2 = 96
				default: // ColorBrightWhite
					result2 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
				result4 = strconv.Itoa(result2)
			default: // GraphicsParamPBackground
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				var result2 int
				switch x3.Field0 {
				case ColorBlack:
					result2 = 30
				case ColorRed:
					result2 = 31
				case ColorGreen:
					result2 = 32
				case ColorYellow:
					result2 = 33
				case ColorBlue:
					result2 = 34
				case ColorMagenta:
					result2 = 35
				case ColorCyan:
					result2 = 36
				case ColorWhite:
					result2 = 37
				case ColorBrightBlack:
					result2 = 90
				case ColorBrightRed:
					result2 = 91
				case ColorBrightGreen:
					result2 = 92
				case ColorBrightYellow:
					result2 = 93
				case ColorBrightBlue:
					result2 = 94
				case ColorBrightMagenta:
					result2 = 95
				case ColorBrightCyan:
					result2 = 96
				default: // ColorBrightWhite
					result2 = 97
				}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
				result4 = strconv.Itoa(result2 + 10)
			}
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:96:44
			list, acc = cs, listStringCells.New(types.ListString{Cons0: result, Cons1: listStringCells.New(types.ListString{Cons0: result3, Cons1: listStringCells.New(types.ListString{Cons0: result4, Cons1: acc})})})
			continue
		}
		return acc
	}
}

//line codes.go:802
func EscapeCodeToStringGo2(b AccInit, value *types.ListString) AccInit {
	for {
		if value == SharedListStringNil {
			return b
		}
		a := value.Cons0
		as := value.Cons1
//line ../../../../../p/lists-7.0.0/src/Data/List/Types.purs:113:17
		b, value = EscapeCodeToStringGo(b, a), as
	}
}

//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
func AppendEscapeCodeToString(buffer []byte, x EscapeCode) []byte {
//line ../../../../../p/prelude-6.0.2/src/Control/Semigroupoid.purs:17:19
	buffer = append(buffer, "\u001b["...)
	switch x.Tag {
	case EscapeCodeUp:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:54:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:54:31
		buffer = append(buffer, 'A')
	case EscapeCodeDown:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:55:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:55:31
		buffer = append(buffer, 'B')
	case EscapeCodeForward:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:56:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:56:31
		buffer = append(buffer, 'C')
	case EscapeCodeBack:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:57:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:57:31
		buffer = append(buffer, 'D')
	case EscapeCodeNextLine:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:58:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:58:31
		buffer = append(buffer, 'E')
	case EscapeCodePreviousLine:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:59:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:59:31
		buffer = append(buffer, 'F')
	case EscapeCodeHorizontalAbsolute:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:60:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:60:31
		buffer = append(buffer, 'G')
	case EscapeCodePosition:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:61:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:61:31
		buffer = append(buffer, ';')
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:61:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA1), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:61:31
		buffer = append(buffer, 'H')
	case EscapeCodeEraseData:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:62:31
		switch x.FieldB0 {
		case EraseParamToEnd:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:96:22
			buffer = append(buffer, '0')
		case EraseParamFromBeginning:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:97:22
			buffer = append(buffer, '1')
		default: // EraseParamEntire
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:98:22
			buffer = append(buffer, '2')
		}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:62:31
		buffer = append(buffer, 'J')
	case EscapeCodeEraseLine:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:63:31
		switch x.FieldB0 {
		case EraseParamToEnd:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:96:22
			buffer = append(buffer, '0')
		case EraseParamFromBeginning:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:97:22
			buffer = append(buffer, '1')
		default: // EraseParamEntire
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:98:22
			buffer = append(buffer, '2')
		}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:63:31
		buffer = append(buffer, 'K')
	case EscapeCodeScrollUp:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:64:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:64:31
		buffer = append(buffer, 'S')
	case EscapeCodeScrollDown:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:65:31
		buffer = strconv.AppendInt(buffer, int64(x.FieldA0), 10)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:65:31
		buffer = append(buffer, 'T')
	case EscapeCodeGraphics:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:116:3
		var result string
		switch x.Graphics.First.Tag {
		case GraphicsParamReset:
			result = "0"
		case GraphicsParamPMode:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
			var result2 int
			switch x.Graphics.First.PMode {
			case RenderingModeBold:
				result2 = 1
			case RenderingModeDim:
				result2 = 2
			case RenderingModeItalic:
				result2 = 3
			case RenderingModeUnderline:
				result2 = 4
			case RenderingModeInverse:
				result2 = 7
			default: // RenderingModeStrikethrough
				result2 = 9
			}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:118:22
			result = strconv.Itoa(result2)
		case GraphicsParamPForeground:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
			var result2 int
			switch x.Graphics.First.Field0 {
			case ColorBlack:
				result2 = 30
			case ColorRed:
				result2 = 31
			case ColorGreen:
				result2 = 32
			case ColorYellow:
				result2 = 33
			case ColorBlue:
				result2 = 34
			case ColorMagenta:
				result2 = 35
			case ColorCyan:
				result2 = 36
			case ColorWhite:
				result2 = 37
			case ColorBrightBlack:
				result2 = 90
			case ColorBrightRed:
				result2 = 91
			case ColorBrightGreen:
				result2 = 92
			case ColorBrightYellow:
				result2 = 93
			case ColorBrightBlue:
				result2 = 94
			case ColorBrightMagenta:
				result2 = 95
			case ColorBrightCyan:
				result2 = 96
			default: // ColorBrightWhite
				result2 = 97
			}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:119:22
			result = strconv.Itoa(result2)
		default: // GraphicsParamPBackground
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
			var result2 int
			switch x.Graphics.First.Field0 {
			case ColorBlack:
				result2 = 30
			case ColorRed:
				result2 = 31
			case ColorGreen:
				result2 = 32
			case ColorYellow:
				result2 = 33
			case ColorBlue:
				result2 = 34
			case ColorMagenta:
				result2 = 35
			case ColorCyan:
				result2 = 36
			case ColorWhite:
				result2 = 37
			case ColorBrightBlack:
				result2 = 90
			case ColorBrightRed:
				result2 = 91
			case ColorBrightGreen:
				result2 = 92
			case ColorBrightYellow:
				result2 = 93
			case ColorBrightBlue:
				result2 = 94
			case ColorBrightMagenta:
				result2 = 95
			case ColorBrightCyan:
				result2 = 96
			default: // ColorBrightWhite
				result2 = 97
			}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:120:22
			result = strconv.Itoa(result2 + 10)
		}
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:116:3
		buffer = prim.AppendString(buffer, EscapeCodeToStringGo2(EscapeCodeToStringGo(AccInit{Init: true, Acc: ""}, result), EscapeCodeToStringChunkedRevMap(SharedListListGraphicsParamNil, x.Graphics.Second)).Acc)
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:116:3
		buffer = append(buffer, 'm')
	case EscapeCodeSavePosition:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:67:31
		buffer = append(buffer, 's')
	case EscapeCodeRestorePosition:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:68:31
		buffer = append(buffer, 'u')
	case EscapeCodeQueryPosition:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:69:31
		buffer = append(buffer, "6n"...)
	case EscapeCodeHideCursor:
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:70:31
		buffer = append(buffer, "?25l"...)
	default: // EscapeCodeShowCursor
//line ../../../../../p/ansi-7.0.0/src/Ansi/Codes.purs:71:31
		buffer = append(buffer, "?25h"...)
	}
	return buffer
}
