package main


import (
	"fmt"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/mesh"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)

const (
	SCREEN_WIDTH = 10
	SCREEN_HEIGHT = 10
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
	point := rasterizer.CreateVec2(float32(pixel_x), float32(pixel_y))

	for _, i := range tris {
		if i.PointInTri(point) {return true}
	}

	return false
}


func triTest() {
	t := rasterizer.CreateTriangle(rasterizer.CreateVec2(2.1, 2.1), rasterizer.CreateVec2(2.1, 2.1), rasterizer.CreateVec2(2.1, 2.1))

	var screen [SCREEN_HEIGHT][SCREEN_WIDTH]int

	for i := 0; i < SCREEN_HEIGHT; i++ {
		for x := 0; x < SCREEN_WIDTH; x++ {
			c := 0

			if t.PointInTri(rasterizer.CreateVec2(float32(x), float32(i))) {
				c = 1
			}

			screen[i][x] = c
		}
	}

	printScreen(screen)
}


func otherTest() {
	tris := mesh.ParseModel("models/cube.obj")

	for _, i := range tris {
		i.Print()
	}

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


func main() {
	//triTest()
	otherTest()

	moveCursor(SCREEN_HEIGHT, false)
}