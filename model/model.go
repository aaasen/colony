package model

import (
	"math/rand"

	"github.com/aaasen/colony/colony"
	structs "github.com/aaasen/colony/model/structs"
	"github.com/aaasen/colony/view"
)

type Model struct {
	renderers chan<- view.Renderer
	control   <-chan bool

	city *colony.Colony
}

func NewModel(renderers chan<- view.Renderer, control <-chan bool) *Model {
	c := colony.NewColony()
	c.AddNode(colony.NewChannelNode(structs.NewVector2(rand.Float32()*100, rand.Float32()*100),
		1.0,
		make(chan *colony.Resource)))
	c.AddNode(colony.NewChannelNode(structs.NewVector2(rand.Float32()*100, rand.Float32()*100),
		1.0,
		make(chan *colony.Resource)))
	c.AddNode(colony.NewChannelNode(structs.NewVector2(rand.Float32()*100, rand.Float32()*100),
		1.0,
		make(chan *colony.Resource)))

	return &Model{
		renderers: renderers,
		control:   control,
		city:      c,
	}
}

func (self *Model) Listen() {
	for {
		select {
		case <-self.control:
			allRenderers := make([]view.Renderer, len(self.city.Nodes))

			for i, node := range self.city.Nodes {
				allRenderers[i] = view.NewSquareRenderer(node.Position.X, node.Position.Y, 10, 10)
			}

			self.renderers <- view.NewMultiRenderer(allRenderers)
		}
	}
}
