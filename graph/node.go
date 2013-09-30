package graph

type Node struct {
	Edges []Edger
}

func NewNode() *Node {
	return &Node{
		Edges: make([]Edger, 0),
	}
}

// adds an edge to the node if one does not already exist
func (self *Node) AddEdge(edge Edger) {
	if self.Adjacent(edge.GetTarget()) == nil {
		self.Edges = append(self.Edges, edge)
	}
}

// deletes the edge connecting the node to a target
func (self *Node) DeleteEdge(target *Node) {
	edgeIndex := self.findEdgeIndex(target)

	if edgeIndex >= 0 {
		self.Edges = append(self.Edges[:edgeIndex], self.Edges[edgeIndex+1:]...)
	}
}

// returns the edge from the node to a target, or nil if there is none
func (self *Node) Adjacent(target *Node) Edger {
	for _, edge := range self.Edges {
		if edge.GetTarget() == target {
			return edge
		}
	}

	return nil
}

// returns all nodes that the node connects to
func (self *Node) Neighbors() []*Node {
	neighbors := make([]*Node, 0)

	for _, edge := range self.Edges {
		neighbors = append(neighbors, edge.GetTarget())
	}

	return neighbors
}

// returns the index of the edge pointing to the target node or -1 if there is none
func (self *Node) findEdgeIndex(target *Node) int {
	for i, edge := range self.Edges {
		if edge.GetTarget() == target {
			return i
		}
	}

	return -1
}
