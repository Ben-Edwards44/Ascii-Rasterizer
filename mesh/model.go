package mesh


import (
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)


type Model struct {
	Triangles []rasterizer.Triangle
	Colour vector.Vec3
}


func (m *Model) Rotate(rot_x float64, rot_y float64, rot_z float64) {
	for i, tri := range m.Triangles {
		m.Triangles[i] = tri.Rotate(rot_x, rot_y, rot_z)
	}
}


func (m *Model) Translate(translation vector.Vec3) {
	for i, tri := range m.Triangles {
		m.Triangles[i] = tri.Translate(translation)
	}
}


func (m *Model) Enlarge(scale_factor float64) {
	for i, tri := range m.Triangles {
		m.Triangles[i] = tri.Enlarge(scale_factor)
	}
}