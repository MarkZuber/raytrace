package raytrace

import ()

type Shape struct {
	position Vector
	material IMaterial
}

func CreateShape(position Vector, material IMaterial) *Shape {
	return &Shape{position, material}
}

func (s *Shape) Position() Vector {
	return s.position
}

func (s *Shape) SetPosition(v Vector) {
	s.position = v
}

func (s *Shape) Material() IMaterial {
	return s.material
}

func (s *Shape) SetMaterial(material IMaterial) {
	s.material = material
}

func (s *Shape) Intersect(ray Ray) *IntersectionInfo {
	return &IntersectionInfo{}
}
