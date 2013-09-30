package colony

import (
	"github.com/aaasen/colony/graph"
)

type ChannelEdge struct {
	graph.Edge

	Resources chan Resource
}

func NewChannelEdge(target *graph.Node) *ChannelEdge {
	return &ChannelEdge{
		graph.Edge{
			Target: target,
		},
		nil,
	}
}

func (self *ChannelEdge) GetTarget() graph.Noder {
	return self.Target
}
