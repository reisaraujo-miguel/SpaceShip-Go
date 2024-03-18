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

	var global_vertices []mgl32.Vec2

	ship_vertices := []mgl32.Vec2{
		{+0.0, +0.05},
		{+0.05, -0.05},
		{-0.05, -0.05},
	}

	box_vertices := []mgl32.Vec2{
		{0.10, -0.05},
		{0.10, 0.05},
		{0.20, -0.05},
		{0.20, 0.05},
	}

	ship := create_body()
	ship.instantiate(ship_vertices, &global_vertices)

	box := create_body()
	box.instantiate(box_vertices, &global_vertices)

	send_to_gpu(&global_vertices, &program)

	glfw.GetCurrentContext().Show()

	const (
		BG_RED   float32 = 0.03
		BG_BLUE  float32 = 0.03
		BG_GREEN float32 = 0.03
		BG_ALPHA float32 = 1.0
	)

	var (
		angle    float32 = 0.0
		s_x, s_y float32 = 1.0, 1.0
		t_x, t_y float32 = 0.0, 0.0
	)

	for !window.ShouldClose() {
		glfw.PollEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.ClearColor(BG_RED, BG_BLUE, BG_GREEN, BG_ALPHA)

		check_scale(window, &s_x, &s_y)
		ship.scale(s_x, s_y)

		check_movement(window, &t_x, &t_y)
		ship.translate(t_x, t_y)

		check_rotation(window, &angle, t_x, t_y)
		ship.rotate(angle)

		ship.draw_body(&program)
		box.draw_body(&program)

		window.SwapBuffers()
	}
}
