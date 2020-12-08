package graph

import (
	"errors"
)

// Graph is an abstract concept in mathematics that has a couple representations.
// ex: adjacency matrix, incidence matrix
//
// Either way, it should implement these methods.
type Graph interface {
	// NewGraph initializes an empty graph to an appropriate representation of the graph.
	//
	// An empty graph is defined as:
	// - A graph with no edges
	//
	// Normally, this is initializing an empty adjacency matrix.
	NewEmptyGraph(...uint8) error
}

// AdjBGraph is a representation of a bipartite graph G.
// It's an adjacency matrix of the form |X| x |Y|, where X and Y are vertex sets in G.
//
type AdjBGraph struct {
	Repr [][]uint8
}

// NewEmptyGraph implements the Graph interface by initializing a graph
// 				to be empty with the appropriate vertex lengths.
func (b *AdjBGraph) NewEmptyGraph(nums ...uint8) error {
	if len(nums) != 2 {
		return errors.New("must use 2 set sizes X and Y to initialize the bipartite graph")
	}
	if nums[0] == 0 || nums[1] == 0 {
		return errors.New("numbers must be non-zero")
	}
	b.Repr = make([][]uint8, nums[0])
	for i := range b.Repr {
		b.Repr[i] = make([]uint8, nums[1])
	}
	return nil
}

// AdjVertexSet is a set of vertices in a given graph G
// It maps the index of the vertex in the adjacency matrix with whether it exists in G (true/false).
//
// Ex: 3x3 adjacency matrix represents |V(G)| = 3
// var a AdjVertexSet
// a[0] => true
// a[2] => true
// a[3] => false
//
// for use in bipartite graphs, do the following:
// - for vertices in X, check with the usual a[n], where a is vertex set and n is col name
// - for vertices in Y, add |X| to n, i.e. do a[n + |X|].
//
type AdjVertexSet map[uint16]bool
