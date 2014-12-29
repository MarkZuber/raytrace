package raytrace

import ()

type Scene struct {
	samplingQuality  int
	rayDepth         int
	renderDiffuse    bool
	renderHighlights bool
	renderShadow     bool
	renderReflection bool
	renderRefraction bool

	background Background
	camera     Camera
	shapes     []IShape
	lights     []Light
}

func (s Scene) SamplingQuality() int {
	return s.samplingQuality
}

func (s Scene) RayDepth() int {
	return s.rayDepth
}

func (s Scene) Camera() Camera {
	return s.camera
}

func (s Scene) Background() Background {
	return s.background
}

func (s Scene) Lights() []Light {
	return s.lights
}

func (s Scene) Shapes() []IShape {
	return s.shapes
}

func (s Scene) RenderDiffuse() bool {
	return s.renderDiffuse
}

func (s Scene) RenderHighlights() bool {
	return s.renderHighlights
}

func (s Scene) RenderReflection() bool {
	return s.renderReflection
}

func (s Scene) RenderRefraction() bool {
	return s.renderRefraction
}

func (s Scene) RenderShadow() bool {
	return s.renderShadow
}

func CreateScene() Scene {
	return Scene{
		samplingQuality:  0,
		rayDepth:         3,
		renderDiffuse:    true,
		renderHighlights: true,
		renderShadow:     true,
		renderReflection: true,
		renderRefraction: true,
		camera:           CreateCameraDefaultUp(Vector{0, 0, -5}, Vector{0, 0, 1}),
		background:       CreateBackground(DoubleColor{0, 0, .5}, 0.2)}
}
