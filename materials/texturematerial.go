package materials

import (
	"github.com/MarkZuber/raytrace"
	"image"
	"math"
)

type TextureMaterial struct {
	raytrace.Material
	image   *image.RGBA
	density float64
}

func CreateTextureMaterial(
	image *image.RGBA,
	density float64,
	gloss float64,
	transparency float64,
	reflection float64,
	refraction float64) *TextureMaterial {
	mat := raytrace.CreateMaterial(
		raytrace.CreateDoubleColor(0, 0, 0),
		gloss,
		transparency,
		reflection,
		refraction)

	return &TextureMaterial{*mat, image, density}
}

func (tm *TextureMaterial) HasTexture() bool {
	return true
}

func (tm *TextureMaterial) GetColor(u float64, v float64) raytrace.DoubleColor {
	// map u, v to [0, 2]
	u = wrapUp(u*tm.density) + 1
	v = wrapUp(v*tm.density) + 1

	width := tm.image.Bounds().Dx()
	height := tm.image.Bounds().Dy()

	// calculate exact position in texture
	nu1 := u * float64(width) / 2
	nv1 := v * float64(height) / 2

	// calculate fractions
	fu := nu1 - math.Floor(nu1)
	fv := nv1 - math.Floor(nv1)
	w1 := (1.0 - fu) * (1.0 - fv)
	w2 := fu * (1.0 - fv)
	w3 := (1.0 - fu) * fv
	w4 := fu * fv

	nu2 := int(math.Floor(nu1)) % width
	nv2 := int(math.Floor(nv1)) % height
	nu3 := int(math.Floor(nu1+1)) % width
	nv3 := int(math.Floor(nv1+1)) % height

	c1 := raytrace.CreateDoubleColorFromRGBA(tm.image.At(nu2, nv2)).MultiplyFloat(w1)
	c2 := raytrace.CreateDoubleColorFromRGBA(tm.image.At(nu3, nv2)).MultiplyFloat(w2)
	c3 := raytrace.CreateDoubleColorFromRGBA(tm.image.At(nu2, nv3)).MultiplyFloat(w3)
	c4 := raytrace.CreateDoubleColorFromRGBA(tm.image.At(nu3, nv3)).MultiplyFloat(w4)

	finalColor := c1.Add(c2).Add(c3).Add(c4)
	return finalColor
}
