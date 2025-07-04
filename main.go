package main


import (
	"github.com/Ben-Edwards44/Ascii-Rasterizer/mesh"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
)

const (
	SCREEN_WIDTH = 100
	SCREEN_HEIGHT = 40
)


func triInPixel(pixel_x int, pixel_y int, tris []rasterizer.Triangle) (bool, rasterizer.Triangle) {
	point := vector.Vec2{X: float64(pixel_x), Y: float64(pixel_y)}

	for _, i := range tris {
		if i.PointInTri(point) {return true, i}
	}

	return false, rasterizer.Triangle{}
}


func otherTest() {
	theta := 0.0
	sun_dir := vector.Vec3{1, 0, 0}
	for {
		theta += 0.001
		tris := mesh.ParseModel("models/cube.obj", theta, theta, theta)

		var screen [SCREEN_HEIGHT][SCREEN_WIDTH]pixel

		for i := 0; i < SCREEN_HEIGHT; i++ {
			for x := 0; x < SCREEN_WIDTH; x++ {
				p := pixel{0, 0, 0}

				hits, tri := triInPixel(x, i, tris)
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

	moveCursor(SCREEN_HEIGHT, false)
}