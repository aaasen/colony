package view

import (
	"log"
	"time"

	"github.com/aaasen/colony/controller"

	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
)

const (
	Title  = "Colony"
	Width  = 640
	Height = 480

	Tick = time.Second / 60.0
)

type View struct {
	renderers <-chan Renderer
	control   chan<- bool

	ticker <-chan time.Time
}

func NewView(renderers <-chan Renderer, control chan<- bool) *View {
	return &View{
		renderers: renderers,
		control:   control,
		ticker:    time.Tick(Tick),
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
		select {
		case <-self.ticker:
			self.control <- true

			select {
			case renderer := <-self.renderers:
				drawScene()

				renderer.Render()

				glfw.SwapBuffers()
			}
		}
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

	gl.ClearColor(0.180, 0.8, 0.443, 0.0)

	gl.Viewport(0, 0, Width, Height)
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, Width, Height, 0, 0, 1)
	gl.Color4f(1, 1, 1, 1)
}
