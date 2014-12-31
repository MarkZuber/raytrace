package materials

import (
	"github.com/MarkZuber/raytrace"
	"math"
)

type ChessboardMaterial struct {
	raytrace.Material
	colorOdd raytrace.DoubleColor
	density  float64
}

func CreateChessboardMaterial(
	colorEven raytrace.DoubleColor,
	colorOdd raytrace.DoubleColor,
	density float64,
	gloss float64,
	transparency float64,
	reflection float64) *ChessboardMaterial {
	mat := raytrace.CreateMaterial(colorEven, gloss, transparency, reflection, 0.5)
	return &ChessboardMaterial{*mat, colorOdd, density}
}

func (sm *ChessboardMaterial) HasTexture() bool {
	return true
}

// wraps any value up in the interval [-1,1] in a rotational manner
func wrapUp(t float64) float64 {
	t = math.Mod(t, 2.0)
	// t = t % 2.0
	if t < -1 {
		t = t + 2.0
	}
	if t >= 1 {
		t = t - 2.0
	}
	return t
}

func (cbm *ChessboardMaterial) GetColor(u float64, v float64) raytrace.DoubleColor {
	t := wrapUp(u) * wrapUp(v)
	if t < 0.0 {
		return cbm.GetBaseColor()
	} else {
		return cbm.colorOdd
	}
}
