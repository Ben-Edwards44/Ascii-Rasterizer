package mesh


import (
	"os"
	"strconv"

	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
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


func appendNum(current_num string, nums []float64) []float64 {
	num, err := strconv.ParseFloat(current_num, 32)
	checkError(err)

	nums = append(nums, float64(num))

	return nums
}


func extractNums(line string) []float64 {
	current_num := ""

	var nums []float64
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


func extractVectors(lines []string, identifier string) []vector.Vec3 {
	var values []vector.Vec3
	for _, i := range lines {
		line_type := split(i, ' ')[0]
		if line_type == identifier {
			vertex_coords := extractNums(i)
			values = append(values, vector.CreateVec3(vertex_coords[0], vertex_coords[1], vertex_coords[2]))
		}
	}

	return values
}


func rotate(vertices []vector.Vec3, normals []vector.Vec3, rot_x float64, rot_y float64, rot_z float64) ([]vector.Vec3, []vector.Vec3) {
	var rotated_verts []vector.Vec3
	var rotated_norms []vector.Vec3

	for _, i := range vertices {
		i.RotX(rot_x)
		i.RotY(rot_y)
		i.RotZ(rot_z)

		rotated_verts = append(rotated_verts, i)
	}

	for _, i := range normals {
		i.RotX(rot_x)
		i.RotY(rot_y)
		i.RotZ(rot_z)

		rotated_norms = append(rotated_norms, i)
	}

	return rotated_verts, rotated_norms
}


func build_triangles(face_vertices []vector.Vec3, face_normal vector.Vec3) []rasterizer.Triangle {
	if len(face_vertices) < 3 {panic("invalid number of vertices in face")}
	
	start := face_vertices[0]
	prev := face_vertices[1]

	var triangles []rasterizer.Triangle
	for _, current := range face_vertices[2:] {
		tri := rasterizer.CreateTriangle(start, prev, current, face_normal)

		prev = current
		triangles = append(triangles, tri)
	}

	return triangles
}


func build_faces(lines []string, vertices []vector.Vec3, normals []vector.Vec3) []rasterizer.Triangle {
	var model_triangles []rasterizer.Triangle
	for _, i := range lines {
		if i[0] != 'f' {continue}

		triplets := split(i, ' ')[1:]

		var face_normal vector.Vec3
		var face_vertices []vector.Vec3

		for _, x := range triplets {
			indexes := split(x, '/')

			vertex_inx, err := strconv.Atoi(indexes[0])
			checkError(err)

			normal_inx, err := strconv.Atoi(indexes[2])
			checkError(err)

			//why obj files are 1-indexed I will never understand...
			face_normal = normals[normal_inx - 1]
			face_vertices = append(face_vertices, vertices[vertex_inx - 1])
		}

		//just assume the face normal is the normal of the last vertex in the face
		face_triangles := build_triangles(face_vertices, face_normal)
		model_triangles = append(model_triangles, face_triangles...)
	}

	return model_triangles
}


func ParseModel(filename string, rot_x float64, rot_y float64, rot_z float64) []rasterizer.Triangle {
	file_data := readFile(filename)
	lines := split(file_data, '\n')

	vertices := extractVectors(lines, "v")
	normals := extractVectors(lines, "vn")

	vertices, normals = rotate(vertices, normals, rot_x, rot_y, rot_z)

	model_triangles := build_faces(lines, vertices, normals)

	return model_triangles
}