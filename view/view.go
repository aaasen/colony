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
	initGLFW()
	defer terminateGLFW()

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

func destroyScene() {

}

func initScene() (err error) {
	gl.Disable(gl.DEPTH_TEST)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)

	gl.Viewport(0, 0, Width, Height)

	return
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
