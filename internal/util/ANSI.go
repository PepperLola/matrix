package util

import (
	"fmt"
)

// ANSIColor is a color in an ANSI format
type ANSIColor int

const (
  // BLACK ANSI color
	BLACK ANSIColor = iota
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
	// BRIGHTBLACK ANSI color
	BRIGHTBLACK
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
	// DEFAULT ANSI color
	DEFAULT
	// RESET ANSI color
	RESET
)

var colors = map[ANSIColor]string{
	BLACK:          "30",
	RED:            "31",
	GREEN:          "32",
	YELLOW:         "33",
	BLUE:           "34",
	MAGENTA:        "35",
	CYAN:           "36",
	WHITE:          "37",
	BRIGHTBLACK:   "90",
	BRIGHTRED:     "91",
	BRIGHTGREEN:   "92",
	BRIGHTYELLOW:  "93",
	BRIGHTBLUE:    "94",
	BRIGHTMAGENTA: "95",
	BRIGHTCYAN:    "96",
	BRIGHTWHITE:   "97",
	DEFAULT:        "39",
	RESET:          "0",
}

// ApplyColor applies an ANSI color to given string by prepending the color code.
func ApplyColor(text string, color ANSIColor) string {
	return "\033[" + colors[color] + "m" + text
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

func SaveCursor() {
	fmt.Print("\033[s")
}

func RestoreCursor() {
	fmt.Print("\033[u")
}

func CursorHome() {
	fmt.Print("\033[H")
}

func CursorPos(c int, r int) {
	fmt.Print("\033[" + fmt.Sprint(r) + ";" + fmt.Sprint(c) + "H")
}

func CursorUp(n int) {
	fmt.Print("\033[" + fmt.Sprint(n) + "A")
}

func CursorDown(n int) {
	fmt.Print("\033[" + fmt.Sprint(n) + "B")
}

func CursorRight(n int) {
	fmt.Print("\033[" + fmt.Sprint(n) + "C")
}

func CursorLeft(n int) {
	fmt.Print("\033[" + fmt.Sprint(n) + "D")
}
