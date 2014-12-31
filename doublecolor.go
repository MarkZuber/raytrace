package raytrace

import (
	"fmt"
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

func (dc DoubleColor) String() string {
	return fmt.Sprintf("(DoubleColor r: %.2f  g: %.2f  b: %.2f)", dc.r, dc.g, dc.b)
}

func CreateDoubleColor(r float64, g float64, b float64) DoubleColor {
	return DoubleColor{r, g, b}
}

func CreateDoubleColorFromRGBA(color color.Color) DoubleColor {
	r, g, b, _ := color.RGBA()
	return CreateDoubleColor(float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0)
}

// implement the Color interface from image/color
func (dc DoubleColor) RGBA() (r, g, b, a uint32) {
	r = uint32(dc.r * 65535)
	g = uint32(dc.g * 65535)
	b = uint32(dc.b * 65535)
	a = 65535
	return
}

func (dc DoubleColor) ToRGBA64() color.RGBA64 {
	rgba := color.RGBA64{uint16(dc.r * 65535), uint16(dc.g * 65535), uint16(dc.b * 65535), 65535}
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

func (dc *DoubleColor) Limit() DoubleColor {
	r := math.Min(math.Max(dc.r, 0.0), 1.0)
	g := math.Min(math.Max(dc.g, 0.0), 1.0)
	b := math.Min(math.Max(dc.b, 0.0), 1.0)
	return DoubleColor{r, g, b}
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
