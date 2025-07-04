package rasterizer


import "github.com/Ben-Edwards44/Ascii-Rasterizer/vector"


func ConvertTo2d(point vector.Vec3) vector.Vec2 {
	return vector.Vec2{X: point.X * 3 + 3, Y: point.Y * 3 + 3}
}