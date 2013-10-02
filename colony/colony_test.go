package colony

import (
	"testing"
	"time"

	"github.com/aaasen/colony/graph"
	structs "github.com/aaasen/colony/model/structs"
)

func TestColony(t *testing.T) {
	colony := graph.NewLabeledGraph()

	to1 := make(chan *Resource, 3)
	to2 := make(chan *Resource, 3)
	to3 := make(chan *Resource, 3)

	node1 := NewFactoryNode(&structs.Vector2{0.0, 0.0}, 10.0, to1)
	node2 := NewChannelNode(&structs.Vector2{0.0, 0.0}, 1.0, to2)
	node3 := NewChannelNode(&structs.Vector2{0.0, 0.0}, 1.0, to3)

	node1.AddEdge(NewChannelEdge(node2, to2))
	node1.AddEdge(NewChannelEdge(node3, to3))
	node2.AddEdge(NewChannelEdge(node1, to1))
	node2.AddEdge(NewChannelEdge(node3, to3))
	node3.AddEdge(NewChannelEdge(node1, to1))
	node3.AddEdge(NewChannelEdge(node2, to2))

	colony.AddNode("node1", node1)
	colony.AddNode("node2", node2)
	colony.AddNode("node3", node3)

	go node1.Listen()
	go node2.Listen()
	go node3.Listen()

	if edge, ok := node1.Edges[0].(*ChannelEdge); ok {
		edge.Resources <- NewResource(10.0)
	}

	time.Sleep(time.Second * 10)
}
