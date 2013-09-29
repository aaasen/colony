package graph

type Grapher interface {

	// returns the edge from x to y, or nil if there is none
	Adjacent(x, y *Node) *Edge

	// returns all nodes the x connects to
	Neighbors(x *Node) []*Edge

	// adds an edge between x and y if one does not already exist
	AddEdge(x, y *Node)

	// deletes the edge connecting x to y
	DeleteEdge(x, y *Node)
}
