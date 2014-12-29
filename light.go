package raytrace

import (
	"fmt"
	"math"
)

type Light struct {
	position Vector
	color    DoubleColor
	strength float64
}

func (l *Light) Position() Vector {
	return l.position
}

func (l *Light) Color() DoubleColor {
	return l.color
}

func (l *Light) Strength() float64 {
	return l.strength
}

func (l *Light) String() string {
	return fmt.Sprintf("Light %v", l.position)
}

func (l *Light) StrengthFromDistance(distance float64) float64 {
	if distance >= l.strength {
		return 0.0
	}

	return math.Pow((l.strength-distance)/l.strength, 0.2)
}
