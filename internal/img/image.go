package img

import (
  "os"
  "fmt"
  "time"
  "math"
  "bufio"
  "image"
  "image/png"
  "image/gif"
  "image/color"
  "golang.org/x/term"
  "golang.org/x/image/draw"
  "github.com/PepperLola/matrix/internal/util"
)

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
  dst := image.NewRGBA(image.Rect(0, 0, int(float64(src.Bounds().Max.X) * amount), int(float64(src.Bounds().Max.Y) * amount)))
  draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
  return dst
}

func DisplayGif(src gif.GIF, scale float64) {
  width, height, err := term.GetSize(0)
  if err != nil {
    panic(err)
  }
  fmt.Println(src.Image[0])
  util.ClearScreen()
  util.HideCursor()
  i := 0
  if scale == -1 {
    // scale to fit
    scale = math.Min(float64(width) / float64(src.Config.Width), float64(height) / float64(src.Config.Width))
  }
  for src.LoopCount < 1 || i < src.LoopCount * len(src.Image) {
    frame := src.Image[i % len(src.Image)]
    frameImg := ResizeImage(frame, scale)
    PrintFrame(frameImg)

    time.Sleep(100 * time.Millisecond)
    i++
  }
}

func PrintFrame(src image.Image) {
  levels := []string{" ", "░", "▒", "▓", "█"}

  for y := src.Bounds().Min.Y; y < src.Bounds().Max.Y; y++ {
    for x := src.Bounds().Min.X; x < src.Bounds().Max.X; x += 1 {
      c := color.RGBAModel.Convert(src.At(x, y)).(color.RGBA)
      g := color.GrayModel.Convert(src.At(x, y)).(color.Gray)
      level := int(math.Min(float64((g.Y / 51 + c.A / 51) / 2), float64(5))) // 51 * 5 = 255
      util.CursorPos(x * 2, y)
      col := util.CreateRGB(int(c.R), int(c.G), int(c.B))
      fmt.Print((&col).ToTrueColor())
      fmt.Print(levels[level] + levels[level])
    }
  }
}
