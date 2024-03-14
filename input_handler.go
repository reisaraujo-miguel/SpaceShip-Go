package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func key_event(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
}

func check_movement(window *glfw.Window, x *float32, y *float32) {
	if window.GetKey(glfw.KeyUp) == glfw.Press {
		*y += 0.01
	}
	if window.GetKey(glfw.KeyDown) == glfw.Press {
		*y -= 0.01
	}
	if window.GetKey(glfw.KeyRight) == glfw.Press {
		*x += 0.01
	}
	if window.GetKey(glfw.KeyLeft) == glfw.Press {
		*x -= 0.01
	}
}

func check_rotation(window *glfw.Window, rot *float32) {
	if window.GetKey(glfw.KeyQ) == glfw.Press {
		*rot += 0.01
	}
	if window.GetKey(glfw.KeyE) == glfw.Press {
		*rot -= 0.01
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
