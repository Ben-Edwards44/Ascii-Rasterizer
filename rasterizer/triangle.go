package rasterizer


import (
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
)

type Triangle struct {
	world_a vector.Vec3
	world_b vector.Vec3
	world_c vector.Vec3

	screen_a vector.Vec2
	screen_b vector.Vec2
	screen_c vector.Vec2

	normal_vec vector.Vec3

	ab_out vector.Vec2
	bc_out vector.Vec2
	ca_out vector.Vec2
}


func CreateTriangle(world_a vector.Vec3, world_b vector.Vec3, world_c vector.Vec3, normal vector.Vec3) Triangle {
	a := convertToScreen(world_a)
	b := convertToScreen(world_b)
	c := convertToScreen(world_c)

	ab := vector.Sub(b, a)
	bc := vector.Sub(c, b)
	ca := vector.Sub(a, c)

	ab_out := ab.Rot90()
	bc_out := bc.Rot90()
	ca_out := ca.Rot90()

	return Triangle{world_a, world_b, world_c, a, b, c, normal, ab_out, bc_out, ca_out}
}


func (tri *Triangle) GetNormal() vector.Vec3 {
	return tri.normal_vec
}


func (tri *Triangle) GetWorldCenter() vector.Vec3 {
	x := (tri.world_a.X + tri.world_b.X + tri.world_c.X) / 3
	y := (tri.world_a.Y + tri.world_b.Y + tri.world_c.Y) / 3
	z := (tri.world_a.Z + tri.world_b.Z + tri.world_c.Z) / 3

	return vector.Vec3{X: x, Y: y, Z: z}
}


func (tri *Triangle) PointInTri(point vector.Vec2) bool {
	ap := vector.Sub(point, tri.screen_a)
	bp := vector.Sub(point, tri.screen_b)
	cp := vector.Sub(point, tri.screen_c)

	return vector.VecsSameDir(ap, tri.ab_out) && vector.VecsSameDir(bp, tri.bc_out) && vector.VecsSameDir(cp, tri.ca_out)
}