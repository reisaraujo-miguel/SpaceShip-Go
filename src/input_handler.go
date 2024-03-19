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
