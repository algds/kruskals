package kruskals

import (
	"sort"

	"github.com/algds/uf"
)

// WeightedEdge is an undirected edge between vertices and a weight.
type WeightedEdge interface {
	// From returns the integer identifier of the first vertex.
	From() int
	// To returns the integer identifier of the second vertex.
	To() int
	// Weight returns the integer identifier of the weight/cost.
	Weight() int
}

// SimpleWeightedEdge is a simple implementation of the interface.
// Initialize with 'F' (From), 'T' (To), 'W' (Weight).
type SimpleWeightedEdge struct {
	F, T, W int
}

// From returns the integer identifier of the first vertex.
func (s SimpleWeightedEdge) From() int {
	return s.F
}

// To returns the integer identifier of the second vertex.
func (s SimpleWeightedEdge) To() int {
	return s.T
}

// Weight returns the integer identifier of the weight/cost.
func (s SimpleWeightedEdge) Weight() int {
	return s.W
}

// MinimumSpanningTree returns the fewest edges to span across the graph
// with the minimum cost based on weights.
func MinimumSpanningTree(edges []WeightedEdge) []WeightedEdge {
	vertices := make(map[int]struct{})
	for _, e := range edges {
		vertices[e.From()] = struct{}{}
		vertices[e.To()] = struct{}{}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight() < edges[j].Weight()
	})
	u := uf.New(len(vertices))
	result := make([]WeightedEdge, 0, len(edges))
	for _, e := range edges {
		if !u.Connected(e.From(), e.To()) {
			u.Union(e.From(), e.To())
			result = append(result, e)
		}
	}
	return result
}
