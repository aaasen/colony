package colony

import (
	"log"
	"time"
)

type FactoryNode struct {
	ChannelNode

	CreateRate float64
}

func NewFactoryNode(resources <-chan *Resource, ticker <-chan time.Time, createRate float64) *FactoryNode {
	return &FactoryNode{
		*NewChannelNode(resources, ticker),
		createRate,
	}
}

func (self *FactoryNode) Listen() {
	for {
		select {
		case <-self.Ticker:
			self.Resources = self.Resources.Subtract(self.burnRate).Add(self.CreateRate)
			select {
			case resource := <-self.ResourceChan:
				log.Printf("get: %v", resource.Amount)

				self.Resources = self.Resources.Add(resource.Amount)
				newResource, toShare := SubtractResources(self.Resources, NewResource(self.Resources.Amount-self.burnRate))
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
