package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	scale_mat, rotation_mat, translation_mat mgl32.Mat4
)

type Body struct {
	vertices []mgl32.Vec2
}

func create_body() (body *Body) {
	new_body := Body{}
	return &new_body
}

func (b *Body) set_vertices(vertices []mgl32.Vec2) {
	b.vertices = vertices
}

func (b *Body) rotate(angle float32) {
	cos := float32(math.Cos(float64(-angle)))
	sin := float32(math.Sin(float64(-angle)))

	rotation_mat = mgl32.Mat4{
		cos, -sin, 0.0, 0.0,
		sin, cos, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) resize(x float32, y float32) {
	scale_mat = mgl32.Mat4{
		x, 0.0, 0.0, 0.0,
		0.0, y, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) translate(x float32, y float32) {
	translation_mat = mgl32.Mat4{
		1.0, 0.0, 0.0, x,
		0.0, 1.0, 0.0, y,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) get_transformation_mat() mgl32.Mat4 {
	return scale_mat.Mul4(rotation_mat.Mul4(translation_mat))
}
