package rasterizer


import "fmt"


type Triangle struct {
	a Vec2
	b Vec2
	c Vec2

	ab_out Vec2
	bc_out Vec2
	ca_out Vec2
}


func CreateTriangle(a Vec2, b Vec2, c Vec2) Triangle {
	ab := sub(b, a)
	bc := sub(c, b)
	ca := sub(a, c)

	ab_out := ab.rot90()
	bc_out := bc.rot90()
	ca_out := ca.rot90()

	return Triangle{a, b, c, ab_out, bc_out, ca_out}
}


func (tri *Triangle) PointInTri(point Vec2) bool {
	ap := sub(point, tri.a)
	bp := sub(point, tri.b)
	cp := sub(point, tri.c)

	return vecsSameDir(ap, tri.ab_out) && vecsSameDir(bp, tri.bc_out) && vecsSameDir(cp, tri.ca_out)
}


func (tri *Triangle) Print() {
	fmt.Printf("triangle: A(%v,%v), B(%v,%v), C(%v,%v)\n", 
			   tri.a.x, 
			   tri.a.y,
			   tri.b.x,
			   tri.b.y,
			   tri.c.x,
			   tri.c.y)
}