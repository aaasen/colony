package colony

import ()

type Colony struct {
	Nodes []ListeningNoder
}

func NewColony() *Colony {
	return &Colony{
		Nodes: make([]ListeningNoder, 0),
	}
}

func (self *Colony) AddNode(node ListeningNoder) {
	for _, otherNode := range self.Nodes {
		node.AddEdge(NewChannelEdge(otherNode, otherNode.GetResourceChan()))
		otherNode.AddEdge(NewChannelEdge(node, node.GetResourceChan()))
	}

	self.Nodes = append(self.Nodes, node)
}

func (self *Colony) Listen() {
	for _, node := range self.Nodes {
		go node.Listen()
	}
}
