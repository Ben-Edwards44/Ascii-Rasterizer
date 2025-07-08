package main


import (
	"math/rand"

	"github.com/Ben-Edwards44/Ascii-Rasterizer/mesh"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/vector"
	"github.com/Ben-Edwards44/Ascii-Rasterizer/rasterizer"
)


var (
	SUN_DIR = vector.Vec3{X: 1, Y: 0.2, Z: -0.3}.Normalise()
	MODEL_TRANSLATION = vector.Vec3{X: 0, Y: 0, Z: 5}
	MODEL_ENLARGEMENT = 1.5
)


func triInPixel(pixel_x int, pixel_y int, tris []rasterizer.Triangle) (bool, rasterizer.Triangle) {
	point := vector.Vec2{X: float64(pixel_x), Y: float64(pixel_y)}

	hit := false
	var nearest_tri rasterizer.Triangle

	for _, i := range tris {
		if i.PointInTri(point) {
			if !hit || nearest_tri.GetWorldCenter().Z < nearest_tri.GetWorldCenter().Z {
				nearest_tri = i
			}

			hit = true
		}
	}

	return hit, nearest_tri
}


func randAngle(max_angle float64) float64 {
	rand_f := rand.Float64()

	return rand_f * max_angle
}


func rotateModel(model *mesh.Model, translation vector.Vec3) {
	rot_x := randAngle(0.05)
	rot_y := randAngle(0.05)
	rot_z := randAngle(0.05)

	inverse_translate := translation.Mul(-1)

	model.Translate(inverse_translate)
	model.Rotate(rot_x, rot_y, rot_z)
	model.Translate(translation)
}


func renderModel(model *mesh.Model) {
	var screen [rasterizer.SCREEN_HEIGHT][rasterizer.SCREEN_WIDTH]pixel

	for i := 0; i < rasterizer.SCREEN_HEIGHT; i++ {
		for x := 0; x < rasterizer.SCREEN_WIDTH; x++ {
			p := pixel{0, 0, 0, 0}

			hits, tri := triInPixel(x, i, model.Triangles)
			if hits {
				normal := tri.GetNormal()
				light := (1 + vector.Dot3(&SUN_DIR, &normal)) * 0.5
				brightness_adjusted := model.Colour.Mul(light)

				p.r = int(brightness_adjusted.X)
				p.g = int(brightness_adjusted.Y)
				p.b = int(brightness_adjusted.Z)

				p.light = light
			}

			screen[i][x] = p
		}
	}

	printScreen(screen)
}


func spinningObject(file_path string) {
	model := mesh.ParseModel(file_path)

	model.Enlarge(MODEL_ENLARGEMENT)
	model.Translate(MODEL_TRANSLATION)

	for {
		rotateModel(&model, MODEL_TRANSLATION)
		renderModel(&model)
	}
}


func main() {
	spinningObject("models/suzanne")
	moveCursor(rasterizer.SCREEN_HEIGHT, false)
}