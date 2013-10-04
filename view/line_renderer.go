package view

import (
	gl "github.com/chsc/gogl/gl21"
)

type LineRenderer struct {
	x1, y1, x2, y2 float32
}

func NewLineRenderer(x1, y1, x2, y2 float32) *LineRenderer {
	return &LineRenderer{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

func (self *LineRenderer) Render() {
	gl.Begin(gl.LINES)

	gl.Vertex2f(gl.Float(self.x1), gl.Float(self.y1))
	gl.Vertex2f(gl.Float(self.x2), gl.Float(self.y2))

	gl.End()
}
