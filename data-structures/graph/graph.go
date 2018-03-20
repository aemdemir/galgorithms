package graph

import (
	"fmt"
)

// Weight is an alias for float64.
// It represents edge weights.
type Weight float64

// Vertex is a block of data.
type Vertex struct {
	Data int
}

// NewVertex returns a new vertex.
func NewVertex(data int) Vertex {
	return Vertex{data}
}

// String prints vertex in a nice format.
func (v Vertex) String() string {
	return fmt.Sprintf("%d", v.Data)
}

// Edge connects a vertex to other vertices.
type Edge struct {
	Source      Vertex
	Destination Vertex
}

// NewEdge returns a new edge between source 'u' and
// destination 'v'.
func NewEdge(u, v Vertex) Edge {
	return Edge{u, v}
}

// String prints edge in a nice format.
func (e Edge) String() string {
	return fmt.Sprintf("(%s) --- (%s)", e.Source.String(), e.Destination.String())
}

// Graph is a set of vertices paired with set of edges.
type Graph struct {
	adj map[Vertex]map[Vertex]Weight
}

// NewGraph returns a new graph.
func NewGraph() *Graph {
	return &Graph{adj: make(map[Vertex]map[Vertex]Weight)}
}

// Count returns the number of vertices in the graph.
func (g *Graph) Count() int {
	return len(g.adj)
}

// Contains tells whether the given vertex exists or not.
func (g *Graph) Contains(u Vertex) bool {
	_, ok := g.adj[u]
	return ok
}

// IsConnected checks if there is an edge between
// source 's' and destination 'v' vertices.
func (g *Graph) IsConnected(s, v Vertex) bool {
	_, ok := g.adj[s][v]
	return ok
}

// Adj returns slice of adjacent vertices of vertex u.
func (g *Graph) Adj(u Vertex) []Vertex {
	vertices := make([]Vertex, len(g.adj[u]))

	i := 0
	for v := range g.adj[u] {
		vertices[i] = v
		i++
	}

	return vertices
}

// AddVertexWith adds a new vertex with given data.
func (g *Graph) AddVertexWith(data int) Vertex {
	u := Vertex{data}
	g.AddVertex(u)
	return u
}

// AddVertex adds a new vertex to the graph.
func (g *Graph) AddVertex(u Vertex) {
	if g.adj[u] != nil {
		return
	}
	g.adj[u] = make(map[Vertex]Weight)
}

// AddUndirectedEdge adds a new undirected edge to the graph where
// u is source, v is destination and w is weight.
func (g *Graph) AddUndirectedEdge(u, v Vertex, w float64) {
	g.AddDirectedEdge(u, v, w)
	g.AddDirectedEdge(v, u, w)
}

// AddDirectedEdge adds a new directed edge to the graph where
// u is source, v is destination and w is weight.
func (g *Graph) AddDirectedEdge(u, v Vertex, w float64) {
	if !(g.Contains(u) && g.Contains(v)) {
		return
	}
	if g.IsConnected(u, v) {
		return
	}
	g.adj[u][v] = Weight(w)
}

// Weight returns the weight of the edge between source 'u' and
// destination 'v'.
func (g *Graph) Weight(u, v Vertex) float64 {
	return float64(g.adj[u][v])
}

// Edges returns edge list for given vertex.
func (g *Graph) Edges(u Vertex) []Edge {
	edges := make([]Edge, len(g.adj[u]))

	i := 0
	for v := range g.adj[u] {
		edges[i] = NewEdge(u, v)
		i++
	}

	return edges
}

// String prints graph in a nice format.
func (g *Graph) String() string {
	if !(g.Count() > 0) {
		return "Empty Graph"
	}

	str := ""
	counter := 0
	length := len(g.adj)

	for u, adjList := range g.adj {
		counter++
		i := 0
		str += fmt.Sprintf("{%v} --> [", u)

		for v := range adjList {
			if i != len(adjList)-1 {
				str += fmt.Sprintf("%v, ", v)
			} else {
				str += fmt.Sprintf("%v", v)
			}
			i++
		}

		if counter != length {
			str += "]\n"
		} else {
			str += "]"
		}
	}
	return str
}
