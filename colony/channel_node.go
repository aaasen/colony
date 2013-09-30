package colony

import (
	"log"
	"math/rand"
	"time"

	"github.com/aaasen/colony/graph"
)

type ChannelNode struct {
	graph.Node

	Resources    *Resource
	ResourceChan <-chan *Resource
	Ticker       <-chan time.Time

	id       int
	burnRate float64
}

func NewChannelNode(resources <-chan *Resource, ticker <-chan time.Time) *ChannelNode {
	return &ChannelNode{
		graph.Node{
			Edges: make([]graph.Edger, 0),
		},
		NewResource(0.0),
		resources,
		ticker,
		rand.Int(),
		1.0,
	}
}

func (self *ChannelNode) Listen() {
	for {
		select {
		case <-self.Ticker:
			log.Println("tick")
			log.Println(self.Resources.Amount)
			self.Resources = self.Resources.Subtract(self.burnRate)
			select {
			case resource := <-self.ResourceChan:
				log.Printf("get: %v", resource)

				self.Resources.Amount += resource.Amount
				toShare := NewResource(self.Resources.Amount - self.burnRate)
				self.Resources.Amount = self.burnRate

				// self.Resources, resource = AddResources(self.Resources, resource)
				// tempResource, share := SubtractResources(self.Resources, self.Resources.Subtract(self.burnRate))
				// self.Resources = tempResource

				log.Printf("giving: %v", toShare.Amount)

				go func(self *ChannelNode, resource *Resource) {
					self.distributeResource(resource, self.Edges)
				}(self, toShare)
			default:
				break
			}
		}
	}
}

func (self *ChannelNode) distributeResource(resource *Resource, edges []graph.Edger) {
	// based on the assumption that all edges are channel edges
	numEdges := len(self.Edges)

	for _, edge := range edges {
		if channelEdge, ok := edge.(*ChannelEdge); ok {
			channelEdge.Resources <- NewResource(resource.Amount / float64(numEdges))
		}
	}
}
