package graph

import (
	"github.com/emre-demir/go-algorithms/data-structures/queue"
)

// BreadthFirstSearch visits vertices in correct order. The first
// vertex to visit is always the source. After that, it will explore the
// source vertex's neighbors, then their neighbors and so on.
func BreadthFirstSearch(g *Graph, source, destination Vertex) []Vertex {
	q := queue.New()
	q.Enqueue(source)

	visits := make(map[Vertex]Vertex)
	visits[source] = source

	for !q.IsEmpty() {
		u := q.Dequeue().(Vertex)

		if u == destination {
			return path(source, destination, visits)
		}

		for _, v := range g.Adj(u) {
			if _, visited := visits[v]; !visited {
				q.Enqueue(v)
				visits[v] = u
			}
		}
	}

	return nil
}

func path(s, v Vertex, predecessor map[Vertex]Vertex) []Vertex {
	path := make([]Vertex, 0)

	for s != v {
		path = append([]Vertex{v}, path...)
		v = predecessor[v]
	}

	path = append([]Vertex{s}, path...)

	return path
}
