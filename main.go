package main

import (
	"github.com/Ben-Edwards44/Ascii-Rasterizer/mesh"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
)


func triInPixel(pixel_x int, pixel_y int, tris []rasterizer.Triangle) (bool, rasterizer.Triangle) {
	point := vector.Vec2{X: float64(pixel_x), Y: float64(pixel_y)}

	hit := false
	var nearest_tri rasterizer.Triangle

	for _, i := range tris {
		if i.PointInTri(point) {
			if !hit || nearest_tri.GetWorldCenter().Z < nearest_tri.GetWorldCenter().Z {
				nearest_tri = i
			}

			hit = true
		}
	}

	return hit, nearest_tri
}


func otherTest() {
	theta := 0.05
	sun_dir := vector.Vec3{1, 0, 0}
	model := mesh.ParseModel("models/cube.obj")
	model.Translate(vector.Vec3{0, 0, 4})

	for {
		model.Translate(vector.Vec3{0, 0, -4})
		model.Rotate(theta, theta, theta)
		model.Translate(vector.Vec3{0, 0, 4})

		var screen [rasterizer.SCREEN_HEIGHT][rasterizer.SCREEN_WIDTH]pixel

		for i := 0; i < rasterizer.SCREEN_HEIGHT; i++ {
			for x := 0; x < rasterizer.SCREEN_WIDTH; x++ {
				p := pixel{0, 0, 0}

				hits, tri := triInPixel(x, i, model.Triangles)
				if hits {
					normal := tri.GetNormal()
					light := (1 + vector.Dot3(&sun_dir, &normal)) * 0.5

					p.r = int(255 * light)
					p.g = int(255 * light)
					p.b = int(255 * light)
				}

				screen[i][x] = p
			}
		}

		printScreen(screen)
	}
}


func main() {
	//triTest()
	otherTest()

	//moveCursor(rasterizer.SCREEN_HEIGHT, false)
}