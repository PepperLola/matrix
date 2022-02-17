package img

import (
	"bufio"
	"fmt"
	"github.com/PepperLola/matrix/internal/util"
	"golang.org/x/image/draw"
	"golang.org/x/term"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"math"
	"os"
	"time"
)

var frameCache []image.Image = []image.Image{}

var width, height, err = term.GetSize(0)

func OpenImage(path string) image.Image {
	fSrc, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fSrc.Close()
	src, err := png.Decode(fSrc)
	if err != nil {
		panic(err)
	}

	return src
}

func OpenGif(path string) gif.GIF {
	fSrc, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fSrc.Close()

	r := bufio.NewReader(fSrc)

	gif, err := gif.DecodeAll(r)
	if err != nil {
		panic(err)
	}

	return *gif
}

func ResizeImage(src image.Image, amount float64) image.Image {
	imgWidth, imgHeight := src.Bounds().Max.X-src.Bounds().Min.X, src.Bounds().Max.Y-src.Bounds().Min.Y
	if amount == -1 {
		// scale to fit
		amount = math.Min(float64(width)/float64(imgWidth), float64(height)/float64(imgHeight))
	}
	dst := image.NewRGBA(image.Rect(0, 0, int(float64(src.Bounds().Max.X)*amount), int(float64(src.Bounds().Max.Y)*amount)))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

func DisplayGif(fps int, src gif.GIF, scale float64) {
	if err != nil {
		panic(err)
	}
	util.ClearScreen()
	util.HideCursor()
	i := 0
	for src.LoopCount < 1 || i < src.LoopCount*len(src.Image) {
		frameIdx := i % len(src.Image)
		var frameImg image.Image
		if len(frameCache) <= frameIdx {
			frame := src.Image[frameIdx]
			frameImg = ResizeImage(frame, scale)
			frameCache = append(frameCache, frameImg)
		} else {
			frameImg = frameCache[frameIdx]
		}
		DisplayImage(frameImg)

		time.Sleep(time.Second / time.Duration(fps))
		i++
	}
}

func DisplayImage(src image.Image) {
	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := src.Bounds().Min.Y; y < src.Bounds().Max.Y; y++ {
		for x := src.Bounds().Min.X; x < src.Bounds().Max.X; x += 1 {
			c := color.RGBAModel.Convert(src.At(x, y)).(color.RGBA)
			g := color.GrayModel.Convert(src.At(x, y)).(color.Gray)
			level := int(math.Min(float64((g.Y/51+c.A/51)/2), float64(4))) // 51 * 5 = 255
			util.CursorPos(x*2, y)
			col := util.CreateRGB(int(c.R), int(c.G), int(c.B))
			fmt.Print((&col).ToTrueColor())
			fmt.Print(levels[level] + levels[level])
		}
	}
}
