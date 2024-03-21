package main

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var ()

type Body struct {
	scale_mat, rotation_mat, translation_mat mgl32.Mat4
	position_x, position_y                   float32
	scale_x, scale_y                         float32
	rotation                                 float32
	start_vertice, vertices_count            int32
}

func create_body() (body *Body) {
	new_body := Body{}

	new_body.rotation_mat = mgl32.Mat4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	new_body.scale_mat = mgl32.Mat4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	new_body.translation_mat = mgl32.Mat4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	return &new_body
}

func (b *Body) instantiate(body_vertices []mgl32.Vec2, global_vertices *[]mgl32.Vec2) {
	b.start_vertice = int32(len(*global_vertices))
	b.vertices_count = int32(len(body_vertices))
	*global_vertices = append(*global_vertices, body_vertices...)
}

func (b *Body) rotate(angle float32) {
	b.rotation = angle

	cos := float32(math.Cos(float64(-angle)))
	sin := float32(math.Sin(float64(-angle)))

	b.rotation_mat = mgl32.Mat4{
		cos, -sin, 0.0, 0.0,
		sin, cos, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) scale(x float32, y float32) {
	b.scale_x = x
	b.scale_y = y

	b.scale_mat = mgl32.Mat4{
		x, 0.0, 0.0, 0.0,
		0.0, y, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) translate(x float32, y float32) {
	b.position_x = x
	b.position_y = y

	b.translation_mat = mgl32.Mat4{
		1.0, 0.0, 0.0, x,
		0.0, 1.0, 0.0, y,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) draw_body(program *uint32) {
	mat_transformation := b.scale_mat.Mul4(b.rotation_mat.Mul4(b.translation_mat))

	loc := gl.GetUniformLocation(*program, gl.Str("mat_transformation\x00"))

	gl.UniformMatrix4fv(loc, 1, true, &mat_transformation[0])

	gl.DrawArrays(gl.TRIANGLE_STRIP, b.start_vertice, b.vertices_count)
}

func move_towards_mouse(window *glfw.Window, x *float32, y *float32, inc float32) {
	width, height := window.GetFramebufferSize()

	cursor_x, cursor_y := window.GetCursorPos()
	origin_x, origin_y := (width / 2), (height / 2)

	cat_x := (cursor_x - float64(origin_x)) - float64(*x*float32(origin_x))
	cat_y := (float64(origin_y) - cursor_y) - float64(*y*float32(origin_y))

	// Geometry rules
	vector_magnitude := math.Sqrt(math.Pow(cat_x, 2) + math.Pow(cat_y, 2))

	x_component := cat_x / vector_magnitude
	y_component := cat_y / vector_magnitude

	*x += inc * float32(x_component)
	*y += inc * float32(y_component)
}

func screen_wrap(x *float32, y *float32) {
	if *x > 1 {
		*x = -1
	} else if *x < -1 {
		*x = 1
	}

	if *y > 1 {
		*y = -1
	} else if *y < -1 {
		*y = 1
	}
}
