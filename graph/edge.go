package graph

type Edge struct {
	Target *Node
}

func NewEdge(target *Node) *Edge {
	return &Edge{
		Target: target,
	}
}

func (self *Edge) GetTarget() *Node {
	return self.Target
}
