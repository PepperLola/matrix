package main

import (
  "flag"
  "github.com/PepperLola/matrix/internal/patterns"
  "github.com/PepperLola/matrix/internal/img"
)

func main() {
  var mode string
  var path string
  var scale float64
  flag.StringVar(&mode, "mode", "matrix", "Mode to run")
  flag.StringVar(&path, "path", "test.png", "File path for image")
  flag.Float64Var(&scale, "scale", -1, "Scale of image")
  flag.Parse()

  switch mode {
    case "matrix":
      patterns.TestMatrix()
      break
    case "image":
      image := img.OpenImage("test.png")
      img.DisplayImage(img.ResizeImage(image, scale))
      break
    case "gif":
      image := img.OpenGif("test.gif")
      img.DisplayGif(image, scale)
      break
  }
}
