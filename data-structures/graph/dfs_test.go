package graph

import (
	"testing"
)

func TestDFS(t *testing.T) {
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

	if actual := DepthFirstSearch(g, v1, v1); actual[0] != v1 {
		t.Errorf("Expected %v, instead got %v", v1, actual[0])
	}
	if actual := DepthFirstSearch(g, v3, v3); actual[0] != v3 {
		t.Errorf("Expected %v, instead got %v", v3, actual[0])
	}

	if actual := DepthFirstSearch(g, v1, v3); actual[0] != v1 && actual[len(actual)-1] != v3 {
		t.Errorf("Expected first and last element (%v, %v), instead got (%v, %v)", v1, v3, actual[0], actual[len(actual)-1])
	}
	if actual := DepthFirstSearch(g, v3, v1); actual[0] != v3 && actual[len(actual)-1] != v1 {
		t.Errorf("Expected first and last element (%v, %v), instead got (%v, %v)", v3, v1, actual[0], actual[len(actual)-1])
	}
	if actual := DepthFirstSearch(g, v1, v4); actual[0] != v1 && actual[len(actual)-1] != v4 {
		t.Errorf("Expected first and last element (%v, %v), instead got (%v, %v)", v1, v4, actual[0], actual[len(actual)-1])
	}
	if actual := DepthFirstSearch(g, v3, v5); actual[0] != v3 && actual[len(actual)-1] != v5 {
		t.Errorf("Expected first and last element (%v, %v), instead got (%v, %v)", v3, v5, actual[0], actual[len(actual)-1])
	}

	v6 := NewVertex(100)
	if actual := DepthFirstSearch(g, v1, v6); actual != nil {
		t.Errorf("Expected %v, instead got %v", nil, actual)
	}

	g.AddVertex(v6)
	if actual := DepthFirstSearch(g, v1, v6); actual != nil {
		t.Errorf("Expected %v, instead got %v", nil, actual)
	}
}
