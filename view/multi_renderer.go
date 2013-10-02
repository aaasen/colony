package view

type MultiRenderer struct {
	Renderers []Renderer
}

func NewMultiRenderer(renderers []Renderer) *MultiRenderer {
	return &MultiRenderer{
		Renderers: renderers,
	}
}

func (self *MultiRenderer) Render() {
	for _, renderer := range self.Renderers {
		renderer.Render()
	}
}
