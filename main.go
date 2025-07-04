package main


import (
	"fmt"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/mesh"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)

const (
	SCREEN_WIDTH = 100
	SCREEN_HEIGHT = 50
)


func moveCursor(lines int, move_up bool) {
	char := 'B'
	if move_up {char = 'A'}

	fmt.Printf("\033[%v%c", lines, char)
}


func printScreen(pixels [SCREEN_HEIGHT][SCREEN_WIDTH]int) {
	for _, row := range pixels {
		for _, pixel := range row {
			if pixel == 0 {
				fmt.Print("#")
			} else {
				fmt.Print("~")
			}
		}

		fmt.Print("\n")
	}

	moveCursor(SCREEN_HEIGHT, true)
}


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

		var screen [SCREEN_HEIGHT][SCREEN_WIDTH]int

		for i := 0; i < SCREEN_HEIGHT; i++ {
			for x := 0; x < SCREEN_WIDTH; x++ {
				c := 0

				if triInPixel(x, i, tris) {
					c = 1
				}

				screen[i][x] = c
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