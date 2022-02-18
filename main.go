package main

import (
	"flag"
	"github.com/PepperLola/matrix/internal/img"
	"github.com/PepperLola/matrix/internal/patterns"
)

func main() {
	var mode string
	var path string
	var scale float64
	var interval int
	var alphabet string
	var word bool
	var fps int
	flag.StringVar(&mode, "mode", "matrix", "Mode to run")
	flag.StringVar(&path, "path", "test.png", "File path for image")
	flag.Float64Var(&scale, "scale", -1, "Scale of image")
	flag.IntVar(&interval, "interval", 10, "Number of frames between spawning matrix lines")
	flag.StringVar(&alphabet, "alphabet", "", "Custom alphabet for matrix lines")
	flag.BoolVar(&word, "word", false, "Word mode preserves order of given alphabet in matrix mode")
	flag.IntVar(&fps, "fps", 10, "Number of frame updates per second")
	flag.Parse()

	switch mode {
	case "matrix":
		patterns.StartMatrix(fps, alphabet, interval, word)
		break
	case "image":
		image := img.OpenImage(path)
		img.DisplayImage(img.ResizeImage(image, scale))
		break
	case "gif":
		image := img.OpenGif(path)
		img.DisplayGif(fps, image, scale)
		break
	}
}
