package colony

import ()

type Colony struct {
	Nodes []*ChannelNode
}

func (self *Colony) AddNode(node *ChannelNode) {
	for _, otherNode := range self.Nodes {
		node.AddEdge(NewChannelEdge(otherNode, otherNode.ResourceChan))
		otherNode.AddEdge(NewChannelEdge(node, node.ResourceChan))
	}

	self.Nodes = append(self.Nodes, node)
}