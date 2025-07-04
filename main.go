package main


import (
	"fmt"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)

const (
	SCREEN_WIDTH = 10
	SCREEN_HEIGHT = 10
)


func moveCursorUp(lines int) {
	fmt.Printf("\033[%vA", lines)
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

	moveCursorUp(SCREEN_HEIGHT)
}


func triTest() {
	t := rasterizer.CreateTriangle(rasterizer.CreateVec2(2.1, 2.1), rasterizer.CreateVec2(2.1, 2.1), rasterizer.CreateVec2(2.1, 2.1))

	k := t.PointInTri(rasterizer.CreateVec2(5, 4))

	fmt.Printf("%t\n", k)

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


func main() {
	triTest()
}