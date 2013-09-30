package colony

import (
	"log"
	"math/rand"

	"github.com/aaasen/colony/graph"
)

type ChannelNode struct {
	graph.Node

	id        int
	Resources <-chan *Resource
}

func NewChannelNode(resources <-chan *Resource) *ChannelNode {
	return &ChannelNode{
		graph.Node{
			Edges: make([]graph.Edger, 0),
		},
		rand.Int(),
		resources,
	}
}

func (self *ChannelNode) Listen() {
	for {
		select {
		case resource := <-self.Resources:
			go func(self *ChannelNode, resource *Resource) {
				log.Printf("%v: %v", self.id, resource)

				self.distributeResource(resource, self.Edges)
			}(self, resource)
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
