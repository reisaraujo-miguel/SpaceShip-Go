package main

import (
	"fmt"
	"math"

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

func check_rotation(window *glfw.Window, rot *float32, rot_type string, width int, height int) {
	switch rot_type {
	case "key":
		if window.GetKey(glfw.KeyQ) == glfw.Press {
			*rot += 0.01
		}
		if window.GetKey(glfw.KeyE) == glfw.Press {
			*rot -= 0.01
		}
	case "mouse":
		if (width | height) <= 0 {
			fmt.Printf("check_rotation(): `width` and `height` must be greater than ZERO.\n")
			panic("You need to pass a valid value for `width` and `height` when using `rot_type: \"mouse\"`.")
		}

		cursor_x, cursor_y := window.GetCursorPos()

		cat_x := cursor_x - float64(width/2)
		cat_y := float64(height/2) - cursor_y

		if cat_y > 0 {
			*rot = float32(math.Atan(cat_x / cat_y))
		} else {
			if cat_x > 0 {
				*rot = float32(math.Atan(-cat_y/cat_x)) + 1.571
			} else {
				*rot = float32(math.Atan(-cat_y/cat_x)) - 1.571
			}
		}
	default:
		fmt.Printf("check_rotation(): `rot_type: \"%v\"` is not a valid type. Try `\"key\"` or `\"mouse\"`.\n", rot_type)
		panic("Invalid `rot_type` argument.")
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
