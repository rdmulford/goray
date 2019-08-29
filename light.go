package main

type Light struct {
	position Vector
}

func LightDiffuse(normal Vector, light_dir Vector) float64 {
	diffuse := normal.Normalize().Dot(light_dir.Normalize())
	if diffuse < 0.2 {
		diffuse = 0.2
	}
	return diffuse
}
