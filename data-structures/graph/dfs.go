package graph

import (
	"github.com/emredir/galgorithms/data-structures/stack"
)

// DepthFirstSearch explores a branch as far as possible until it
// reaches the end of the branch. At that point, it backtracks,
// and explores the next available branch.
func DepthFirstSearch(g *Graph, source, destination Vertex) []Vertex {
	s := stack.New()
	s.Push(source)

	visited := make(map[Vertex]bool)
	visited[source] = true

OUTER:
	for vertex := s.Peek(); vertex != nil && vertex != destination; vertex = s.Peek() {
		u := vertex.(Vertex)

		adjacents := g.Adj(u)
		if !(len(adjacents) > 0) {
			s.Pop()
			continue
		}
		for _, v := range adjacents {
			if _, ok := visited[v]; !ok {
				s.Push(v)
				visited[v] = true
				continue OUTER
			}
		}

		s.Pop()
	}

	if s.IsEmpty() {
		return nil
	}

	slice := s.ToSlice()
	vertices := make([]Vertex, len(slice))
	for i, e := range slice {
		vertices[i] = e.(Vertex)
	}

	return vertices
}
