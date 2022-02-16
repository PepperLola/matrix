package util

import (
  "math"
)

func HSVtoRGB(color HSVColor) RGBColor {
  r, g, b := 0.0, 0.0, 0.0
  h, s, v := float64(color.getH()), float64(color.getS()), float64(color.getV())

  h /= 360
  s /= 100
  v /= 100

  i := math.Floor(h * 6)
  f := h * 6 - i
  p := v * (1 - s)
  q := v * (1 - f * s)
  t := v * (1 - (1 - f) * s)

  switch (int(i) % 6) {
    case 0:
      r, g, b = v, t, p
      break
    case 1:
      r, g, b = q, v, p
      break
    case 2:
      r, g, b = p, v, t
      break
    case 3:
      r, g, b = p, q, v
      break
    case 4:
      r, g, b = t, p, v
      break
    case 5:
      r, g, b = v, p, q
      break
  }

  return CreateRGB(int(r * 255), int(g * 255), int(b * 255))
}

func RGBtoHSV(color RGBColor) HSVColor {
  r := float64(color.getR()) / 255.0
  g := float64(color.getG()) / 255.0
  b := float64(color.getB()) / 255.0
  max := math.Max(r, math.Max(g, b))
  min := math.Min(r, math.Min(g, b))

  h, s, v := max, max, max
  d := max - min
  if max == 0 {
    s = 0
  } else {
    s = d / max
  }

  if max == min {
    h = 0
  } else {
    switch max {
      case r:
        h = (g - b) / d
        if g < b {
          h += 6
        }
        break
      case g:
        h = (b - r) / d + 2
        break
      case b:
        h = (r - g) / d + 4
        break
    }

    h /= 6
  }

  return CreateHSV(int(math.Round(h * 360)), int(math.Round(s * 100)), int(math.Round(v * 100)))
}
