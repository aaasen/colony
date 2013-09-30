package graph

type Edge struct {
	Target Noder
}

func NewEdge(target Noder) *Edge {
	return &Edge{
		Target: target,
	}
}

func (self *Edge) GetTarget() Noder {
	return self.Target
}
