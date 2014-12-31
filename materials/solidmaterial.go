package materials

import (
	"github.com/MarkZuber/raytrace"
)

type SolidMaterial struct {
	raytrace.Material
}

func CreateSolidMaterial(color raytrace.DoubleColor, gloss float64, transparency float64, reflection float64, refraction float64) *SolidMaterial {
	mat := raytrace.CreateMaterial(color, gloss, transparency, reflection, refraction)
	return &SolidMaterial{*mat}
}
