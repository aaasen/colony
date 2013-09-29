package graph

type Grapher interface {

	// returns the edge from x to y, or nil if there is none
	Adjacent(x, y *Noder) *Edger

	// returns all nodes the x connects to
	Neighbors(x *Noder) []*Edger

	// adds an edge between x and y if one does not already exist
	Add(x, y *Noder)

	// deletes the edge connecting x to y
	Delete(x, y *Noder)
}
