package main

type Material struct {
	color      Color
	reflective int64
}

func NewMaterial(color Color, reflective int64) *Material {
	return &Material{color, reflective}
}
