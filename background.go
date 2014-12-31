package raytrace

import ()

type Background struct {
	color    DoubleColor
	ambience float64
}

func (b Background) Color() DoubleColor {
	return b.color
}

func (b Background) Ambience() float64 {
	return b.ambience
}

func CreateBackground(color DoubleColor, ambience float64) *Background {
	return &Background{color, ambience}
}
