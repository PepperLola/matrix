package patterns

import (
	"github.com/PepperLola/matrix/internal/img"
)

// StartGIF starts rendering a GIF with fps, path and scale
func StartGIF(fps int, path string, scale float64) {
	image := img.OpenGif(path)
	img.DisplayGif(fps, image, scale)
}
