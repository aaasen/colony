package model

import (
	"log"
	"math/rand"

	"github.com/aaasen/colony/colony"
	structs "github.com/aaasen/colony/model/structs"
	"github.com/aaasen/colony/view"
)

type Model struct {
	renderers    chan<- view.Renderer
	keyEvents    <-chan int
	requestModel <-chan bool

	city *colony.Colony
}

func NewModel(renderers chan<- view.Renderer, keyEvents <-chan int, requestModel <-chan bool) *Model {
	c := colony.NewColony()
	c.AddNode(colony.NewFactoryNode(structs.NewVector2(rand.Float32()*100, rand.Float32()*100),
		1.0,
		make(chan *colony.Resource)))
	c.AddNode(colony.NewChannelNode(structs.NewVector2(rand.Float32()*100, rand.Float32()*100),
		1.0,
		make(chan *colony.Resource)))
	c.AddNode(colony.NewChannelNode(structs.NewVector2(rand.Float32()*100, rand.Float32()*100),
		1.0,
		make(chan *colony.Resource)))

	return &Model{
		renderers:    renderers,
		keyEvents:    keyEvents,
		requestModel: requestModel,
		city:         c,
	}
}

func (self *Model) Listen() {
	go self.city.Listen()

	go self.controllerListen()
	go self.viewListen()
}

func (self *Model) viewListen() {
	for {
		select {
		case <-self.requestModel:
			allRenderers := make([]view.Renderer, 0)

			for _, node := range self.city.Nodes {
				allRenderers = append(allRenderers, view.NewSquareRenderer(node.GetPosition().X, node.GetPosition().Y, 10, 10))

				for _, edge := range node.GetEdges() {
					allRenderers = append(allRenderers, view.NewLineRenderer(0, 0, 100, 100))

					log.Println(edge)
				}
			}

			self.renderers <- view.NewMultiRenderer(allRenderers)
		}
	}
}

func (self *Model) controllerListen() {
	for {
		select {
		case event := <-self.keyEvents:
			self.city.AddNode(
				colony.NewChannelNode(structs.NewVector2(rand.Float32()*100, rand.Float32()*100),
					1.0,
					make(chan *colony.Resource)))

			log.Println(event)
		}
	}
}
