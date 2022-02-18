package util

import (
	"fmt"
	"strconv"
)

// ANSIColor is a color in an ANSI format
type ANSIColor int

// ANSIFormat is a text format in an ANSI format
type ANSIFormat int

// Defaults
const (
	// DEFAULT ANSI color
	DEFAULT ANSIFormat = 39
	// RESET ANSI color
	RESET = 0
)

// Formatting
const (
	BOLD          ANSIFormat = 1
	DIM                      = 2
	ITALIC                   = 3
	UNDERLINE                = 4
	BLINKING                 = 5
	INVERSE                  = 7
	INVISIBLE                = 8
	STRIKETHROUGH            = 9
)

// Not bright colors
const (
	// BLACK ANSI color
	BLACK ANSIColor = iota + 30
	// RED ANSI color
	RED
	// GREEN ANSI color
	GREEN
	// YELLOW ANSI color
	YELLOW
	// BLUE ANSI color
	BLUE
	// MAGENTA ANSI color
	MAGENTA
	// CYAN ANSI color
	CYAN
	// WHITE ANSI color
	WHITE
)

// Bright colors
const (
	// BRIGHTBLACK ANSI color
	BRIGHTBLACK ANSIColor = iota + 90
	// BRIGHTRED ANSI color
	BRIGHTRED
	// BRIGHTGREEN ANSI color
	BRIGHTGREEN
	// BRIGHTYELLOW ANSI color
	BRIGHTYELLOW
	// BRIGHTBLUE ANSI color
	BRIGHTBLUE
	// BRIGHTMAGENTA ANSI color
	BRIGHTMAGENTA
	// BRIGHTCYAN ANSI color
	BRIGHTCYAN
	// BRIGHTWHITE ANSI color
	BRIGHTWHITE
)

// ApplyForegroundColor applies an ANSI color to the foreground of a given string
func ApplyForegroundColor(text string, color ANSIColor) string {
	return "\033[" + strconv.Itoa(int(color)) + "m" + text
}

// ApplyBackgroundColor applies an ANSI color to the background of a given string
func ApplyBackgroundColor(text string, color ANSIColor) string {
	return "\033[" + strconv.Itoa(int(color)+10) + "m" + text
}

// ApplyFormatting applies ANSI format to text
func ApplyFormatting(text string, format ANSIFormat) string {
	return "\033[" + strconv.Itoa(int(format)) + "m" + text
}

// ResetFormatting resets formatting of text
func ResetFormatting(text string, format ANSIFormat) string {
	return "\033[" + strconv.Itoa(int(format)+20) + "m" + text
}

// ClearScreen uses an ASCII escape sequence to clear the screen
func ClearScreen() {
	fmt.Print("\033[2J")
}

// ShowCursor uses an ASCII escape sequence to show the cursor
func ShowCursor() {
	fmt.Print("\033[?25h")
}

// HideCursor uses an ASCII escape sequence to hide the cursor
func HideCursor() {
	fmt.Print("\033[?25l")
}

// SaveCursor uses an ASCII escape sequence to save the cursor position
func SaveCursor() {
	fmt.Print("\033[s")
}

// RestoreCursor uses an ASCII escape sequence to restore the cursor position
func RestoreCursor() {
	fmt.Print("\033[u")
}

// CursorHome uses an ASCII escape sequence to move the cursor to its home position, usually (1, 1), or top left
func CursorHome() {
	fmt.Print("\033[H")
}

// CursorPos uses an ASCII escape sequence to move the cursor to a set column and row (x,y). Top left is (1,1)
func CursorPos(c int, r int) {
	fmt.Print("\033[" + fmt.Sprint(r) + ";" + fmt.Sprint(c) + "H")
}

// CursorUp uses an ASCII escape sequence to move the cursor up n rows
func CursorUp(n int) {
	fmt.Print("\033[" + fmt.Sprint(n) + "A")
}

// CursorDown uses an ASCII escape sequence to move the cursor down n rows
func CursorDown(n int) {
	fmt.Print("\033[" + fmt.Sprint(n) + "B")
}

// CursorRight uses an ASCII escape sequence to move the cursor right n columns
func CursorRight(n int) {
	fmt.Print("\033[" + fmt.Sprint(n) + "C")
}

// CursorLeft uses an ASCII escape sequence to move the cursor left n columns
func CursorLeft(n int) {
	fmt.Print("\033[" + fmt.Sprint(n) + "D")
}
