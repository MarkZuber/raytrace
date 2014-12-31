package shapes

import (
	"fmt"
	"github.com/MarkZuber/raytrace"
)

type PlaneShape struct {
	raytrace.Shape
	d float64
}

func CreatePlaneShape(position raytrace.Vector, material raytrace.IMaterial, d float64) *PlaneShape {
	shape := raytrace.CreateShape(position, material)
	return &PlaneShape{*shape, d}
}

func (ps *PlaneShape) Intersect(ray raytrace.Ray) *raytrace.IntersectionInfo {
	info := &raytrace.IntersectionInfo{}
	vd := ps.Position().Dot(ray.Direction())
	if vd == 0 {
		// no intersection
		return info
	}

	t := -((ps.Position().Dot(ray.Position()) + ps.d) / vd)
	if t <= 0 {
		return info
	}

	info.SetElement(ps)
	info.SetIsHit(true)
	info.SetPosition(ray.Position().Add(ray.Direction().MultiplyFloat(t)))
	info.SetNormal(ps.Position())
	info.SetDistance(t)

	if ps.Material().HasTexture() {
		vecU := raytrace.CreateVector(ps.Position().Y(), ps.Position().Z(), -ps.Position().X())
		vecV := vecU.Cross(ps.Position())

		u := info.Position().Dot(vecU)
		v := info.Position().Dot(vecV)
		info.SetColor(ps.Material().GetColor(u, v))
	} else {
		info.SetColor(ps.Material().GetColor(0, 0))
	}

	return info
}

func (ps *PlaneShape) String() string {
	return fmt.Sprintf("(Plane Position %v D=%.2f", ps.Position(), ps.d)
}
