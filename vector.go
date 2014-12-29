package raytrace

import (
	"fmt"
	"math"
)

type Vector struct {
	x float64
	y float64
	z float64
}

func CreateVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z}
}

func (v Vector) X() float64 {
	return v.x
}

func (v Vector) Y() float64 {
	return v.y
}

func (v Vector) Z() float64 {
	return v.z
}

func (v Vector) Add(w Vector) Vector {
	return Vector{x: v.x + w.x, y: v.y + w.y, z: v.z + w.z}
}

func (v Vector) Subtract(w Vector) Vector {
	return Vector{x: v.x - w.x, y: v.y - w.y, z: v.z - w.z}
}

func (v Vector) MultiplyVector(w Vector) Vector {
	return Vector{x: v.x * w.x, y: v.y * w.y, z: v.z * w.z}
}

func (v Vector) MultiplyFloat(f float64) Vector {
	return Vector{x: v.x * f, y: v.y * f, z: v.z * f}
}

func (v Vector) Divide(f float64) Vector {
	return Vector{x: v.x / f, y: v.y / f, z: v.z / f}
}

func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	return Vector{x: v.x / mag, y: v.y / mag, z: v.z / mag}
}

func (v Vector) Dot(w Vector) float64 {
	return (v.x * w.x) + (v.y * w.y) + (v.z * w.z)
}

func (v Vector) Cross(w Vector) Vector {
	return Vector{
		x: (-v.z * w.y) + (v.y * w.z),
		y: (v.z * w.x) - (v.x * w.z),
		z: (-v.y * w.x) + (v.x * w.y)}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y) + (v.z * v.z))
}

func (v *Vector) String() string {
	return fmt.Sprintf("(%.3f, %.3f, %.3f)", v.x, v.y, v.z)
}
