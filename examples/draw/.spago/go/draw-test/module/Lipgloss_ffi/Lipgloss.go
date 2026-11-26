package Lipgloss_ffi

import (
	"image/color"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

type ColorType = color.Color

type Style = lipgloss.Style

type Position = lipgloss.Position

type BorderType = lipgloss.Border

func ColorImpl(text string) color.Color { return lipgloss.Color(text) }

var NoColor color.Color = lipgloss.NoColor{}

func LightDarkImpl(isDark bool, light color.Color, dark color.Color) color.Color {
	return lipgloss.LightDark(isDark)(light, dark)
}

var NewStyle = lipgloss.NewStyle()

func RenderImpl(style lipgloss.Style, text string) string { return style.Render(text) }

func StringWidthImpl(text string) int  { return lipgloss.Width(text) }
func StringHeightImpl(text string) int { return lipgloss.Height(text) }

func ForegroundImpl(c color.Color, s lipgloss.Style) lipgloss.Style { return s.Foreground(c) }
func BackgroundImpl(c color.Color, s lipgloss.Style) lipgloss.Style { return s.Background(c) }

func BoldImpl(on bool, s lipgloss.Style) lipgloss.Style          { return s.Bold(on) }
func ItalicImpl(on bool, s lipgloss.Style) lipgloss.Style        { return s.Italic(on) }
func UnderlineImpl(on bool, s lipgloss.Style) lipgloss.Style     { return s.Underline(on) }
func StrikethroughImpl(on bool, s lipgloss.Style) lipgloss.Style { return s.Strikethrough(on) }
func FaintImpl(on bool, s lipgloss.Style) lipgloss.Style         { return s.Faint(on) }
func ReverseImpl(on bool, s lipgloss.Style) lipgloss.Style       { return s.Reverse(on) }
func BlinkImpl(on bool, s lipgloss.Style) lipgloss.Style         { return s.Blink(on) }
func InlineImpl(on bool, s lipgloss.Style) lipgloss.Style        { return s.Inline(on) }

func PaddingImpl(top, right, bottom, left int, s lipgloss.Style) lipgloss.Style {
	return s.Padding(top, right, bottom, left)
}
func PaddingTopImpl(n int, s lipgloss.Style) lipgloss.Style    { return s.PaddingTop(n) }
func PaddingRightImpl(n int, s lipgloss.Style) lipgloss.Style  { return s.PaddingRight(n) }
func PaddingBottomImpl(n int, s lipgloss.Style) lipgloss.Style { return s.PaddingBottom(n) }
func PaddingLeftImpl(n int, s lipgloss.Style) lipgloss.Style   { return s.PaddingLeft(n) }

func MarginImpl(top, right, bottom, left int, s lipgloss.Style) lipgloss.Style {
	return s.Margin(top, right, bottom, left)
}
func MarginTopImpl(n int, s lipgloss.Style) lipgloss.Style    { return s.MarginTop(n) }
func MarginRightImpl(n int, s lipgloss.Style) lipgloss.Style  { return s.MarginRight(n) }
func MarginBottomImpl(n int, s lipgloss.Style) lipgloss.Style { return s.MarginBottom(n) }
func MarginLeftImpl(n int, s lipgloss.Style) lipgloss.Style   { return s.MarginLeft(n) }
func MarginBackgroundImpl(c color.Color, s lipgloss.Style) lipgloss.Style {
	return s.MarginBackground(c)
}

func WidthImpl(n int, s lipgloss.Style) lipgloss.Style     { return s.Width(n) }
func HeightImpl(n int, s lipgloss.Style) lipgloss.Style    { return s.Height(n) }
func MaxWidthImpl(n int, s lipgloss.Style) lipgloss.Style  { return s.MaxWidth(n) }
func MaxHeightImpl(n int, s lipgloss.Style) lipgloss.Style { return s.MaxHeight(n) }

var (
	Top    = lipgloss.Top
	Bottom = lipgloss.Bottom
	Left   = lipgloss.Left
	Right  = lipgloss.Right
	Center = lipgloss.Center
)

func AlignImpl(horizontal, vertical lipgloss.Position, s lipgloss.Style) lipgloss.Style {
	return s.Align(horizontal, vertical)
}
func AlignHorizontalImpl(p lipgloss.Position, s lipgloss.Style) lipgloss.Style {
	return s.AlignHorizontal(p)
}
func AlignVerticalImpl(p lipgloss.Position, s lipgloss.Style) lipgloss.Style {
	return s.AlignVertical(p)
}

var (
	NormalBorder   = lipgloss.NormalBorder()
	RoundedBorder  = lipgloss.RoundedBorder()
	ThickBorder    = lipgloss.ThickBorder()
	DoubleBorder   = lipgloss.DoubleBorder()
	BlockBorder    = lipgloss.BlockBorder()
	HiddenBorder   = lipgloss.HiddenBorder()
	AsciiBorder    = lipgloss.ASCIIBorder()
	MarkdownBorder = lipgloss.MarkdownBorder()
)

func BorderImpl(b lipgloss.Border, top, right, bottom, left bool, s lipgloss.Style) lipgloss.Style {
	return s.Border(b, top, right, bottom, left)
}
func BorderForegroundImpl(c color.Color, s lipgloss.Style) lipgloss.Style {
	return s.BorderForeground(c)
}
func BorderBackgroundImpl(c color.Color, s lipgloss.Style) lipgloss.Style {
	return s.BorderBackground(c)
}

func PlaceImpl(width, height int, horizontal, vertical lipgloss.Position, text string) string {
	return lipgloss.Place(width, height, horizontal, vertical, text)
}
func PlaceHorizontalImpl(width int, p lipgloss.Position, text string) string {
	return lipgloss.PlaceHorizontal(width, p, text)
}
func PlaceVerticalImpl(height int, p lipgloss.Position, text string) string {
	return lipgloss.PlaceVertical(height, p, text)
}

// PlaceOverlayImpl draws overlay over base with its top left corner at column, row,
// keeping the styles of both.
func PlaceOverlayImpl(column, row int, overlay, base string) string {
	baseLines := strings.Split(base, "\n")
	overlayWidth := lipgloss.Width(overlay)
	for i, line := range strings.Split(overlay, "\n") {
		at := row + i
		if at < 0 || at >= len(baseLines) {
			continue
		}
		under := baseLines[at]
		before := ansi.Truncate(under, column, "")
		if gap := column - lipgloss.Width(before); gap > 0 {
			before += strings.Repeat(" ", gap)
		}
		baseLines[at] = before + line + ansi.TruncateLeft(under, column+overlayWidth, "")
	}
	return strings.Join(baseLines, "\n")
}

func JoinHorizontalImpl(p lipgloss.Position, texts []string) string {
	return lipgloss.JoinHorizontal(p, texts...)
}
func JoinVerticalImpl(p lipgloss.Position, texts []string) string {
	return lipgloss.JoinVertical(p, texts...)
}
