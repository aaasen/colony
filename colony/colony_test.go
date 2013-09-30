package colony

import (
	"testing"

	"github.com/aaasen/colony/graph"
)

func TestColony(t *testing.T) {
	colony := graph.NewLabeledGraph()

	from1to2 := make(chan Resource)
	from1to3 := make(chan Resource)
	from2to3 := make(chan Resource)
	from3to1 := make(chan Resource)

	node1 := NewChannelNode([]<-chan Resource{from3to1})
	node2 := NewChannelNode([]<-chan Resource{from1to2})
	node3 := NewChannelNode([]<-chan Resource{from1to3, from2to3})

	node1.AddEdge(NewChannelEdge(node2, from1to2))
	node1.AddEdge(NewChannelEdge(node3, from1to3))
	node2.AddEdge(NewChannelEdge(node3, from2to3))
	node3.AddEdge(NewChannelEdge(node1, from3to1))

	colony.AddNode("node1", node1)
	colony.AddNode("node2", node2)
	colony.AddNode("node3", node3)
}
