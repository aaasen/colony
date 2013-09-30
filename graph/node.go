package graph

type Node struct {
	Edges []*Edge
}

func NewNode() *Node {
	return &Node{
		Edges: make([]*Edge, 0),
	}
}

// adds an edge to the node if one does not already exist
func (self *Node) AddEdge(edge *Edge) {
	if self.Adjacent(edge.Target) == nil {
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
func (self *Node) Adjacent(target *Node) *Edge {
	for _, edge := range self.Edges {
		if edge.Target == target {
			return edge
		}
	}

	return nil
}

// returns all nodes that the node connects to
func (self *Node) Neighbors() []*Node {
	neighbors := make([]*Node, 0)

	for _, edge := range self.Edges {
		neighbors = append(neighbors, edge.Target)
	}

	return neighbors
}

// returns the index of the edge pointing to the target node or -1 if there is none
func (self *Node) findEdgeIndex(target *Node) int {
	for i, edge := range self.Edges {
		if edge.Target == target {
			return i
		}
	}

	return -1
}
