package raytrace

type IMaterial interface {
	// specifies the Gloss (or shininess) of the element
	// value must be between 1 (very shiny) and 5 (matt) for a realistic effect
	Gloss() float64
	SetGloss(value float64)

	// defines the transparency of the element.
	// values must be between 0 (opaque) and 1 (fully transparent);
	Transparency() float64
	SetTransparency(value float64)

	// specifies how much light the element will reflect
	// value must be between 0 (no reflection) to 1 (total reflection/mirror)
	Reflection() float64
	SetReflection(value float64)

	// refraction index
	// specifies how the material will bend the light rays
	// value must be between <0,1] (total reflection/mirror)
	Refraction() float64
	SetRefraction(value float64)

	// indicates that the material has a texture and therefore the exact
	// u,v coordinates are to be calculated by the element
	// and passed on in the GetColor function
	HasTexture() bool

	// retrieves the actual color of the material
	// the color can change depending on the u,v coordinates in the texture map
	GetColor(u float64, v float64) DoubleColor
}
