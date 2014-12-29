package shapes

import (
	"math"
)

type SphereShape struct {
	position Vector
	material IMaterial
	radius   float64
}

func (ss *SphereShape) Position() Vector {
	return ss.position
}

func (ss *SphereShape) SetPosition(v Vector) {
	ss.position = v
}

func (ss *SphereShape) Material() IMaterial {
	return ss.material
}

func (ss *SphereShape) SetMaterial(IMaterial material) {
	ss.material = material
}

func (ss *SphereShape) Intersect(ray Ray) IntersectionInfo {
	info = IntersectionInfo{}
  info.SetElement(ss)

  dst := ray.Position().subtract(ss.Position())
  b := dst.Dot(ray.Direction())
  c := dst.Dot(dst).subtract(radius * radius)
  d = b * b - c

  if (d > 0) {
  	// yes, that's it, we found the intersection!
    info.SetIsHit(true)
    info.SetDistance(-b - math.Sqrt(d))
    info.SetPosition(ray.Position() + ray.Direction() * info.Distance())
    info.SetNormal(info.Position().subtract(ss.position).Normalize())

    if ss.material.HasTexture() {
      vn := Vector{0, 1, 0}.Normalize() // north pole / up
      ve := Vector{0, 0, 1}.Normalize() // equator / sphere orientation
      vp := info.Position().subtract(ss.position).Normalize() // points from center of sphere to intersection 
      
      phi = math.Acos(-vp.Dot(vn))
      v = (phi*2 / math.PI) - 1

      sinphi = ve.Dot(vp) / math.Sin(phi)
      sinphi = sinphi < -1 ? -1 : sinphi > 1 ? 1 : sinphi
      theta := math.Acos(sinphi) * 2 / math.PI

      u := 0.0

      if vn.Cross(ve).Dot(vp) > 0 {
        u = theta;
      } else {
        u = 1 - theta;
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
