package main

import (
  //"github.com/PepperLola/matrix/internal/patterns"
  "github.com/PepperLola/matrix/internal/img"
)

func main() {
  //patterns.TestMatrix()
  image := img.OpenGif("test.gif")
  img.DisplayGif(image, -1)
}
