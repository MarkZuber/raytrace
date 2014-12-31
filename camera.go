package raytrace

import (
	"fmt"
)

type Camera struct {
	position Vector
	lookAt   Vector
	equator  Vector
	up       Vector // defines the tilt of the camera
	screen   Vector // defines the center position of the viewport/screen in 3D space
}

func CreateCameraDefaultUp(position Vector, lookAt Vector) *Camera {
	return CreateCamera(position, lookAt, Vector{0.0, 1.0, 0.0})
}

func CreateCamera(position Vector, lookAt Vector, up Vector) *Camera {
	normalizedUp := up.Normalize()

	return &Camera{
		up:       normalizedUp,
		position: position,
		lookAt:   lookAt,
		equator:  lookAt.Normalize().Cross(normalizedUp),
		screen:   position.Add(lookAt)}
}

func (c *Camera) Position() Vector {
	return c.position
}

func (c *Camera) LookAt() Vector {
	return c.lookAt
}

func (c *Camera) Equator() Vector {
	return c.equator
}

func (c *Camera) Up() Vector {
	return c.up
}

func (c *Camera) Screen() Vector {
	return c.screen
}

func (c *Camera) GetRay(vx float64, vy float64) Ray {
	// pos = Screen - Up * vy - Equator * vx;
	pos := c.screen.Subtract(c.up.MultiplyFloat(vy)).Subtract(c.equator.MultiplyFloat(vx))
	dir := pos.Subtract(c.position)

	ray := Ray{position: pos, direction: dir.Normalize()}
	return ray
}

func (c *Camera) String() string {
	return fmt.Sprintf("(Camera Pos %v LookAt %v Equator %v  Up %v  Screen %v )", c.Position(), c.LookAt(), c.Equator(), c.Up(), c.Screen())
}
