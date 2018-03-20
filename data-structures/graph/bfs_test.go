package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	g := NewGraph()
	v1 := g.AddVertexWith(1)
	v2 := g.AddVertexWith(2)
	v3 := g.AddVertexWith(3)
	v4 := g.AddVertexWith(4)
	v5 := g.AddVertexWith(5)

	g.AddUndirectedEdge(v1, v2, 0)
	g.AddUndirectedEdge(v1, v5, 0)
	g.AddUndirectedEdge(v2, v3, 0)
	g.AddUndirectedEdge(v2, v4, 0)
	g.AddUndirectedEdge(v2, v5, 0)
	g.AddUndirectedEdge(v3, v4, 0)
	g.AddUndirectedEdge(v4, v5, 0)

	if actual := BreadthFirstSearch(g, v1, v1); actual[0] != v1 {
		t.Errorf("Expected %v, instead got %v", v1, actual[0])
	}
	if actual := BreadthFirstSearch(g, v3, v3); actual[0] != v3 {
		t.Errorf("Expected %v, instead got %v", v3, actual[0])
	}

	if actual := BreadthFirstSearch(g, v1, v2); len(actual) != 2 {
		t.Errorf("Expected %v, instead got %v", 2, len(actual))
	}
	if actual := BreadthFirstSearch(g, v2, v1); len(actual) != 2 {
		t.Errorf("Expected %v, instead got %v", 2, len(actual))
	}
	if actual := BreadthFirstSearch(g, v1, v3); len(actual) != 3 {
		t.Errorf("Expected %v, instead got %v", 3, len(actual))
	}
	if actual := BreadthFirstSearch(g, v3, v1); len(actual) != 3 {
		t.Errorf("Expected %v, instead got %v", 3, len(actual))
	}
	if actual := BreadthFirstSearch(g, v3, v5); len(actual) != 3 {
		t.Errorf("Expected %v, instead got %v", 3, len(actual))
	}
}
