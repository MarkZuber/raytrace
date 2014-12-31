package raytrace

import ()

type Material struct {
	gloss        float64
	transparency float64
	reflection   float64
	refraction   float64
	color        DoubleColor
}

func CreateMaterial(color DoubleColor, gloss float64, transparency float64, reflection float64, refraction float64) *Material {
	return &Material{color: color, gloss: gloss, transparency: transparency, reflection: reflection, refraction: refraction}
}

func (sm *Material) Gloss() float64 {
	return sm.gloss
}

func (sm *Material) SetGloss(value float64) {
	sm.gloss = value
}

func (sm *Material) Transparency() float64 {
	return sm.transparency
}

func (sm *Material) SetTransparency(value float64) {
	sm.transparency = value
}

func (sm *Material) Reflection() float64 {
	return sm.reflection
}

func (sm *Material) SetReflection(value float64) {
	sm.reflection = value
}

func (sm *Material) Refraction() float64 {
	return sm.refraction
}

func (sm *Material) SetRefraction(value float64) {
	sm.refraction = value
}

func (sm *Material) HasTexture() bool {
	return false
}

func (sm *Material) GetColor(u float64, v float64) DoubleColor {
	return sm.color
}
