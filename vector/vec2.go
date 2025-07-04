package vector


type Vec2 struct {
	X float64
	Y float64
}


func (vec *Vec2) Rot90() Vec2 {
	new_x := -vec.Y
	new_y := vec.X

	return Vec2{new_x, new_y}
}


func dot2(a Vec2, b Vec2) float64 {
	return a.X * b.X + a.Y * b.Y
}


func VecsSameDir(a Vec2, b Vec2) bool {
	return dot2(a, b) > 0
}


func Sub(a Vec2, b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}