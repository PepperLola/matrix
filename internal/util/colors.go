package util

import "fmt"

type RGBColor struct {
  r int
  g int
  b int
}

type HSVColor struct {
  h int
  s int
  v int
}

type Color interface {
  ToHSV() HSVColor
  ToRGB() RGBColor
}

func CreateRGB(r int, g int, b int) RGBColor {
  return RGBColor{r, g, b}
}

func (c *RGBColor) getR() int {
  return c.r
}

func (c *RGBColor) getG() int {
  return c.g
}

func (c *RGBColor) getB() int {
  return c.b
}

func (c *RGBColor) ToHSV() HSVColor {
  return RGBtoHSV(*c)
}

func (c *RGBColor) ToRGB() RGBColor {
  return *c
}

func (c *RGBColor) ToTrueColor() string {
  return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.getR(), c.getG(), c.getB())
}

func CreateHSV(h int, s int, v int) HSVColor {
  return HSVColor{h, s, v}
}

func (c *HSVColor) getH() int {
  return c.h
}

func (c *HSVColor) getS() int {
  return c.s
}

func (c *HSVColor) getV() int {
  return c.v
}

func (c *HSVColor) ToHSV() HSVColor {
  return *c
}

func (c *HSVColor) ToRGB() RGBColor {
  return HSVtoRGB(*c)
}

func (c *HSVColor) ToTrueColor() string {
  rgb := c.ToRGB()
  return (&rgb).ToTrueColor()
}
