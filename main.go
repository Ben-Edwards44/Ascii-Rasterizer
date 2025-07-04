package main


import (
	"github.com/Ben-Edwards44/Ascii-Rasterizer/mesh"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)

const (
	SCREEN_WIDTH = 100
	SCREEN_HEIGHT = 50
)


func triInPixel(pixel_x int, pixel_y int, tris []rasterizer.Triangle) bool {
	point := vector.Vec2{X: float64(pixel_x), Y: float64(pixel_y)}

	for _, i := range tris {
		if i.PointInTri(point) {return true}
	}

	return false
}


func otherTest() {
	theta := 0.0
	for {
		theta += 0.001
		tris := mesh.ParseModel("models/cube.obj", theta, theta, theta)

		var screen [SCREEN_HEIGHT][SCREEN_WIDTH]pixel

		for i := 0; i < SCREEN_HEIGHT; i++ {
			for x := 0; x < SCREEN_WIDTH; x++ {
				p := pixel{0, 0, 0}

				if triInPixel(x, i, tris) {
					p.r = 255
					p.g = 255
					p.b = 255
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