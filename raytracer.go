package raytrace

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type RayTracer struct {
	viewport image.Rectangle
	scene    *Scene
}

func CreateRayTracer(viewport image.Rectangle, scene *Scene) *RayTracer {
	return &RayTracer{viewport, scene}
}

func (rt *RayTracer) SimpleRender(image *image.RGBA) {
	for y := 0; y < image.Bounds().Size().Y; y++ {
		// fmt.Printf("Rendering line: %d\n", y)
		fmt.Printf(".")
		for x := 0; x < image.Bounds().Size().X; x++ {
			color := rt.GetPixelColor(x, y)
			// fmt.Printf("Setting (%d, %d) to %v\n", x, y, color)
			image.Set(x, y, color)
		}
	}
	fmt.Println()
}

func (rt *RayTracer) SetViewport(viewport image.Rectangle) {
	rt.viewport = viewport
}

func (rt *RayTracer) Viewport() image.Rectangle {
	return rt.viewport
}

func (rt *RayTracer) SetScene(scene *Scene) {
	rt.scene = scene
}

func (rt *RayTracer) Scene() *Scene {
	return rt.scene
}

func (rt *RayTracer) GetRgssOffsets(quality int) []Point {
	sampleCount := quality * quality
	samplesArray := make([]Point, sampleCount)

	if sampleCount == 1 {
		samplesArray[0] = Point{0.0, 0.0}
	} else {
		for i := 0; i < sampleCount; i++ {
			y := float64(i+1) / float64(sampleCount+1)
			x := y * float64(quality)

			x -= math.Floor(x)

			samplesArray[i] = Point{float64(x - 0.5), float64(y - 0.5)}
		}
	}

	return samplesArray
}

func BlendColors(colors []color.RGBA64) color.RGBA64 {
	if len(colors) == 0 {
		return color.RGBA64{65535, 65535, 65535, 65535}
	}

	var rSum, gSum, bSum uint32 = 0, 0, 0

	for i := 0; i < len(colors); i++ {
		c := colors[i]
		r, g, b, _ := c.RGBA()
		rSum += r
		gSum += g
		bSum += b
	}

	r := uint16(rSum / uint32(len(colors)))
	g := uint16(gSum / uint32(len(colors)))
	b := uint16(bSum / uint32(len(colors)))

	return color.RGBA64{r, g, b, 65535}
}

func (rt *RayTracer) GetPixelColor(x int, y int) color.RGBA64 {
	xd := float64(x)
	yd := float64(y)

	// this will trigger the raytracing algorithm
	if rt.Scene().SamplingQuality() == 0 {
		xp := xd/float64(rt.viewport.Size().X)*2 - 1
		yp := yd/float64(rt.viewport.Size().Y)*2 - 1

		ray := rt.Scene().Camera().GetRay(xp, yp)
		return rt.CalculateColor(ray, rt.Scene()).ToRGBA64()
	} else {
		samples := rt.GetRgssOffsets(rt.Scene().SamplingQuality())
		colors := make([]color.RGBA64, len(samples))
		for i, sample := range samples {
			xp := (xd+sample.X)/float64(rt.viewport.Size().X)*2 - 1
			yp := (yd+sample.Y)/float64(rt.viewport.Size().Y)*2 - 1

			ray := rt.Scene().Camera().GetRay(xp, yp)
			colors[i] = rt.CalculateColor(ray, rt.Scene()).ToRGBA64()
		}
		return BlendColors(colors)
	}
}

// this implementation is used for debugging purposes.
// the color is calculated following the normal raytrace procedure
// execpt it is calculated for 1 particula ray
func (rt *RayTracer) CalculateColor(ray Ray, scene *Scene) DoubleColor {
	intersectionInfo := rt.TestIntersection(ray, scene, nil)
	if intersectionInfo.IsHit() {
		c := rt.RayTrace(intersectionInfo, ray, scene, 0)
		return c
	}

	return rt.scene.background.color
}

// This is the main RayTrace controller algorithm, the core of the RayTracer
// recursive method setup
// this does the actual tracing of the ray and determines the color of each pixel
// supports:
// - ambient lighting
// - diffuse lighting
// - Gloss lighting
// - shadows
// - reflections
func (rt *RayTracer) RayTrace(intersectionInfo *IntersectionInfo, ray Ray, scene *Scene, depth int) DoubleColor {
	// calculate ambient light
	color := intersectionInfo.Color().MultiplyFloat(scene.Background().Ambience())
	shininess := math.Pow(10, intersectionInfo.Element().Material().Gloss()+1)

	for _, light := range scene.Lights() {
		// calculate diffuse lighting
		v := light.Position().Subtract(intersectionInfo.Position()).Normalize()

		if scene.RenderDiffuse() {
			l := v.Dot(intersectionInfo.Normal())
			if l > 0.0 {
				color = color.Add(intersectionInfo.Color().MultiplyColor(light.Color()).MultiplyFloat(l))
			}
		}

		// this is the max depth of raytracing.
		// increasing depth will calculate more accurate color, however it will
		// also take longer (exponentially)
		if depth < scene.RayDepth() {
			// calculate reflection ray
			if scene.RenderReflection() && intersectionInfo.Element().Material().Reflection() > 0 {
				reflectionray := rt.GetReflectionRay(intersectionInfo.Position(), intersectionInfo.Normal(), ray.Direction())
				refl := rt.TestIntersection(reflectionray, scene, intersectionInfo.Element())
				if refl.IsHit() && refl.Distance() > 0 {
					// recursive call, this makes reflections expensive
					refl.SetColor(rt.RayTrace(refl, reflectionray, scene, depth+1))
				} else {
					// does not reflect an object, then reflect background color
					refl.SetColor(scene.Background().Color())
				}

				color = color.Blend(refl.Color(), intersectionInfo.Element().Material().Reflection())
			}

			// calculate refraction ray
			if scene.RenderRefraction() && intersectionInfo.Element().Material().Transparency() > 0 {
				refractionray := rt.GetRefractionRay(intersectionInfo.Position(), intersectionInfo.Normal(), ray.Direction(), intersectionInfo.Element().Material().Refraction())
				refr := intersectionInfo.Element().Intersect(refractionray)
				if refr.IsHit() {
					// refractionray = new Ray(refr.Position, ray.Direction);
					refractionray = rt.GetRefractionRay(refr.Position(), refr.Normal(), refractionray.Direction(), refr.Element().Material().Refraction())
					refr = rt.TestIntersection(refractionray, scene, intersectionInfo.Element())
					if refr.IsHit() && refr.Distance() > 0 {
						// recursive call, this makes refractions expensive
						refr.SetColor(rt.RayTrace(refr, refractionray, scene, depth+1))
					} else {
						refr.SetColor(scene.Background().Color())
					}
				} else {
					refr.SetColor(scene.Background().Color())
				}
				color = color.Blend(refr.Color(), intersectionInfo.Element().Material().Transparency())
			}
		}

		shadow := &IntersectionInfo{}
		if scene.RenderShadow() {
			// calculate shadow, create ray from intersection point to light
			shadowray := Ray{intersectionInfo.Position(), v}

			// find any element in between intersection point and light
			shadow = rt.TestIntersection(shadowray, scene, intersectionInfo.Element())
			if shadow.IsHit() && shadow.Element() != intersectionInfo.Element() {
				// only cast shadow if the found interesection is another
				// element than the current element

				// Math.Pow(.5, shadow.HitCount);
				color = color.MultiplyFloat(0.5 + 0.5*math.Pow(shadow.Element().Material().Transparency(), 0.5))
			}
		}

		// only show highlights if it is not in the shadow of another object
		if scene.RenderHighlights() && !shadow.IsHit() && intersectionInfo.Element().Material().Gloss() > 0 {
			// only show Gloss light if it is not in a shadow of another element.
			// calculate Gloss lighting (Phong)
			lv := intersectionInfo.Element().Position().Subtract(light.Position()).Normalize()
			e := scene.Camera().Position().Subtract(intersectionInfo.Element().Position()).Normalize()
			h := e.Subtract(lv).Normalize()

			glossweight := math.Pow(math.Max(intersectionInfo.Normal().Dot(h), 0), shininess)
			color = color.Add(light.Color().MultiplyFloat(glossweight))
		}
	}

	// normalize the color
	color = color.Limit()
	return color
}

// this method tests for an intersection. It will try to find the closest
// object that intersects with the ray.
// it will inspect every object in the scene. also here there is room for increased performance.
func (rt *RayTracer) TestIntersection(ray Ray, scene *Scene, exclude IShape) *IntersectionInfo {
	hitcount := 0
	best := &IntersectionInfo{}
	best.SetDistance(math.MaxFloat64)

	for _, elt := range scene.Shapes() {
		if elt == exclude {
			continue
		}

		intersectionInfo := elt.Intersect(ray)
		if intersectionInfo.IsHit() && intersectionInfo.Distance() < best.Distance() && intersectionInfo.Distance() >= 0 {
			best = intersectionInfo
			hitcount++
		}
	}
	best.SetHitCount(hitcount)
	return best
}

func (rt *RayTracer) GetReflectionRay(p Vector, n Vector, v Vector) Ray {
	c1 := -n.Dot(v)
	rl := v.Add(n.MultiplyFloat(2).MultiplyFloat(c1))
	return Ray{p, rl}
}

func (rt *RayTracer) GetRefractionRay(p Vector, n Vector, v Vector, refraction float64) Ray {
	// V = V * -1;
	// double n = -0.55; // refraction constant for now
	// if (n < 0 || n > 1) return new Ray(P, V); // no refraction

	c1 := n.Dot(v)
	c2 := 1.0 - refraction*refraction*(1-c1*c1)

	// TODO: This may be a bug.  There was originally blank space after the
	// "if (c2 < 0)" line which was probably correcting before doing the Sqrt().
	// if (c2 < 0)

	c2 = math.Sqrt(c2)
	t := (n.MultiplyFloat(refraction*c1 - c2).Subtract(v.MultiplyFloat(refraction))).MultiplyFloat(-1.0)
	t.Normalize()

	return Ray{p, t} // no refraction
}
