package graph

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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
	NewEmptyGraph(...uint16) error

	// NumEdges counts the number of edges in the given graph.
	// if its a matrix, you must count every element in the matrix.
	NumEdges() uint16

	// Marshal saves the graph to a .json file (ex: array of arrays)
	// returns the byte array of the json for whatever usage is necessary
	Marshal() ([]byte, error)

	// Unmarshal opens a .json file and initializes the graph object to the data of the file.
	Unmarshal(path string) error
}

// AdjBGraph is a representation of a bipartite graph G.
// It's an adjacency matrix of the form |X| x |Y|, where X and Y are vertex sets in G.
//
type AdjBGraph [][]uint16

// NewEmptyGraph implements the Graph interface by initializing a graph
// 				to be empty with the appropriate vertex lengths.
func (G *AdjBGraph) NewEmptyGraph(nums ...uint8) error {
	if len(nums) != 2 {
		return errors.New("must use 2 set sizes X and Y to initialize the bipartite graph")
	}
	if nums[0] == 0 || nums[1] == 0 {
		return errors.New("numbers must be non-zero")
	}
	*G = make([][]uint16, nums[0])
	for i := range *G {
		(*G)[i] = make([]uint16, nums[1])
	}
	return nil
}

// NumEdges implements the method in the Graph interface
// this calculates the number of edges by iterating through the array, and
// increments the count by the value in the adjacency matrix.
func (G AdjBGraph) NumEdges() uint16 {
	var count uint16 = 0
	for i := 0; i < len(G); i++ {
		for j := 0; j <= len(G[i]); j++ {
			count += G[i][j]
		}
	}
	return count
}

// Marshal implements the marshal function from the Graph interface
// saves an adjacency matrix as a JSON array
func (G AdjBGraph) Marshal() ([]byte, error) {
	b, err := json.Marshal(G)
	return b, err
}

// Unmarshal implements the Graph interface Unmarshal
// It opens a valid JSON file.
// "valid" means it is an array of arrays of type uint16
func (G AdjBGraph) Unmarshal(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &G)
	return err
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
type AdjVertexSet map[uint32]bool

// AdjMatching is implemented as a subgraph of AdjBGraph,
// as it uses a subset of edges of the given graph.
// The vertices can be ignored, we will only focus on the edges.
type AdjMatching struct {
	Repr AdjBGraph // representation of the matching
}

// Len gets the length (i.e. number of edges) in the matching.
func (M AdjMatching) Len() uint16 {
	return M.Repr.NumEdges()
}
