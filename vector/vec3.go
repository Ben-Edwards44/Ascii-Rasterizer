package vector

import (
	"math"
)


type Vec3 struct {
	X float64
	Y float64
	Z float64
}


func CreateVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}


func matMul(mat_a [][]float64, mat_b [][]float64) [][]float64 {
	var result [][]float64
	for i := 0; i < len(mat_a); i++ {
		result = append(result, []float64{})
		
		for x := 0; x < len(mat_b[0]); x++ {
			var sum float64

			for k := 0; k < len(mat_a[0]); k++ {
				sum += mat_a[i][k] * mat_b[k][x]
			}

			result[len(result) - 1] = append(result[len(result) - 1], sum)
		}
	}

	return result
}


func (vec *Vec3) applyRot(rot_mat [][]float64) {
	vec_mat := [][]float64 {
		{vec.X},
		{vec.Y},
		{vec.Z},
	}

	rotated_mat := matMul(rot_mat, vec_mat)

	vec.X = rotated_mat[0][0]
	vec.Y = rotated_mat[1][0]
	vec.Z = rotated_mat[2][0]
}


func (vec *Vec3) RotX(angle float64) {
	rot_mat := [][]float64{
		{1, 0, 0},
		{0, math.Cos(angle), -math.Sin(angle)},
		{0, math.Sin(angle), math.Cos(angle)},
	}

	vec.applyRot(rot_mat)
}


func (vec *Vec3) RotY(angle float64) {
	rot_mat := [][]float64{
		{math.Cos(angle), 0, math.Sin(angle)},
		{0, 1, 0},
		{-math.Sin(angle), 0, math.Cos(angle)},
	}

	vec.applyRot(rot_mat)
}


func (vec *Vec3) RotZ(angle float64) {
	rot_mat := [][]float64{
		{math.Cos(angle), -math.Sin(angle), 0},
		{math.Sin(angle), math.Cos(angle), 0},
		{0, 0, 1},
	}

	vec.applyRot(rot_mat)
}