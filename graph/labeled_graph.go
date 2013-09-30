package graph

type LabeledGraph struct {
	nodes map[string]Noder
}

func NewLabeledGraph() *LabeledGraph {
	return &LabeledGraph{
		nodes: make(map[string]Noder, 0),
	}
}

func (self *LabeledGraph) AddNode(label string, node Noder) {
	self.nodes[label] = node
}

func (self *LabeledGraph) GetNode(label string) Noder {
	return self.nodes[label]
}
