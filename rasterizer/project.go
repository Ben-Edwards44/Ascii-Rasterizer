package rasterizer


import "github.com/Ben-Edwards44/Ascii-Rasterizer/vector"


func convertTo2d(point vector.Vec3) vector.Vec2 {
	return vector.Vec2{X: point.X * 10 + 10, Y: point.Y * 10 + 10}
}