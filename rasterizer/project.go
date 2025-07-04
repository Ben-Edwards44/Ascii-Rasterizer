package rasterizer


func ConvertTo2d(point []float32) Vec2 {
	return CreateVec2(point[0] * 3 + 3, point[1] * 3 + 3)
}