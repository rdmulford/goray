package main

import "math"

type RayHit struct {
	ray    Ray
	normal Vector
	mat    Material
	t      float64
}

// find the closest hit from a list of hits
func RayHitNearest(hits []RayHit) RayHit {
	t := math.MaxFloat64
	var best_hit RayHit
	for _, hit := range hits {
		if hit.t <= t {
			t = hit.t
			best_hit = hit
		}
	}
	return best_hit
}
