package raytrace

import (
	"fmt"
)

type Scene struct {
	samplingQuality  int
	rayDepth         int
	renderDiffuse    bool
	renderHighlights bool
	renderShadow     bool
	renderReflection bool
	renderRefraction bool

	background *Background
	camera     *Camera
	shapes     []IShape
	lights     []*Light
}

func (s *Scene) SamplingQuality() int {
	return s.samplingQuality
}

func (s *Scene) RayDepth() int {
	return s.rayDepth
}

func (s *Scene) Camera() *Camera {
	return s.camera
}

func (s *Scene) SetCamera(value *Camera) {
	s.camera = value
}

func (s *Scene) Background() *Background {
	return s.background
}

func (s *Scene) Lights() []*Light {
	return s.lights
}

func (s *Scene) AddLight(light *Light) {
	s.lights = append(s.lights, light)
}

func (s *Scene) Shapes() []IShape {
	return s.shapes
}

func (s *Scene) AddShape(shape IShape) {
	s.shapes = append(s.shapes, shape)
}

func (s *Scene) RenderDiffuse() bool {
	return s.renderDiffuse
}

func (s *Scene) RenderHighlights() bool {
	return s.renderHighlights
}

func (s *Scene) RenderReflection() bool {
	return s.renderReflection
}

func (s *Scene) RenderRefraction() bool {
	return s.renderRefraction
}

func (s *Scene) RenderShadow() bool {
	return s.renderShadow
}

func CreateScene() *Scene {
	return &Scene{
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

func (s *Scene) String() string {
	str := fmt.Sprintf("Scene: Diffuse(%v) Highlights(%v) Reflections(%v) Refraction(%v) Shadow (%v) NumLights (%d) NumShapes (%d)",
		s.RenderDiffuse(), s.RenderHighlights(), s.RenderReflection(), s.RenderRefraction(), s.RenderShadow(), len(s.Lights()), len(s.Shapes()))

	str += fmt.Sprintf(" Camera (%v)", s.Camera())
	return str
}
