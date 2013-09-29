package graph

import (
	"testing"
)

func TestAdjacent(t *testing.T) {
	graph := makeLabeledTestGraph()

	edge := graph.GetNode("node1").Adjacent(graph.GetNode("node2"))

	if edge == nil {
		t.Fail()
	}
}

func TestAdjacentInverse(t *testing.T) {
	graph := makeLabeledTestGraph()

	edge := graph.GetNode("node2").Adjacent(graph.GetNode("node1"))

	if edge != nil {
		t.Fail()
	}
}

func TestNeighbors(t *testing.T) {
	graph := makeLabeledTestGraph()

	neighbors := graph.GetNode("node1").Neighbors()

	if len(neighbors) != 2 {
		t.Fail()
	}
}

func makeLabeledTestGraph() *LabeledGraph {
	graph := NewLabeledGraph()

	node1 := NewNode()
	node2 := NewNode()
	node3 := NewNode()

	node1.AddEdge(node2)
	node1.AddEdge(node3)
	node2.AddEdge(node3)
	node3.AddEdge(node1)

	graph.AddNode("node1", node1)
	graph.AddNode("node2", node2)
	graph.AddNode("node3", node3)

	return graph
}
