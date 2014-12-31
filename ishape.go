package raytrace

type IShape interface {
	// indicates the position of the element
	Position() Vector
	SetPosition(v Vector)

	// specifies the ambient and diffuse color of the element
	Material() IMaterial
	SetMaterial(material IMaterial)

	// this method is to be implemented by each element seperately. This is the core
	// function of each element, to determine the intersection with a ray.
	Intersect(ray Ray) *IntersectionInfo
}
