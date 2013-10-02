package model

import (
	"github.com/aaasen/colony/colony"
	"github.com/aaasen/colony/view"
)

type Model struct {
	renderers chan<- view.Renderer
	control   <-chan bool

	colony *colony.Colony
}

func NewModel(renderers chan<- view.Renderer, control <-chan bool) *Model {
	return &Model{
		renderers: renderers,
		control:   control,
	}
}

func (self *Model) Listen() {
	for {
		select {
		case <-self.control:
			self.renderers <- view.NewMultiRenderer([]view.Renderer{
				view.NewSquareRenderer(0, 0, 100, 100),
				view.NewSquareRenderer(100, 100, 100, 100),
			})
		}
	}
}
