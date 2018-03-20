package graph

import (
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := NewGraph()

	g.AddVertexWith(1)
	g.AddVertexWith(2)

	expected := 2
	if actual := g.Count(); actual != expected {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	v1 := NewVertex(1)
	g.AddVertex(v1)    // Does nothing.
	g.AddVertexWith(1) // Does nothing.

	expected = 2
	if actual := g.Count(); actual != expected {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	if !(g.Contains(v1)) {
		t.Errorf("Expected %v, instead got %v", true, false)
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()

	v1 := NewVertex(1)
	v2 := NewVertex(2)
	v3 := NewVertex(3)
	v4 := NewVertex(4)

	// Because vertices 'v1' and 'v2' are not in the graph,
	// below statements does nothing.
	g.AddDirectedEdge(v1, v2, 12)   // Does nothing.
	g.AddUndirectedEdge(v1, v2, 12) // Does nothing.

	if actual := len(g.Edges(v1)); actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}
	if actual := len(g.Edges(v2)); actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}

	// Because vertex 'v2' is not in the graph,
	// below function calls 'AddEdge' does nothing.
	g.AddVertex(v1)
	g.AddDirectedEdge(v1, v2, 12)   // Does nothing.
	g.AddUndirectedEdge(v1, v2, 12) // Does nothing.

	if actual := len(g.Edges(v1)); actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}
	if actual := len(g.Edges(v2)); actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}

	g.AddVertex(v2)
	g.AddUndirectedEdge(v1, v2, 12)

	if actual := len(g.Edges(v1)); actual != 1 {
		t.Errorf("Expected %v, instead got %v", 1, actual)
	}
	if actual := len(g.Edges(v2)); actual != 1 {
		t.Errorf("Expected %v, instead got %v", 1, actual)
	}

	g.AddVertex(v3)
	g.AddVertex(v4)
	g.AddDirectedEdge(v3, v4, 34)

	if actual := len(g.Edges(v3)); actual != 1 {
		t.Errorf("Expected %v, instead got %v", 1, actual)
	}
	if actual := len(g.Edges(v4)); actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}

	// Because edges (v1, v2), and (v3, v4) already exist,
	// below statements does nothing.
	g.AddUndirectedEdge(v1, v2, 1212) // Does nothing.
	g.AddDirectedEdge(v3, v4, 3434)   // Does nothing.

	if actual := g.Weight(v1, v2); actual != 12 {
		t.Errorf("Expected %v, instead got %v", 12, actual)
	}
	if actual := g.Weight(v3, v4); actual != 34 {
		t.Errorf("Expected %v, instead got %v", 34, actual)
	}
	if actual := len(g.Edges(v1)); actual != 1 {
		t.Errorf("Expected %v, instead got %v", 1, actual)
	}
	if actual := len(g.Edges(v2)); actual != 1 {
		t.Errorf("Expected %v, instead got %v", 1, actual)
	}
	if actual := len(g.Edges(v3)); actual != 1 {
		t.Errorf("Expected %v, instead got %v", 1, actual)
	}
	if actual := len(g.Edges(v4)); actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}
}
