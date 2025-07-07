package main


import (
	"fmt"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)


const CHARS = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'."


type pixel struct {
	r int
	g int
	b int

	light float64
}


func moveCursor(lines int, move_up bool) {
	char := 'B'
	if move_up {char = 'A'}

	fmt.Printf("\033[%v%c", lines, char)
}


func setColour (r int, g int, b int) {
	fmt.Printf("\033[38;2;%v;%v;%vm", r, g, b)
}


func getChar(light float64) string {
	chosen_char := int(light * float64(len(CHARS)))
	char_inx := len(CHARS) - chosen_char - 1

	return string(CHARS[char_inx])
}


func printScreen(pixels [rasterizer.SCREEN_HEIGHT][rasterizer.SCREEN_WIDTH]pixel) {
	for _, row := range pixels {
		for _, pixel := range row {
			setColour(pixel.r, pixel.g, pixel.b)
			char := getChar(pixel.light)

			fmt.Print(char)
		}

		fmt.Print("\n")
	}

	moveCursor(rasterizer.SCREEN_HEIGHT, true)
}