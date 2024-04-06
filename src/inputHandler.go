package main

import (
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func checkMovement(window *glfw.Window, speed *float32, factor float32) {
	if window.GetKey(glfw.KeyUp) == glfw.Press {
		*speed += factor
	}
	if window.GetKey(glfw.KeyDown) == glfw.Press {
		*speed -= factor
	}
}

func checkScale(window *glfw.Window, sX *float32, sY *float32, factor float32) {
	minusKey := window.GetKey(glfw.KeyMinus) == glfw.Press
	shiftKey := (window.GetKey(glfw.KeyLeftShift) == glfw.Press) || (window.GetKey(glfw.KeyRightShift) == glfw.Press)
	plusKey := (window.GetKey(glfw.KeyEqual) == glfw.Press) && shiftKey

	if plusKey {
		*sY += factor
		*sX += factor
	} else if minusKey {
		if (*sX > 0) && (*sY > 0) {
			*sY -= factor
			*sX -= factor
		}
	}
}

func checkRotation(window *glfw.Window, x float32, y float32, angle *float32) {
	width, height := window.GetFramebufferSize()

	cursorX, cursorY := window.GetCursorPos()
	originX, originY := (width / 2), (height / 2)

	catX := (cursorX - float64(originX)) - float64(x*float32(originX))
	catY := (float64(originY) - cursorY) - float64(y*float32(originY))

	if catY > 0 {
		*angle = float32(math.Atan(catX / catY))
	} else {
		if catX > 0 {
			*angle = float32(math.Atan((-catY)/catX)) + 1.571
		} else {
			*angle = float32(math.Atan((-catY)/catX)) - 1.571
		}
	}
}
