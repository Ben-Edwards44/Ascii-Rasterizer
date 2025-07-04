package rasterizer


import (
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
)

type Triangle struct {
	a vector.Vec2
	b vector.Vec2
	c vector.Vec2

	ab_out vector.Vec2
	bc_out vector.Vec2
	ca_out vector.Vec2
}


func CreateTriangle(a vector.Vec2, b vector.Vec2, c vector.Vec2) Triangle {
	ab := vector.Sub(b, a)
	bc := vector.Sub(c, b)
	ca := vector.Sub(a, c)

	ab_out := ab.Rot90()
	bc_out := bc.Rot90()
	ca_out := ca.Rot90()

	return Triangle{a, b, c, ab_out, bc_out, ca_out}
}


func (tri *Triangle) PointInTri(point vector.Vec2) bool {
	ap := vector.Sub(point, tri.a)
	bp := vector.Sub(point, tri.b)
	cp := vector.Sub(point, tri.c)

	return vector.VecsSameDir(ap, tri.ab_out) && vector.VecsSameDir(bp, tri.bc_out) && vector.VecsSameDir(cp, tri.ca_out)
}