package main

import (
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func getShaders() (vertex uint32, fragment uint32) {
	vertexCode :=
		`
		attribute vec2 position;
		uniform mat4 matTransformation;
		void main() {
			gl_Position = matTransformation * vec4(position, 0.0, 1.0);
		}
		` + "\x00"

	fragmentCode :=
		`
		void main() {
			gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
		}
		` + "\x00"

	cVertexCode, cVertexFree := gl.Strs(vertexCode)
	cFragmentCode, cFragmentFree := gl.Strs(fragmentCode)

	defer cVertexFree()
	defer cFragmentFree()

	newVertex := gl.CreateShader(gl.VERTEX_SHADER)
	newFragment := gl.CreateShader(gl.FRAGMENT_SHADER)

	gl.ShaderSource(newVertex, 1, cVertexCode, nil)
	gl.ShaderSource(newFragment, 1, cFragmentCode, nil)

	return newVertex, newFragment
}

func compileShader(shaderObj *uint32) {
	gl.CompileShader(*shaderObj)

	var status int32
	gl.GetShaderiv(*shaderObj, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(*shaderObj, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))

		gl.GetShaderInfoLog(*shaderObj, logLength, nil, gl.Str(log))

		println("gl.CompileShader(*shaderObj): ", log)
		panic("shaderObj compile error.")
	}
}

func linkProgram(programObj *uint32) {
	gl.LinkProgram(*programObj)

	var status int32
	gl.GetProgramiv(*programObj, gl.LINK_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(*programObj, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))

		gl.GetProgramInfoLog(*programObj, logLength, nil, gl.Str(log))

		println("gl.LinkProgram(*programObj): ", log)
		panic("Program link error.")
	}
}

func sendToGpu(vertices *[]mgl32.Vec2, program *uint32) {
	var buffer uint32
	gl.GenBuffers(1, &buffer)

	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	vertexData := gl.Ptr(*vertices)
	vertexSize := int(unsafe.Sizeof(vertices)) * len(*vertices)

	gl.BufferData(gl.ARRAY_BUFFER, vertexSize, vertexData, gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	loc := gl.GetAttribLocation(*program, gl.Str("position\x00"))

	gl.EnableVertexAttribArray(uint32(loc))

	gl.VertexAttribPointer(uint32(loc), 2, gl.FLOAT, false, int32(unsafe.Sizeof(vertices)), nil)
}
