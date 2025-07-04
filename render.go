package main


import (
	"fmt"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)


type pixel struct {
	r int
	g int
	b int
}


func moveCursor(lines int, move_up bool) {
	char := 'B'
	if move_up {char = 'A'}

	fmt.Printf("\033[%v%c", lines, char)
}


func setColour (r int, g int, b int) {
	fmt.Printf("\033[38;2;%v;%v;%vm", r, g, b)
}


func printScreen(pixels [rasterizer.SCREEN_HEIGHT][rasterizer.SCREEN_WIDTH]pixel) {
	for _, row := range pixels {
		for _, pixel := range row {
			setColour(pixel.r, pixel.g, pixel.b)
			fmt.Print("#")
		}

		fmt.Print("\n")
	}

	moveCursor(rasterizer.SCREEN_HEIGHT, true)
}