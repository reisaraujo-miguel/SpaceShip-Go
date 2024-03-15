package main

import (
	"math"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func get_shaders() (vertex uint32, fragment uint32) {
	vertex_code := `
	attribute vec2 position;
	uniform mat4 mat_transformation;

	void main() {
		gl_Position = mat_transformation * vec4(position, 0.0, 1.0);
	}
	`

	fragment_code := `
	void main() {
		gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
	}
	`

	c_vertex_code, c_vertex_free := gl.Strs(vertex_code)
	c_fragment_code, c_fragment_free := gl.Strs(fragment_code)

	defer c_vertex_free()
	defer c_fragment_free()

	new_vertex := gl.CreateShader(gl.VERTEX_SHADER)
	new_fragment := gl.CreateShader(gl.FRAGMENT_SHADER)

	gl.ShaderSource(new_vertex, 1, c_vertex_code, nil)
	gl.ShaderSource(new_fragment, 1, c_fragment_code, nil)

	return new_vertex, new_fragment
}

func compile_shader(shader_obj *uint32) {
	gl.CompileShader(*shader_obj)

	var status int32
	gl.GetShaderiv(*shader_obj, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var log_length int32
		gl.GetShaderiv(*shader_obj, gl.INFO_LOG_LENGTH, &log_length)

		log := strings.Repeat("\x00", int(log_length))

		gl.GetShaderInfoLog(*shader_obj, log_length, nil, gl.Str(log))

		print("gl.CompileShader(*shader_obj): ", log)
		panic("shader_obj compile error.")
	}
}

func link_program(program_obj *uint32) {
	gl.LinkProgram(*program_obj)

	var status int32
	gl.GetProgramiv(*program_obj, gl.LINK_STATUS, &status)

	if status == gl.FALSE {
		var log_length int32
		gl.GetShaderiv(*program_obj, gl.INFO_LOG_LENGTH, &log_length)

		log := strings.Repeat("\x00", int(log_length))

		gl.GetShaderInfoLog(*program_obj, log_length, nil, gl.Str(log))

		print("gl.LinkProgram(*program_obj): ", log)
		panic("Program link error.")
	}
}

func send_to_gpu(vertices *[]mgl32.Vec2, program *uint32) {
	var buffer uint32
	gl.GenBuffers(1, &buffer)

	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	vertex_data := gl.Ptr(*vertices)
	vertex_size := int(unsafe.Sizeof(vertices)) * len(*vertices)

	gl.BufferData(gl.ARRAY_BUFFER, vertex_size, vertex_data, gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	loc := gl.GetAttribLocation(*program, gl.Str("position\x00"))

	gl.EnableVertexAttribArray(uint32(loc))

	gl.VertexAttribPointer(uint32(loc), 2, gl.FLOAT, false, int32(unsafe.Sizeof(vertices)), nil)
}

func apply_transformation(s_x float32, s_y float32, rot float32, t_x float32, t_y float32) mgl32.Mat4 {
	scale := mgl32.Mat4{
		s_x, 0.0, 0.0, 0.0,
		0.0, s_y, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	cos := float32(math.Cos(float64(-rot)))
	sin := float32(math.Sin(float64(-rot)))

	rotation := mgl32.Mat4{
		cos, -sin, 0.0, 0.0,
		sin, cos, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	translation := mgl32.Mat4{
		1.0, 0.0, 0.0, t_x,
		0.0, 1.0, 0.0, t_y,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	mat_transformation := scale.Mul4(rotation.Mul4(translation))

	return mat_transformation
}
