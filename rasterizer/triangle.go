package rasterizer


//import "fmt"


type triangle struct {
	a Vec2
	b Vec2
	c Vec2

	ab_out Vec2
	bc_out Vec2
	ca_out Vec2
}


func CreateTriangle(a Vec2, b Vec2, c Vec2) triangle {
	ab := sub(b, a)
	bc := sub(c, b)
	ca := sub(a, c)

	//fmt.Printf("ab: %v,%v\n", ab.x, ab.y)
	//fmt.Printf("bc: %v,%v\n", bc.x, bc.y)
	//fmt.Printf("ca: %v,%v\n", ca.x, ca.y)

	ab_out := ab.rot90()
	bc_out := bc.rot90()
	ca_out := ca.rot90()

	return triangle{a, b, c, ab_out, bc_out, ca_out}
}


func (tri *triangle) PointInTri(point Vec2) bool {
	ap := sub(point, tri.a)
	bp := sub(point, tri.b)
	cp := sub(point, tri.c)

	//fmt.Printf("out: %v,%v\n", tri.ab_out.x, tri.ab_out.y)
	//fmt.Printf("out: %v,%v\n", tri.bc_out.x, tri.bc_out.y)
	//fmt.Printf("out: %v,%v\n", tri.ca_out.x, tri.ca_out.y)

	//fmt.Printf("same: %t\n", vecsSameDir(ap, tri.ab_out))
	//fmt.Printf("same: %t\n", vecsSameDir(bp, tri.bc_out))
	//fmt.Printf("same: %t\n", vecsSameDir(cp, tri.ca_out))

	return vecsSameDir(ap, tri.ab_out) && vecsSameDir(bp, tri.bc_out) && vecsSameDir(cp, tri.ca_out)
}