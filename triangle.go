package main

import "math"

type Triangle struct {
	a, b, c Vector
	mat     Material
}

// Compute position where ray (may) intersect triangle
func (tri Triangle) Intersect(ray Ray) RayHit {
	miss := RayHit{ray, Vector{0, 0, 0}, Material{Black, 0}, math.MaxFloat64}

	t := tri.ComputeT(ray)

	// account for bump ray hits (and negative t)
	if t < 0.0001 || t == math.MaxFloat64 {
		return miss
	}

	intersect_pos := ray.HitPos(t)
	normal := tri.Normal(intersect_pos)

	return RayHit{
		ray:    Ray{intersect_pos, ray.direction},
		normal: normal,
		mat:    tri.mat,
		t:      t,
	}
}

// compute distance to intersection
func (tri Triangle) ComputeT(ray Ray) float64 {
	t := math.MaxFloat64

	xa := tri.a.x
	xb := tri.b.x
	xc := tri.c.x
	xd := ray.direction.x
	xe := ray.origin.x
	ya := tri.a.y
	yb := tri.b.y
	yc := tri.c.y
	yd := ray.direction.y
	ye := ray.origin.y
	za := tri.a.z
	zb := tri.b.z
	zc := tri.c.z
	zd := ray.direction.z
	ze := ray.origin.z

	a := xa - xb
	b := ya - yb
	c := za - zb
	d := xa - xc
	e := ya - yc
	f := za - zc
	g := xd
	h := yd
	i := zd
	j := xa - xe
	k := ya - ye
	l := za - ze

	m := (a * ((e * i) - (h * f))) + (b * ((g * f) - (d * i))) + (c * ((d * h) - (e * g)))
	t = (-((f * ((a * k) - (j * b))) + (e * ((j * c) - (a * l))) + (d * ((b * l) - (k * c)))) / m)

	gamma := (((i * ((a * k) - (j * b))) + (h * ((j * c) - (a * l))) + (g * ((b * l) - (k * c)))) / m)

	if (gamma < 0) || (gamma > 1) {
		return t
	}

	beta := ((j*((e*i)-(h*f)) + (k * ((g * f) - (d * i))) + (l * ((d * h) - (e * g)))) / m)

	if (beta < 0) || (beta > (1 - gamma)) {
		return t
	}

	return t
}

// Get triangle intersection point normal
func (tri Triangle) Normal(intersect_pos Vector) Vector {
	return intersect_pos.Sub(tri.c.Sub(tri.a).Cross(tri.b.Sub(tri.a)).Normalize()).Normalize()
}
