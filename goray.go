package main

import "fmt"
import "math"
import "image"
import "image/png"
import "os"

// TODO
// add hit boolean to rayhit
// abstract some code into more functions
// misc. refactoring
// write scene parser (JSON?, OBJ?)

func main() {
	p := Perspective{Vector{0, 0, 0}, 2.0, 2.0, 512}
	filename := "out.png"
	refl := Material{Color{0, 0, 0}, 1}
	red := Material{Color{255, 0, 0}, 0}
	blue := Material{Color{0, 0, 255}, 0}
	white := Material{Color{255, 255, 255}, 0}

	sph1 := Sphere{Vector{0, 0, -16}, 2, refl}
	sph2 := Sphere{Vector{3, -1, -14}, 1, refl}
	sph3 := Sphere{Vector{-3, -1, -14}, 1, red}

	back1 := Triangle{Vector{-8, -2, -20}, Vector{8, -2, -20}, Vector{8, 10, -20}, blue}
	back2 := Triangle{Vector{-8, -2, -20}, Vector{8, 10, -20}, Vector{-8, 10, -20}, blue}
	bot1 := Triangle{Vector{-8, -2, -20}, Vector{8, -2, -10}, Vector{8, -2, -20}, white}
	bot2 := Triangle{Vector{-8, -2, -20}, Vector{-8, -2, -10}, Vector{8, -2, -10}, white}
	right := Triangle{Vector{8, -2, -20}, Vector{8, -2, -10}, Vector{8, 10, -20}, red}

	light := Vector{3, 5, -15}

	triangles := []Triangle{}
	spheres := []Sphere{}

	triangles = append(triangles, back1)
	triangles = append(triangles, back2)
	triangles = append(triangles, bot1)
	triangles = append(triangles, bot2)
	triangles = append(triangles, right)
	spheres = append(spheres, sph1)
	spheres = append(spheres, sph2)
	spheres = append(spheres, sph3)

	scene := Scene{triangles, 5, spheres, 3, light, 1}

	fmt.Printf("Rendering Scene %v\n", scene)
	Render(scene, p, filename)

	return
}

func Render(scene Scene, p Perspective, filename string) {
	output := image.NewRGBA(image.Rect(0, 0, int(p.screen_width_pixels), int(p.screen_width_pixels)))
	output_file, err := os.Create(filename)
	if err != nil {
		// handle error
	}
	for y := 0; y < int(p.screen_width_pixels); y += 1 {
		for x := 0; x < int(p.screen_width_pixels); x += 1 {
			color := Sample(scene, p, x, y, 1)
			output.Set(x, y, color)
		}
	}
	png.Encode(output_file, output)
	output_file.Close()
}

func Sample(scene Scene, p Perspective, x int, y int, samples int64) Color {
	colors := []Color{}
	pixel_length := p.screen_width_world / p.screen_width_pixels
	for i := 0; int64(i) < samples; i += 1 {
		screen_cord := ScreenCoordRandom(int64(x), int64(y), pixel_length, -p.screen_dist)
		ray := RayGet(p, screen_cord)
		colors = append(colors, Trace(scene, ray, 0))
	}
	return AverageSamples(colors)
}

func Trace(scene Scene, ray Ray, depth int64) Color {
	hit := IntersectObjects(scene, ray)
	color := hit.mat.color
	if hit.t != math.MaxFloat64 {
		if hit.mat.reflective == 1 && depth < 10 {
			color = Trace(scene, ray.Reflect(hit), depth+1)
		} else {
			light_dir := scene.light.Sub(hit.ray.origin).Normalize()
			diffuse := LightDiffuse(hit.normal, light_dir)
			light_ray := Ray{hit.ray.origin, light_dir}
			light_hit := IntersectObjects(scene, light_ray)
			d1 := light_hit.ray.origin.Distance(scene.light)
			d2 := light_hit.ray.origin.Distance(hit.ray.origin)

			if light_hit.t != math.MaxFloat64 && d2 < d1 {
				color = color.MultiplyScalar(0.2)
			} else {
				color = color.MultiplyScalar(diffuse)
			}
		}
	}
	return color
}
