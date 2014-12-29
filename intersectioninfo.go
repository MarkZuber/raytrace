package raytrace

import ()

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

func (i *IntersectionInfo) HitCount() int {
	return i.hitCount
}

func (i *IntersectionInfo) SetHitCount(value int) {
	i.hitCount = value
}

func (i *IntersectionInfo) Element() IShape {
	return i.element
}

func (i *IntersectionInfo) Position() Vector {
	return i.position
}

func (i *IntersectionInfo) Normal() Vector {
	return i.normal
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
