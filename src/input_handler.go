package main

import (
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func key_event(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
}

func check_movement(window *glfw.Window, inc *float32) {
	if window.GetKey(glfw.KeyUp) == glfw.Press {
		*inc += 0.0005
	}
	if window.GetKey(glfw.KeyDown) == glfw.Press {
		*inc -= 0.0005
	}
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

func check_rotation(window *glfw.Window, angle *float32, x float32, y float32) {
	width, height := window.GetFramebufferSize()

	cursor_x, cursor_y := window.GetCursorPos()
	origin_x, origin_y := (width / 2), (height / 2)

	cat_x := (cursor_x - float64(origin_x)) - float64(x*float32(origin_x))
	cat_y := (float64(origin_y) - cursor_y) - float64(y*float32(origin_y))

	if cat_y > 0 {
		*angle = float32(math.Atan(cat_x / cat_y))
	} else {
		if cat_x > 0 {
			*angle = float32(math.Atan((-cat_y)/cat_x)) + 1.571
		} else {
			*angle = float32(math.Atan((-cat_y)/cat_x)) - 1.571
		}
	}
}

func check_scale(window *glfw.Window, s_x *float32, s_y *float32) {
	if window.GetKey(glfw.KeyW) == glfw.Press {
		*s_y += 0.01
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		*s_y -= 0.01
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		*s_x += 0.01
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		*s_x -= 0.01
	}
}

func mouse_event(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	print("mouse")
}
