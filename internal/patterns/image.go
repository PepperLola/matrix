package patterns

import (
	"github.com/PepperLola/matrix/internal/img"
)

// DisplayImage displays an image from path with scale
func DisplayImage(path string, scale float64) {
	image := img.OpenImage(path)
	img.DisplayImage(img.ResizeImage(image, scale))
}
