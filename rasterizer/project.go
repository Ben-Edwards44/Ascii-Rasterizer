package rasterizer


import (
	"math"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
)


const (
	SCREEN_WIDTH = 100
	SCREEN_HEIGHT = 40
	CAM_FOV = 1.5
)

var VIEW_PLANE_HEIGHT = math.Tan(CAM_FOV / 2)


func convertToViewPlane(point vector.Vec3) vector.Vec2 {
	view_plane_y := point.Y / point.Z
	fraction_up_plane := view_plane_y / VIEW_PLANE_HEIGHT

	view_plane_x := point.X / point.Z
	fraction_along_plane := view_plane_x / VIEW_PLANE_HEIGHT

	return vector.Vec2{X: fraction_along_plane * 2, Y: fraction_up_plane}
}


func convertToScreen(point vector.Vec3) vector.Vec2 {
	view_plane_point := convertToViewPlane(point)

	scale := float64(SCREEN_HEIGHT) / 2

	scaled_y := view_plane_point.Y * scale
	scaled_x := view_plane_point.X * scale

	offset_y := float64(SCREEN_HEIGHT) / 2 - scaled_y
	offset_x := float64(SCREEN_WIDTH) / 2 + scaled_x

	return  vector.Vec2{X: offset_x, Y: offset_y}
}