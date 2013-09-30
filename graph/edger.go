package graph

type Edger interface {

	// returns the target node of the edge
	GetTarget() Noder
}
