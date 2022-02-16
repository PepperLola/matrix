package patterns

import (
  "fmt"
  "time"
  "math"
  "math/rand"
  "golang.org/x/term"
  "github.com/PepperLola/matrix/internal/util"
)

type Line struct {
  x int
  y int
  length int
  letters []rune
}

var lines = []*Line{}
var lastLineForCol = map[int]*Line{}

var MIN_LINE_GAP = 4
var MIN_LINE_LENGTH = 30
var MAX_LINE_LENGTH = 50

var defaultCooldown = 5
var cooldown = defaultCooldown

var kana []rune = []rune{'ア','ァ','カ','サ','ナ','ハ','マ','ヤ','ラ','ガ','ザ','ダ','バ','パ','イ','ィ','キ','シ','チ','ニ','ヒ','ミ','リ','ヰ','ギ','ジ','ヂ','ウ','ゥ','ク','ス','ツ','ヌ','ユ','ュ','ル','グ','ヅ','プ','エ','ェ','ケ','セ','テ','ネ','ヘ','メ','レ','ゼ','デ','ベ','ペ','オ','ォ','コ','ノ','ヨ','ョ','ロ','ヲ','ゴ','ゾ','ド','ボ','ポ','ヴ','ッ','ン'}
var latin []rune = []rune{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z','a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}
var numbers []rune = []rune{'0','1','2','3','4','5','6','7','8','9'}
var alphabet = append(kana, append(latin, numbers...)...)

func RemoveElement(slice []*Line, s int) []*Line {
    return append(slice[:s], slice[s+1:]...)
}

func RandomLetter() rune {
  return alphabet[util.RandomBetween(0, len(alphabet))]
}

func PrintMatrix(height int) {
  for _, line := range lines {
    x := line.x
    y := line.y
    if (y - line.length > 0) {
      util.CursorPos(x, y - line.length)
      fmt.Print(" ")
    }
    line.letters = append(line.letters, RandomLetter())
    for r := y; r > y - int(math.Min(float64(len(line.letters)), float64(line.length))); r -- {
      if r < height {
        util.CursorPos(x, r)
        var color util.ANSIColor = util.BLACK
        if (y == r) {
          color = util.BRIGHT_GREEN
        } else if (y - r < (3 * line.length / 4)) {
          color = util.GREEN
        } else {
          color = util.BRIGHT_BLACK
        }
        fmt.Printf(util.ApplyColor("%c", color), line.letters[y - r])
      }
    }
  }
}

func TestMatrix() {
  width, height, err := term.GetSize(0)
  if err != nil {
    panic(err)
  }
  //util.ClearScreen()
  fmt.Print("\033[?25l")
  fmt.Print("\033[2J")

  for true {
    if cooldown <= 0 {
      cooldown = defaultCooldown
      col := rand.Intn(width)
      lastLine, contains := lastLineForCol[col]
      for contains && (*lastLine).y - (*lastLine).length < MIN_LINE_GAP {
        col = rand.Intn(width)
        lastLine, contains = lastLineForCol[col]
      }
      newLine := Line{
        col, 0, util.RandomBetween(MIN_LINE_LENGTH, MAX_LINE_LENGTH), []rune{},
      }
      lines = append(lines, &newLine)
      lastLineForCol[col] = &newLine
    } else {
      cooldown --
    }
    for i, line := range lines {
      line.y++
      if line.y - line.length > height {
        lines = RemoveElement(lines, i)
      }
    }
    // move lines down

    util.SaveCursor()
    PrintMatrix(height)
    util.RestoreCursor()
    time.Sleep(100 * time.Millisecond)
  }
}
