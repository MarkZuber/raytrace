package raytrace

import (
	"image/color"
	"math"
)

type DoubleColor struct {
	r float64
	g float64
	b float64
}

func (dc DoubleColor) R() float64 {
	return dc.r
}

func (dc DoubleColor) G() float64 {
	return dc.g
}

func (dc DoubleColor) B() float64 {
	return dc.b
}

// implement the Color interface from image/color
func (dc DoubleColor) RGBA() (r, g, b, a uint32) {
	return dc.ToRGBA().RGBA()
}

func (dc DoubleColor) ToRGBA() color.RGBA {
	rgba := color.RGBA{uint8(dc.r * 255), uint8(dc.g * 255), uint8(dc.b * 255), 0}
	return rgba
}

func (dc DoubleColor) Add(c2 DoubleColor) DoubleColor {
	return DoubleColor{dc.r + c2.r, dc.g + c2.g, dc.b + c2.b}
}

func (dc DoubleColor) Subtract(c2 DoubleColor) DoubleColor {
	return DoubleColor{dc.r - c2.r, dc.g - c2.g, dc.b - c2.b}
}

func (dc DoubleColor) MultiplyColor(c2 DoubleColor) DoubleColor {
	return DoubleColor{dc.r * c2.r, dc.g * c2.g, dc.b * c2.b}
}

func (dc DoubleColor) MultiplyFloat(f float64) DoubleColor {
	return DoubleColor{dc.r * f, dc.g * f, dc.b * f}
}

func (dc DoubleColor) Divide(f float64) DoubleColor {
	return DoubleColor{dc.r / f, dc.g / f, dc.b / f}
}

func (dc *DoubleColor) Limit() {
	dc.r = math.Min(math.Max(dc.r, 0.0), 1.0)
	dc.g = math.Min(math.Max(dc.g, 0.0), 1.0)
	dc.b = math.Min(math.Max(dc.b, 0.0), 1.0)
}

func (dc *DoubleColor) ToBlack() {
	dc.r = 0.0
	dc.g = 0.0
	dc.b = 0.0
}

func (dc DoubleColor) Blend(other DoubleColor, weight float64) DoubleColor {
	newColor := dc.MultiplyFloat(1.0 - weight).Add(other.MultiplyFloat(weight))
	return newColor
}
