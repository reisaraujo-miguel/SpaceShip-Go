package main

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var ()

type Body struct {
	scaleMat, rotationMat, translationMat mgl32.Mat4
	startVertice, verticesCount           int32
}

func createBody() (body *Body) {
	newBody := Body{}

	newBody.rotationMat = mgl32.Mat4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	newBody.scaleMat = mgl32.Mat4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	newBody.translationMat = mgl32.Mat4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	return &newBody
}

func (b *Body) instantiate(bodyVertices []mgl32.Vec2, globalVertices *[]mgl32.Vec2) {
	b.startVertice = int32(len(*globalVertices))
	b.verticesCount = int32(len(bodyVertices))
	*globalVertices = append(*globalVertices, bodyVertices...)
}

func (b *Body) rotate(angle float32) {
	cos := float32(math.Cos(float64(-angle)))
	sin := float32(math.Sin(float64(-angle)))

	b.rotationMat = mgl32.Mat4{
		cos, -sin, 0.0, 0.0,
		sin, cos, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) scale(x float32, y float32) {
	b.scaleMat = mgl32.Mat4{
		x, 0.0, 0.0, 0.0,
		0.0, y, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) translate(x float32, y float32) {
	b.translationMat = mgl32.Mat4{
		1.0, 0.0, 0.0, x,
		0.0, 1.0, 0.0, y,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (b *Body) drawBody(program *uint32) {
	matTransformation := b.scaleMat.Mul4(b.rotationMat.Mul4(b.translationMat))

	loc := gl.GetUniformLocation(*program, gl.Str("matTransformation\x00"))

	gl.UniformMatrix4fv(loc, 1, true, &matTransformation[0])

	gl.DrawArrays(gl.TRIANGLE_STRIP, b.startVertice, b.verticesCount)
}

func moveTowardsMouse(window *glfw.Window, x *float32, y *float32, inc float32) {
	width, height := window.GetFramebufferSize()

	cursorX, cursorY := window.GetCursorPos()
	originX, originY := (width / 2), (height / 2)

	catX := (cursorX - float64(originX)) - float64(*x*float32(originX))
	catY := (float64(originY) - cursorY) - float64(*y*float32(originY))

	// Geometry rules
	vectorMagnitude := math.Sqrt(math.Pow(catX, 2) + math.Pow(catY, 2))

	xComponent := catX / vectorMagnitude
	yComponent := catY / vectorMagnitude

	*x += inc * float32(xComponent)
	*y += inc * float32(yComponent)
}

func screenWrap(x *float32, y *float32) {
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
