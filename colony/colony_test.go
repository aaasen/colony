package colony

import (
	"testing"

	"github.com/aaasen/colony/graph"
)

func TestColony(t *testing.T) {
	colony := graph.NewLabeledGraph()

	node1 := graph.NewNode()
	node2 := graph.NewNode()
	node3 := graph.NewNode()

	node1.AddEdge(NewChannelEdge(node2))
	node1.AddEdge(NewChannelEdge(node3))
	node2.AddEdge(NewChannelEdge(node3))
	node3.AddEdge(NewChannelEdge(node1))

	colony.AddNode("node1", node1)
	colony.AddNode("node2", node2)
	colony.AddNode("node3", node3)
}
