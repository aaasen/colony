package graph

type Noder interface {

	// adds an edge to the node if one does not already exist
	AddEdge(Edger)

	// deletes the edge connecting the node to a target
	DeleteEdge(Noder)

	// returns the edge from the node to a target, or nil if there is none
	Adjacent(Noder) Edger

	// returns all nodes that the node connects to
	Neighbors() []Noder
}
