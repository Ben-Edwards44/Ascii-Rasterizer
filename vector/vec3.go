package vector


import "math"


type Vec3 struct {
	X float64
	Y float64
	Z float64
}


func CreateVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}


func Add(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
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


func (vec *Vec3) applyRot(rot_mat [][]float64) Vec3 {
	vec_mat := [][]float64 {
		{vec.X},
		{vec.Y},
		{vec.Z},
	}

	rotated_mat := matMul(rot_mat, vec_mat)

	new_x := rotated_mat[0][0]
	new_y := rotated_mat[1][0]
	new_z := rotated_mat[2][0]

	return Vec3{new_x, new_y, new_z}
}


func (vec *Vec3) rotX(angle float64) Vec3 {
	rot_mat := [][]float64{
		{1, 0, 0},
		{0, math.Cos(angle), -math.Sin(angle)},
		{0, math.Sin(angle), math.Cos(angle)},
	}

	return vec.applyRot(rot_mat)
}


func (vec *Vec3) rotY(angle float64) Vec3 {
	rot_mat := [][]float64{
		{math.Cos(angle), 0, math.Sin(angle)},
		{0, 1, 0},
		{-math.Sin(angle), 0, math.Cos(angle)},
	}

	return vec.applyRot(rot_mat)
}


func (vec *Vec3) rotZ(angle float64) Vec3 {
	rot_mat := [][]float64{
		{math.Cos(angle), -math.Sin(angle), 0},
		{math.Sin(angle), math.Cos(angle), 0},
		{0, 0, 1},
	}

	return vec.applyRot(rot_mat)
}


func (vec *Vec3) Rotate(rot_x float64, rot_y float64, rot_z float64) Vec3 {
	new_vec := vec.rotX(rot_x)
	new_vec = new_vec.rotX(rot_y)
	new_vec = new_vec.rotZ(rot_z)

	return new_vec
}


func Dot3(a *Vec3, b *Vec3) float64 {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}