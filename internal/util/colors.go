package util

import "fmt"

// RGBColor represents a color with red, green, and blue components
type RGBColor struct {
	r int
	g int
	b int
}

// HSVColor represents a color with hue, saturation, and value components
type HSVColor struct {
	h int
	s int
	v int
}

// Color defines the methods each type of color needs to implement
type Color interface {
	ToHSV() HSVColor
	ToRGB() RGBColor
}

// CreateRGB generates an RGBColor from the r, g, b components
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

// ToHSV converts an RGBColor to an HSVColor
func (c *RGBColor) ToHSV() HSVColor {
	return RGBtoHSV(*c)
}

// ToRGB converts an RGBColor to an RGBColor
func (c *RGBColor) ToRGB() RGBColor {
	return *c
}

// ToTrueColor generates an ANSI Truecolor code representing an RGBColor
func (c *RGBColor) ToTrueColor() string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.getR(), c.getG(), c.getB())
}

// CreateHSV creates an HSVColor from the h, s, v components
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

// ToHSV converts an HSVColor to an HSVColor
func (c *HSVColor) ToHSV() HSVColor {
	return *c
}

// ToRGB converts an HSVColor to an RGBColor
func (c *HSVColor) ToRGB() RGBColor {
	return HSVtoRGB(*c)
}

// ToTrueColor converts an HSVColor to an RGBColor, then to an ANSI Truecolor code representing that RGBColor
func (c *HSVColor) ToTrueColor() string {
	rgb := c.ToRGB()
	return (&rgb).ToTrueColor()
}
