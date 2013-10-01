package main

import (
	"log"

	"github.com/aaasen/colony/controller"

	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
)

const (
	Title  = "Colony"
	Width  = 640
	Height = 480
)

func main() {
	initGLFW()
	glfw.SetKeyCallback(controller.KeyHandler)
	defer terminateGLFW()

	initScene()
	defer destroyScene()

	for glfw.WindowParam(glfw.Opened) == 1 {
		drawScene()
		glfw.SwapBuffers()
	}
}

func destroyScene() {

}

func initScene() {
	if err := gl.Init(); err != nil {
		log.Fatalf("gl: %s\n", err)
	}

	gl.Disable(gl.DEPTH_TEST)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)

	gl.Viewport(0, 0, Width, Height)
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
