package rasterizer


type Vec2 struct {
	x float32
	y float32
}


func CreateVec2(x float32, y float32) Vec2 {
	return Vec2{x, y}
}


func (vec *Vec2) rot90() Vec2 {
	new_x := -vec.y
	new_y := vec.x

	return Vec2{new_x, new_y}
}


func dot(a Vec2, b Vec2) float32 {
	return a.x * b.x + a.y * b.y
}


func vecsSameDir(a Vec2, b Vec2) bool {
	return dot(a, b) > 0
}


func sub(a Vec2, b Vec2) Vec2 {
	return Vec2{a.x - b.x, a.y - b.y}
}