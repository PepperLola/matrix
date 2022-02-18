package patterns

import (
  "fmt"
  "strings"
  "strconv"
  "golang.org/x/term"
  "github.com/PepperLola/matrix/internal/util"
)

var cracked []bool = []bool{}

var defaultAlphabet []string = []string{"0","1","2","3","4","5","6","7","8","9"}

// StartCode starts a matrix code cracking animation
func StartCode(alphabetString string, text string, defaultCooldown int) {
	width, height, err := term.GetSize(0)
  if err != nil {
    panic(err)
  }

  var alphabet []string
  if len(alphabetString) == 0 {
    alphabet = defaultAlphabet
  } else {
    alphabet = strings.Split(alphabetString, "")
  }

  cooldown = 20

  for i := 0; i < width; i++ {
    if len(text) == 0{
      text += strconv.Itoa(util.RandomBetween(0, 9))
    }
    cracked = append(cracked, false)
  }


  fmt.Print(util.ApplyForegroundColor("", util.GREEN))
  util.HideCursor()

  for true {
    if cooldown <= 0 {
      cracked[util.RandomBetween(0, len(cracked))] = true
      cooldown = defaultCooldown
    } else {
      cooldown --
    }
    for x := 1; x < width; x++ {
      if cracked[x - 1] {
        util.CursorPos(x, 1)
        fmt.Printf("%c", text[(x - 1) % len(text)])
        for y := 2; y < height; y++ {
          util.CursorPos(x, y)
          fmt.Print(" ")
        }
      } else {
        for y := 1; y < height; y++ {
          util.CursorPos(x, y)
          rand := util.RandomBetween(0, len(alphabet))
          fmt.Println(alphabet[rand])
        }
      }
    }
  }
}
