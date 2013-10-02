package view

import (
	gl "github.com/chsc/gogl/gl21"
)

type SquareRenderer struct {
	x, y, w, h float32
}

func NewSquareRenderer(x, y, w, h float32) *SquareRenderer {
	return &SquareRenderer{
		x: x,
		y: y,
		w: w,
		h: h,
	}
}

func (self *SquareRenderer) Render() {
	gl.Begin(gl.QUADS)

	gl.Vertex2f(gl.Float(self.x), gl.Float(self.y))
	gl.Vertex2f(gl.Float(self.x+self.w), gl.Float(self.y))
	gl.Vertex2f(gl.Float(self.x+self.w), gl.Float(self.y+self.h))
	gl.Vertex2f(gl.Float(self.x), gl.Float(self.y+self.h))

	gl.End()
}
