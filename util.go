package main

import "time"
import "math/rand"

// convert world coord to screen cord
func ScreenCoordCenter(x int64, y int64, pixel_length float64, z_dist float64) Vector {
	xpos := -1 + float64(x)*pixel_length + pixel_length/2
	ypos := 1 - float64(y)*pixel_length - pixel_length/2

	return Vector{
		x: xpos,
		y: ypos,
		z: z_dist,
	}
}

func ScreenCoordRandom(x int64, y int64, pixel_length float64, z_dist float64) Vector {
	xpos := -1 + float64(x)*pixel_length + RandomRange(0, pixel_length)
	ypos := 1 - float64(y)*pixel_length - RandomRange(0, pixel_length)

	return Vector{
		x: xpos,
		y: ypos,
		z: z_dist,
	}
}

func RandomRange(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

func AverageSamples(colors []Color) Color {
	var average Color
	for _, color := range colors {
		average = average.Add(color)
	}
	result := average.DivideScalar(float64(len(colors)))
	return result
}

// intersect with each object in the scene and return the closest hit
func IntersectObjects(scene Scene, ray Ray) RayHit {
	hits := []RayHit{}
	for i := 0; i < int(scene.t_size); i += 1 {
		hit := scene.triangles[i].Intersect(ray)
		hits = append(hits, hit)
	}
	for i := 0; i < int(scene.s_size); i += 1 {
		hit := scene.spheres[i].Intersect(ray)
		hits = append(hits, hit)
	}
	hit := RayHitNearest(hits)
	return hit
}
