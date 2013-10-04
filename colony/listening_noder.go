package colony

import (
	"github.com/aaasen/colony/graph"
	structs "github.com/aaasen/colony/model/structs"
)

type ListeningNoder interface {
	Listen()

	GetResourceChan() chan *Resource
	GetPosition() *structs.Vector2
	GetEdges() []graph.Edger

	// adds an edge to the node if one does not already exist
	AddEdge(graph.Edger)

	// deletes the edge connecting the node to a target
	DeleteEdge(graph.Noder)

	// returns the edge from the node to a target, or nil if there is none
	Adjacent(graph.Noder) graph.Edger

	// returns all nodes that the node connects to
	Neighbors() []graph.Noder
}
