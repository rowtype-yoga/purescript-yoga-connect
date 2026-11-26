package shared

import (
	lipgloss "iris.local/draw-test/Lipgloss_ffi"
	bytes "iris.local/draw-test/Yoga_Protobuf_Bytes_ffi"
)

type Seen struct {
	ByAda string
	ByBob string
}

type ListenOptions struct {
	Host string
	Port int
}

type Size struct {
	Height int
	Width  int
}

type Theme struct {
	Accent            lipgloss.ColorType
	BgBase            lipgloss.ColorType
	BgLeastVisible    lipgloss.ColorType
	BgLessVisible     lipgloss.ColorType
	BgMostVisible     lipgloss.ColorType
	Busy              lipgloss.ColorType
	Destructive       lipgloss.ColorType
	Error             lipgloss.ColorType
	FgBase            lipgloss.ColorType
	FgMoreSubtle      lipgloss.ColorType
	FgMostSubtle      lipgloss.ColorType
	FgSubtle          lipgloss.ColorType
	Info              lipgloss.ColorType
	InfoMoreSubtle    lipgloss.ColorType
	InfoMostSubtle    lipgloss.ColorType
	Keyword           lipgloss.ColorType
	OnPrimary         lipgloss.ColorType
	Primary           lipgloss.ColorType
	Secondary         lipgloss.ColorType
	Separator         lipgloss.ColorType
	Success           lipgloss.ColorType
	SuccessMoreSubtle lipgloss.ColorType
	SuccessMostSubtle lipgloss.ColorType
	Warning           lipgloss.ColorType
	WarningSubtle     lipgloss.ColorType
}

type XY struct {
	X int
	Y int
}

type ErrorDetail struct {
	Bytes    bytes.Bytes
	TypeName string
}
