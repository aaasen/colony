package colony

import (
	"github.com/aaasen/colony/graph"
)

type ChannelNode struct {
	graph.Node

	Resources []<-chan Resource
}

func NewChannelNode(resources []<-chan Resource) *ChannelNode {
	return &ChannelNode{
		graph.Node{
			Edges: make([]graph.Edger, 0),
		},
		resources,
	}
}
