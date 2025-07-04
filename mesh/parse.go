package mesh

import (
	"os"
	"fmt"
	"strconv"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)


const DIGITS = "0123456789.-"


func checkError(err error) {
	if err != nil {panic(err)}
}


func readFile(filename string) string {
	data, err := os.ReadFile(filename)

	checkError(err)

	return string(data)
}


func split(str string, seperator rune) []string {
	current := ""

	var splitted []string
	for _, i := range str {
		if i == seperator {
			splitted = append(splitted, current)
			current = ""
		} else {
			current += string(i)
		}
	}

	if len(current) > 0 {splitted = append(splitted, current)}

	return splitted
}


func isNum(char rune) bool {
	for _, i := range DIGITS {
		if i == char {return true}
	}

	return false
}


func appendNum(current_num string, nums []float32) []float32 {
	num, err := strconv.ParseFloat(current_num, 32)
	checkError(err)

	nums = append(nums, float32(num))

	return nums
}


func extractNums(line string) []float32 {
	current_num := ""

	var nums []float32
	for _, i := range line {
		if isNum(i) {
			current_num += string(i)
		} else if i == ' ' && len(current_num) > 0 {
			nums = appendNum(current_num, nums)
			current_num = ""
		}
	}

	nums = appendNum(current_num, nums)

	return nums
}


func extractValues(lines []string, identifier rune) [][]float32 {
	var values [][]float32
	for _, i := range lines {
		if i[0] == byte(identifier) {
			values = append(values, extractNums(i))
		}
	}

	return values
}


func projectVertices(vertices [][]float32) []rasterizer.Vec2 {
	var projected_vertices []rasterizer.Vec2
	for _, i := range vertices {
		projected_vertex := rasterizer.ConvertTo2d(i)
		projected_vertices = append(projected_vertices, projected_vertex)
	}

	return projected_vertices
}


func build_triangles(face_vertices []rasterizer.Vec2) []rasterizer.Triangle {
	if len(face_vertices) < 3 {panic("invalid number of vertices in face")}
	
	start := face_vertices[0]
	prev := face_vertices[1]

	var triangles []rasterizer.Triangle
	for _, current := range face_vertices[2:] {
		tri := rasterizer.CreateTriangle(start, prev, current)

		prev = current
		triangles = append(triangles, tri)
	}

	fmt.Printf("%v,%v\n", len(triangles), len(face_vertices))

	return triangles
}


func build_faces(lines []string, projected_vertices []rasterizer.Vec2) []rasterizer.Triangle {
	var model_triangles []rasterizer.Triangle
	for _, i := range lines {
		if i[0] != 'f' {continue}

		triplets := split(i, ' ')[1:]

		var face_vertices []rasterizer.Vec2
		for _, x := range triplets {
			vertex := split(x, '/')[0]
			vertex_inx, err := strconv.Atoi(vertex)
			checkError(err)

			//why obj files are 1-indexed I will never understand...
			face_vertices = append(face_vertices, projected_vertices[vertex_inx - 1])
		}

		face_triangles := build_triangles(face_vertices)
		model_triangles = append(model_triangles, face_triangles...)
	}

	return model_triangles
}


func ParseModel(filename string) []rasterizer.Triangle {
	file_data := readFile(filename)
	lines := split(file_data, '\n')
	vertices := extractValues(lines, 'v')
	projected_vertices := projectVertices(vertices)
	model_triangles := build_faces(lines, projected_vertices)

	return model_triangles
}