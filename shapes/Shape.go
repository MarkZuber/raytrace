package shapes

import (
	"github.com/MarkZuber/raytrace"
)

type Shape struct {
	position raytrace.Vector
	material raytrace.IMaterial
}

func (s *Shape) Position() raytrace.Vector {
	return s.position
}

func (s *Shape) SetPosition(v raytrace.Vector) {
	s.position = v
}

func (s *Shape) Material() raytrace.IMaterial {
	return s.material
}

func (s *Shape) SetMaterial(material raytrace.IMaterial) {
	s.material = material
}
