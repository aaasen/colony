package graph

type LabeledGraph struct {
	nodes map[string]*Node
}

func NewLabeledGraph() *LabeledGraph {
	return &LabeledGraph{
		nodes: make(map[string]*Node, 0),
	}
}

func (self *LabeledGraph) AddNode(label string, node *Node) {
	self.nodes[label] = node
}

func (self *LabeledGraph) GetNode(label string) *Node {
	return self.nodes[label]
}
