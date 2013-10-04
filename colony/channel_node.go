package colony

import (
	"log"
	"time"

	"github.com/aaasen/colony/graph"
	structs "github.com/aaasen/colony/model/structs"
)

const (
	TickDuration = time.Second
)

type ChannelNode struct {
	graph.Node

	Position *structs.Vector2

	BurnRate float64

	Resources    *Resource
	ResourceChan chan *Resource

	Ticker <-chan time.Time
}

func NewChannelNode(position *structs.Vector2, burnRate float64, resources chan *Resource) *ChannelNode {
	return &ChannelNode{
		graph.Node{
			Edges: make([]graph.Edger, 0),
		},
		position,
		burnRate,
		NewResource(0.0),
		resources,
		time.Tick(TickDuration),
	}
}

func (self *ChannelNode) Listen() {
	for {
		select {
		case <-self.Ticker:
			self.Resources = self.Resources.Subtract(self.BurnRate)
			select {
			case resource := <-self.ResourceChan:
				log.Printf("get: %v", resource.Amount)

				self.Resources = self.Resources.Add(resource.Amount)
				newResource, toShare := SubtractResources(self.Resources, NewResource(self.Resources.Amount-self.BurnRate))
				self.Resources = newResource

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

func (self *ChannelNode) GetResourceChan() chan *Resource {
	return self.ResourceChan
}

func (self *ChannelNode) GetPosition() *structs.Vector2 {
	return self.Position
}

func (self *ChannelNode) GetEdges() []graph.Edger {
	return self.Edges
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
