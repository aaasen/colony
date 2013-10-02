package model

import (
	"github.com/aaasen/colony/colony"
	"github.com/aaasen/colony/view"
)

type Model struct {
	Renderers chan<- view.Renderer
	Control   <-chan bool

	colony *colony.Colony
}

func NewModel(renderers chan<- view.Renderer, control <-chan bool) *Model {
	return &Model{
		Renderers: renderers,
		Control:   control,
	}
}

func (self *Model) Listen() {
	for {
		select {
		case <-self.Control:

		}
	}
}
