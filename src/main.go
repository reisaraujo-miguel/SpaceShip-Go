package main

import (
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize glfw: ", err)
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(600, 600, "Pong-Go!", nil, nil)

	if err != nil {
		log.Fatalln("Could not create window: ", err)
	}

	window.MakeContextCurrent()

	glfw.GetCurrentContext().SetKeyCallback(key_event)
	glfw.GetCurrentContext().SetMouseButtonCallback(mouse_event)

	if err := gl.Init(); err != nil {
		log.Fatalln("Could not init gl: ", err)
	}

	program := gl.CreateProgram()
	vertex, fragment := get_shaders()

	compile_shader(&vertex)
	compile_shader(&fragment)

	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)

	link_program(&program)

	gl.UseProgram(program)

	vertices := []mgl32.Vec2{
		/*// left square
		{-0.89, -0.18},
		{-0.89, 0.18},
		{-0.95, -0.18},
		{-0.95, 0.18},
		// right square
		{0.89, -0.18},
		{0.89, 0.18},
		{0.95, -0.18},
		{0.95, 0.18},
		*/
		{+0.0, +0.05},
		{+0.05, -0.05},
		{-0.05, -0.05},
	}

	send_to_gpu(&vertices, &program)

	glfw.GetCurrentContext().Show()

	const (
		BG_RED   float32 = 0.03
		BG_BLUE  float32 = 0.03
		BG_GREEN float32 = 0.03
		BG_ALPHA float32 = 1.0
	)

	var (
		s_x float32 = 1.0
		s_y float32 = 1.0
		rot float32 = 0.0
		t_x float32 = 0.0
		t_y float32 = 0.0
	)

	width, height := window.GetFramebufferSize()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.ClearColor(BG_RED, BG_BLUE, BG_GREEN, BG_ALPHA)

		check_scale(window, &s_x, &s_y)
		check_rotation(window, &rot, "mouse", width, height)
		check_movement(window, &t_x, &t_y)

		mat_transformation := apply_transformation(s_x, s_y, rot, t_x, t_y)

		var loc uint8
		gl.GetUniformLocation(program, &loc)

		gl.UniformMatrix4fv(int32(loc), 1, true, &mat_transformation[0])

		gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 3)
		//gl.DrawArrays(gl.TRIANGLE_STRIP, 4, 4)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
