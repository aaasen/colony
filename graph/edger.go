package graph

type Edger interface {

	// gets the node referenced by the edge
	GetNode() *Noder
}
