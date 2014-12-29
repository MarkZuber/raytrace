package raytrace

type Ray struct {
	position  Vector
	direction Vector
}

func (r *Ray) Position() Vector {
	return r.position
}

func (r *Ray) Direction() Vector {
	return r.direction
}
