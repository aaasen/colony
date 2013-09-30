package colony

import (
	"log"

	"github.com/aaasen/colony/graph"
)

type ChannelNode struct {
	graph.Node

	Resources <-chan *Resource
}

func NewChannelNode(resources <-chan *Resource) *ChannelNode {
	return &ChannelNode{
		graph.Node{
			Edges: make([]graph.Edger, 0),
		},
		resources,
	}
}

func (self *ChannelNode) Listen() {
	for {
		select {
		case resource := <-self.Resources:
			log.Println(resource)

			numEdges := len(self.Edges)

			for _, edge := range self.Edges {
				if channelEdge, ok := edge.(*ChannelEdge); ok {
					channelEdge.Resources <- NewResource(resource.Amount / float64(numEdges))
				}
			}
		}
	}
}
