package shapes

import (
	"github.com/MarkZuber/raytrace"
	"math"
)

type SphereShape struct {
	Shape
	radius float64
}

func (ss *SphereShape) Intersect(ray raytrace.Ray) raytrace.IntersectionInfo {
	info := raytrace.IntersectionInfo{}
	info.SetElement(ss)

	dst := ray.Position().Subtract(ss.Position())
	b := dst.Dot(ray.Direction())
	c := dst.Dot(dst) - (ss.radius * ss.radius)
	d := b*b - c

	if d > 0 {
		// yes, that's it, we found the intersection!
		info.SetIsHit(true)
		info.SetDistance(-b - math.Sqrt(d))
		info.SetPosition(ray.Position().Add(ray.Direction().MultiplyFloat(info.Distance())))
		info.SetNormal(info.Position().Subtract(ss.position).Normalize())

		if ss.material.HasTexture() {
			vn := raytrace.CreateVector(0, 1, 0).Normalize()        // north pole / up
			ve := raytrace.CreateVector(0, 0, 1).Normalize()        // equator / sphere orientation
			vp := info.Position().Subtract(ss.position).Normalize() // points from center of sphere to intersection

			phi := math.Acos(-vp.Dot(vn))
			v := (phi * 2 / math.Pi) - 1

			sinphi := ve.Dot(vp) / math.Sin(phi)

			sinphi = math.Min(math.Max(sinphi, -1), 1)
			theta := math.Acos(sinphi) * 2 / math.Pi

			u := 0.0

			if vn.Cross(ve).Dot(vp) > 0 {
				u = theta
			} else {
				u = 1 - theta
			}

			// alternative but worse implementation
			// double u = Math.Atan2(vp.x, vp.z);
			// double v = Math.Acos(vp.y);
			info.SetColor(ss.Material().GetColor(u, v))
		} else {
			// skip uv calculation, just get the color
			info.SetColor(ss.Material().GetColor(0, 0))
		}
	} else {
		info.SetIsHit(false)
	}

	return info
}
