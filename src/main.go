package main

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize glfw: ", err)
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(600, 600, "SpaceShip-Go!", nil, nil)

	if err != nil {
		log.Fatalln("Could not create window: ", err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("Could not init gl: ", err)
	}

	vertex, fragment := getShaders()

	compileShader(&vertex)
	compileShader(&fragment)

	program := gl.CreateProgram()

	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)

	linkProgram(&program)

	gl.UseProgram(program)

	var globalVertices []mgl32.Vec2

	shipVertices := []mgl32.Vec2{
		{+0.0, +0.05},
		{+0.05, -0.05},
		{-0.05, -0.05},
	}

	boxVertices := []mgl32.Vec2{
		{0.10, -0.05},
		{0.10, 0.05},
		{0.20, -0.05},
		{0.20, 0.05},
	}

	ship := createBody()
	ship.instantiate(shipVertices, &globalVertices)

	box := createBody()
	box.instantiate(boxVertices, &globalVertices)

	sendToGpu(&globalVertices, &program)

	window.Show()

	const (
		BG_RED   float32 = 0.03
		BG_BLUE  float32 = 0.03
		BG_GREEN float32 = 0.03
		BG_ALPHA float32 = 1.0
	)

	var (
		angle       float32 = 0.0
		sX, sY      float32 = 1.0, 1.0
		tX, tY      float32 = 0.0, 0.0
		speed       float32 = 0.0
		sizeFactor  float32 = 0.05
		speedFactor float32 = 0.05
		deltaTime   float32
		lastFrame   float32 = float32(glfw.GetTime())
	)

	for !window.ShouldClose() {
		currentFrame := float32(glfw.GetTime())
		deltaTime = currentFrame - lastFrame
		lastFrame = currentFrame

		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.ClearColor(BG_RED, BG_BLUE, BG_GREEN, BG_ALPHA)

		checkScale(window, &sX, &sY, sizeFactor)
		ship.scale(sX, sY)

		checkMovement(window, &speed, speedFactor)
		moveTowardsMouse(window, &tX, &tY, speed*deltaTime)
		screenWrap(&tX, &tY)
		ship.translate(tX, tY)

		checkRotation(window, tX, tY, &angle)
		ship.rotate(angle)

		ship.drawBody(&program)
		box.drawBody(&program)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
