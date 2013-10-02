package view

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

type View struct {
	Renderers <-chan Renderer
}

func NewView(renderChan <-chan Renderer) *View {
	return &View{
		renderChan,
	}
}

func (self *View) Init() {
	initGLFW()
	glfw.SetKeyCallback(controller.KeyHandler)
	glfw.SetMouseButtonCallback(controller.MouseButtonHandler)
	glfw.SetMousePosCallback(controller.MousePosHandler)

	initScene()

}

func (self *View) Listen() {
	for glfw.WindowParam(glfw.Opened) == 1 {
		// select {
		// case renderer := <-self.Renderers:
		// 	renderer.Render()
		// }
		drawScene()
		glfw.SwapBuffers()
	}
}

func (self *View) Destroy() {
	terminateGLFW()
	destroyScene()
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
