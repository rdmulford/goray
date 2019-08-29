package main

import "math"

type Sphere struct {
	center Vector
	radius float64
	mat    Material
}

// compute an intersection between ray and sphere
func (s Sphere) Intersect(ray Ray) RayHit {
	miss := RayHit{ray, Vector{0, 0, 0}, Material{Black, 0}, math.MaxFloat64}

	t := s.ComputeT(ray)

	// ignore if t is really close or no hit
	if t < 0.0001 || t == math.MaxFloat64 {
		return miss
	}

	intersect_pos := ray.HitPos(t)
	normal := s.Normal(intersect_pos)

	return RayHit{
		ray:    Ray{intersect_pos, ray.direction},
		normal: normal,
		mat:    s.mat,
		t:      t,
	}
}

// find the distance to an intersection
func (s Sphere) ComputeT(ray Ray) float64 {
	t := math.MaxFloat64
	toCenter := ray.origin.Sub(s.center)
	r := s.radius * s.radius
	e := ray.origin
	d := ray.direction
	c := s.center

	discriminant := (d.Dot(toCenter) * d.Dot(toCenter)) - d.Dot(d)*(toCenter.Dot(toCenter)-(r*r))

	var pos_t, neg_t float64
	if discriminant == 0 {
		t = (-d.Dot(e.Sub(c)) + math.Sqrt(discriminant)/d.Dot(d))
	} else if discriminant > 0 {
		pos_t = (-d.Dot(e.Sub(c))) + math.Sqrt(discriminant)/d.Dot(d)
		neg_t = (-d.Dot(e.Sub(c))) - math.Sqrt(discriminant)/d.Dot(d)

		if pos_t < 0 && neg_t < 0 {
			return t
		} else if pos_t < 0 && neg_t >= 0 {
			t = neg_t
		} else if neg_t < 0 && pos_t >= 0 {
			t = pos_t
		} else {
			if pos_t < neg_t {
				t = pos_t
			} else {
				t = neg_t
			}
		}
	} else {
		return t
	}
	return t
}

// compute the normal of an intersection
func (s Sphere) Normal(intersect_pos Vector) Vector {
	return intersect_pos.Sub(s.center).Normalize()
}
