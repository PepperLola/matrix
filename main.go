package main

import (
  "fmt"
  "flag"
  "github.com/PepperLola/matrix/internal/patterns"
  "github.com/PepperLola/matrix/internal/img"
)

func main() {
  var mode string
  var path string
  flag.StringVar(&mode, "mode", "matrix", "Mode to run")
  flag.StringVar(&path, "path", "test.png", "File path for image")
  flag.Parse()
  fmt.Println(mode)
  switch mode {
    case "matrix":
      patterns.TestMatrix()
      break
    case "image":
      image := img.OpenImage("test.png")
      img.DisplayImage(image)
      break
    case "gif":
      image := img.OpenGif("test.gif")
      img.DisplayGif(image, -1)
      break
  }
}
