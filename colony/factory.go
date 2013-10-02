package colony

import (
	"log"

	structs "github.com/aaasen/colony/model/structs"
)

type FactoryNode struct {
	ChannelNode

	CreateRate float64
}

func NewFactoryNode(position *structs.Vector2, createRate float64, resources <-chan *Resource) *FactoryNode {
	return &FactoryNode{
		*NewChannelNode(position, 0.0, resources),
		createRate,
	}
}

func (self *FactoryNode) Listen() {
	for {
		select {
		case <-self.Ticker:
			self.Resources = self.Resources.Subtract(self.BurnRate).Add(self.CreateRate)
			select {
			case resource := <-self.ResourceChan:
				log.Printf("get: %v", resource.Amount)

				self.Resources = self.Resources.Add(resource.Amount)
				newResource, toShare := SubtractResources(self.Resources, NewResource(self.Resources.Amount-self.BurnRate))
				self.Resources = newResource

				log.Printf("giving: %v", toShare.Amount)

				go func(node *ChannelNode, resource *Resource) {
					node.distributeResource(resource, self.Edges)
				}(&self.ChannelNode, toShare)
			default:
				break
			}
		}
	}
}
