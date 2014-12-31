package raytrace

type IntersectionInfo struct {
	// indicates if the shape was hit
	isHit bool

	// counts the number of shapes that were hit
	hitCount int

	// the closest shape that was intersected
	element IShape

	// position of intersection
	position Vector

	// normal vector on intesection point
	normal Vector

	// color at intersection
	color DoubleColor

	// distance from point to screen
	distance float64
}

func (i *IntersectionInfo) IsHit() bool {
	return i.isHit
}

func (i *IntersectionInfo) SetIsHit(value bool) {
	i.isHit = value
}

func (i *IntersectionInfo) HitCount() int {
	return i.hitCount
}

func (i *IntersectionInfo) SetHitCount(value int) {
	i.hitCount = value
}

func (i *IntersectionInfo) Element() IShape {
	return i.element
}

func (i *IntersectionInfo) SetElement(value IShape) {
	i.element = value
}

func (i *IntersectionInfo) Position() Vector {
	return i.position
}

func (i *IntersectionInfo) SetPosition(value Vector) {
	i.position = value
}

func (i *IntersectionInfo) Normal() Vector {
	return i.normal
}

func (i *IntersectionInfo) SetNormal(value Vector) {
	i.normal = value
}

func (i *IntersectionInfo) Color() DoubleColor {
	return i.color
}

func (i *IntersectionInfo) SetColor(value DoubleColor) {
	i.color = value
}

func (i *IntersectionInfo) Distance() float64 {
	return i.distance
}

func (i *IntersectionInfo) SetDistance(value float64) {
	i.distance = value
}
