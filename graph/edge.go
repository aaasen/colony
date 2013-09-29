package graph

type Edge struct {
	Target *Node
}

func NewEdge(target *Node) *Edge {
	return &Edge{
		Target: target,
	}
}
