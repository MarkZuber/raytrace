package materials

import (
	"github.com/MarkZuber/raytrace"
)

type SolidMaterial struct {
	gloss        float64
	transparency float64
	reflection   float64
	refraction   float64
	color        raytrace.DoubleColor
}

func (sm *SolidMaterial) Gloss() float64 {
	return sm.gloss
}

func (sm *SolidMaterial) SetGloss(value float64) {
	sm.gloss = value
}

func (sm *SolidMaterial) Transparency() float64 {
	return sm.transparency
}

func (sm *SolidMaterial) SetTransparency(value float64) {
	sm.transparency = value
}

func (sm *SolidMaterial) Reflection() float64 {
	return sm.reflection
}

func (sm *SolidMaterial) SetReflection(value float64) {
	sm.reflection = value
}

func (sm *SolidMaterial) Refraction() float64 {
	return sm.refraction
}

func (sm *SolidMaterial) SetRefraction(value float64) {
	sm.refraction = value
}

func (sm *SolidMaterial) HasTexture() bool {
	return false
}

func (sm *SolidMaterial) GetColor(u float64, v float64) raytrace.DoubleColor {
	return sm.color
}
