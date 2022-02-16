package util

import (
  "fmt"
  "os"
  "os/exec"
  "runtime"
)

type ANSIColor int

const (
  BLACK ANSIColor = iota
  RED
  GREEN
  YELLOW
  BLUE
  MAGENTA
  CYAN
  WHITE
  BRIGHT_BLACK
  BRIGHT_RED
  BRIGHT_GREEN
  BRIGHT_YELLOW
  BRIGHT_BLUE
  BRIGHT_MAGENTA
  BRIGHT_CYAN
  BRIGHT_WHITE
  DEFAULT
  RESET
)

var colors = map[ANSIColor]string {
  BLACK: "30",
  RED: "31",
  GREEN: "32",
  YELLOW: "33",
  BLUE: "34",
  MAGENTA: "35",
  CYAN: "36",
  WHITE: "37",
  BRIGHT_BLACK: "90",
  BRIGHT_RED: "91",
  BRIGHT_GREEN: "92",
  BRIGHT_YELLOW: "93",
  BRIGHT_BLUE: "94",
  BRIGHT_MAGENTA: "95",
  BRIGHT_CYAN: "96",
  BRIGHT_WHITE: "97",
  DEFAULT: "39",
  RESET: "0",
}

func ApplyColor(text string, color ANSIColor) string {
  return "\033[" + colors[color] + "m" + text
}

func ClearScreen() {
  switch (runtime.GOOS) {
    case "linux":
    case "darwin":
      cmd := exec.Command("clear")
      cmd.Stdout = os.Stdout
      cmd.Run()
      break
    case "windows":
      cmd := exec.Command("cmd", "/c", "cls")
      cmd.Stdout = os.Stdout
      cmd.Run()
      break
  }
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
