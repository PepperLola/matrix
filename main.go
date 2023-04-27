package main

import (
	"flag"
	"github.com/PepperLola/matrix/internal/patterns"
)

func main() {
	var mode string
	var path string
	var scale float64
	var interval int
	var alphabet string
      var text string
	var word bool
    var reverse bool
	var fps int
	flag.StringVar(&mode, "mode", "matrix", "Mode to run")
	flag.StringVar(&path, "path", "test.png", "File path for image")
	flag.Float64Var(&scale, "scale", -1, "Scale of image")
	flag.IntVar(&interval, "interval", 10, "Number of frames between spawning matrix lines")
	flag.StringVar(&alphabet, "alphabet", "", "Custom alphabet for matrix lines and code cracking")
  flag.StringVar(&text, "text", "", "Custom text for matrix cracking")
	flag.BoolVar(&word, "word", false, "Word mode preserves order of given alphabet in matrix mode")
    flag.BoolVar(&reverse, "reverse", false, "Reverse the alphabet in word mode")
	flag.IntVar(&fps, "fps", 10, "Number of frame updates per second")
	flag.Parse()

	switch mode {
	case "matrix":
		patterns.StartMatrix(fps, alphabet, interval, word, reverse)
		break
	case "image":
		patterns.DisplayImage(path, scale)
		break
	case "gif":
		patterns.StartGIF(fps, path, scale)
		break
  case "code":
    patterns.StartCode(alphabet, text, interval)
    break
	}
}
