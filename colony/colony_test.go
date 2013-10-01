package colony

import (
	"testing"
	"time"

	"github.com/aaasen/colony/graph"
)

func TestColony(t *testing.T) {
	tickDuration := time.Second

	colony := graph.NewLabeledGraph()

	to1 := make(chan *Resource, 3)
	to2 := make(chan *Resource, 3)
	to3 := make(chan *Resource, 3)

	node1 := NewFactoryNode(to1, time.Tick(tickDuration), 10.0)
	node2 := NewChannelNode(to2, time.Tick(tickDuration))
	node3 := NewChannelNode(to3, time.Tick(tickDuration))

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
