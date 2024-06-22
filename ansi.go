package ansi

import (
	"fmt"
	"strings"

	"github.com/danbrakeley/ansi/sgr"
)

// Useful references:
// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
// https://en.wikipedia.org/wiki/ANSI_escape_code
// http://www.termsys.demon.co.uk/vtansi.htm

const (
	EscRune = '\u001b'
	Esc     = string(EscRune)
	CSI     = Esc + "[" // Control Sequence Introducer

	// Move cursor to screen-relative locations

	TopLeft     = CSI + "H"
	LeftMost    = CSI + "1G" // stays on current line
	BottomRight = CSI + "32767;32767H"

	// Save/Restore cursor position

	PosSave    = CSI + "s" // saves the current cursor position
	PosRestore = CSI + "u" // restores cursor position to the last save

	// Erase commands do not move cursor, escept where noted

	EraseEOL    = CSI + "K"  // erase to end of current line
	EraseSOL    = CSI + "1K" // erase to start of current line
	EraseLine   = CSI + "2K" // erase from start to end of current line
	EraseDown   = CSI + "J"  // erase to end of current line, then everything down to bottom of screen
	EraseUp     = CSI + "1J" // erase to start of current line, then everything up to top of screen
	EraseScreen = CSI + "2J" // erases everything to background colour, then moves cursor home

	// DECTCEM commands

	ShowCursor = CSI + "?25h"
	HideCursor = CSI + "?25l"

	// DECXCPR commands

	GetCursorPos = CSI + "6n" // responds with `ESC [ <r> ; <c> R`, where <r> is row and <c> is column
)

func Up(n int) string {
	return fmt.Sprintf(CSI+"%dA", n)
}

func Down(n int) string {
	return fmt.Sprintf(CSI+"%dB", n)
}

func Right(n int) string {
	return fmt.Sprintf(CSI+"%dC", n)
}

func Left(n int) string {
	return fmt.Sprintf(CSI+"%dD", n)
}

// NextLine moves cursor to start of line n below current
// (stops at bottom of viewable area, does not cause a scroll)
func NextLine(n int) string {
	return fmt.Sprintf(CSI+"%dE", n)
}

// PrevLine moves cursor to start of line n above current
// (stops at top of viewable area)
func PrevLine(n int) string {
	return fmt.Sprintf(CSI+"%dF", n)
}

// SGR applies the sgr params (found in the sgr package) in the order specified
// (later commands may override earlier commands)
func SGR(params ...sgr.SGR) string {
	var sb strings.Builder
	sb.Grow(len(params)*4 + 2) // worst case: n 3-char params, n-1 semicolons, <esc>, '[', and 'm'
	sb.WriteString(CSI)
	for i, v := range params {
		if i != 0 {
			sb.WriteRune(';')
		}
		sb.WriteString(string(v))
	}
	sb.WriteRune('m')
	return sb.String()
}

// stand alone values to shortcut calling SGR when you only have one argument
const (
	Reset         = CSI + string(sgr.Reset) + "m"
	Bold          = CSI + string(sgr.Bold) + "m"
	Underline     = CSI + string(sgr.Underline) + "m"
	Reverse       = CSI + string(sgr.Reverse) + "m"
	UnderlineOff  = CSI + string(sgr.UnderlineOff) + "m"
	ReverseOff    = CSI + string(sgr.ReverseOff) + "m"
	FgBlack       = CSI + string(sgr.FgBlack) + "m"
	FgDarkRed     = CSI + string(sgr.FgDarkRed) + "m"
	FgDarkGreen   = CSI + string(sgr.FgDarkGreen) + "m"
	FgDarkYellow  = CSI + string(sgr.FgDarkYellow) + "m"
	FgDarkBlue    = CSI + string(sgr.FgDarkBlue) + "m"
	FgDarkMagenta = CSI + string(sgr.FgDarkMagenta) + "m"
	FgDarkCyan    = CSI + string(sgr.FgDarkCyan) + "m"
	FgLightGray   = CSI + string(sgr.FgLightGray) + "m"
	FgReset       = CSI + string(sgr.FgReset) + "m"
	BgBlack       = CSI + string(sgr.BgBlack) + "m"
	BgDarkRed     = CSI + string(sgr.BgDarkRed) + "m"
	BgDarkGreen   = CSI + string(sgr.BgDarkGreen) + "m"
	BgDarkYellow  = CSI + string(sgr.BgDarkYellow) + "m"
	BgDarkBlue    = CSI + string(sgr.BgDarkBlue) + "m"
	BgDarkMagenta = CSI + string(sgr.BgDarkMagenta) + "m"
	BgDarkCyan    = CSI + string(sgr.BgDarkCyan) + "m"
	BgLightGray   = CSI + string(sgr.BgLightGray) + "m"
	BgReset       = CSI + string(sgr.BgReset) + "m"
	FgDarkGray    = CSI + string(sgr.FgDarkGray) + "m"
	FgRed         = CSI + string(sgr.FgRed) + "m"
	FgGreen       = CSI + string(sgr.FgGreen) + "m"
	FgYellow      = CSI + string(sgr.FgYellow) + "m"
	FgBlue        = CSI + string(sgr.FgBlue) + "m"
	FgMagenta     = CSI + string(sgr.FgMagenta) + "m"
	FgCyan        = CSI + string(sgr.FgCyan) + "m"
	FgWhite       = CSI + string(sgr.FgWhite) + "m"
	BgDarkGray    = CSI + string(sgr.BgDarkGray) + "m"
	BgRed         = CSI + string(sgr.BgRed) + "m"
	BgGreen       = CSI + string(sgr.BgGreen) + "m"
	BgYellow      = CSI + string(sgr.BgYellow) + "m"
	BgBlue        = CSI + string(sgr.BgBlue) + "m"
	BgMagenta     = CSI + string(sgr.BgMagenta) + "m"
	BgCyan        = CSI + string(sgr.BgCyan) + "m"
	BgWhite       = CSI + string(sgr.BgWhite) + "m"
)
