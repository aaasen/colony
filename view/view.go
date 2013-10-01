package main

import (
	"log"

	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
)

const (
	Title  = "Colony"
	Width  = 640
	Height = 480
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("glfw: %s\n", err)
		return
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		log.Fatalf("glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	if err := gl.Init(); err != nil {
		log.Fatalf("gl: %s\n", err)
	}

	if err := initScene(); err != nil {
		log.Fatalf("init: %s\n", err)
		return
	}
	defer destroyScene()

	for glfw.WindowParam(glfw.Opened) == 1 {

		drawScene()
		glfw.SwapBuffers()
	}
}

func initScene() (err error) {
	gl.Disable(gl.DEPTH_TEST)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)

	gl.Viewport(0, 0, Width, Height)

	return
}

func destroyScene() {
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, Width, Height, 0, 0, 1)
	gl.Color4f(1, 1, 1, 1)

	gl.Begin(gl.LINES)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(100, 100)
	gl.End()
}
