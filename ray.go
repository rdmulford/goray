package main

type Ray struct {
	origin    Vector
	direction Vector
}

// calculate position where intersection occured
func (r Ray) HitPos(t float64) Vector {
	return r.direction.MultiplyByScalar(t).Add(r.origin)
}

// compute reflection ray based off where the ray hit
func (r Ray) Reflect(hit RayHit) Ray {
	reflection_dir := hit.ray.direction.Sub(hit.normal.MultiplyByScalar(2.0 * hit.ray.direction.Dot(hit.normal))).Normalize()
	reflection_origin := reflection_dir.MultiplyByScalar(0.001).Add(hit.ray.origin)
	return Ray{
		origin:    reflection_origin,
		direction: reflection_dir,
	}
}

func RayGet(p Perspective, screen_cord Vector) Ray {
	return Ray{
		origin:    Vector{0, 0, 0},
		direction: screen_cord.Sub(p.camera_pos).Normalize(),
	}
}
