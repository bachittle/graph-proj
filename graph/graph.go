package graph

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
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

	// Set graph by its representation
	Set(interface{}) error

	// NumEdges counts the number of edges in the given graph.
	// if its a matrix, you must count every element in the matrix.
	NumEdges() uint16

	// Marshal saves the graph to a .json file (ex: array of arrays)
	// returns the byte array of the json for whatever usage is necessary
	Marshal() ([]byte, error)

	// Unmarshal opens a .json file and initializes the graph object to the data of the file.
	Unmarshal(r io.Reader) error
}

// AdjBGraph is a representation of a bipartite graph G.
// It's an adjacency matrix of the form |X| x |Y|, where X and Y are vertex sets in G.
type AdjBGraph struct {
	Repr [][]uint16
	X    AdjVertexSet
	Y    AdjVertexSet
}

// EmptyGraph takes another graph and initializes a copied empty graph
func EmptyGraph(G AdjBGraph) AdjBGraph {
	G.NewEmptyGraph(uint16(len(G.X.Repr)), uint16(len(G.Y.Repr)))
	return G
}

// NewEmptyGraph implements the Graph interface by initializing a graph
// to be empty with the appropriate vertex lengths.
func (G *AdjBGraph) NewEmptyGraph(nums ...uint16) error {
	if len(nums) != 2 {
		return errors.New("must use 2 set sizes X and Y to initialize the bipartite graph")
	}
	if nums[0] == 0 || nums[1] == 0 {
		// zeroes, initialize a nil array
		return errors.New("initialized a nil array")
	}
	G.Repr = make([][]uint16, nums[0])
	for i := range G.Repr {
		G.Repr[i] = make([]uint16, nums[1])
	}
	G.X.Init(G)
	G.Y.Init(G)
	return nil
}

// Set a graph to its representation
func (G *AdjBGraph) Set(adjMat interface{}) error {
	mat, ok := adjMat.([][]uint16)
	if !ok {
		return errors.New("not uint16")
	}
	G.Repr = mat
	G.X.Init(G)
	G.Y.Init(G)
	for i := 0; i < len(mat); i++ {
		G.X.Repr[uint16(i)] = true
	}
	if len(mat) == 0 {
		// zero valued matrix
		return nil
	}
	for j := 0; j < len(mat[0]); j++ {
		G.Y.Repr[uint16(j)] = true
	}
	return nil
}

// NumEdges implements the method in the Graph interface
// this calculates the number of edges by iterating through the array, and
// increments the count by the value in the adjacency matrix.
func (G AdjBGraph) NumEdges() uint16 {
	var count uint16 = 0
	for i := range G.X.Repr {
		for j := range G.Y.Repr {
			count += G.Repr[i][j]
		}
	}
	return count
}

// Marshal implements the marshal function from the Graph interface
// saves an adjacency matrix as 2 types:
// - "json" => json file
// - "tex" => tex file of the graph (for visualization)
func (G AdjBGraph) Marshal(Types ...string) (b []byte, err error) {
	for _, t := range Types {
		switch t {
		case "json":
			b, err = json.Marshal(G)
			break
		case "tex":
			b, err = texMarshal(G, nil, nil)
			break
		}
		if err != nil {
			return
		}
	}
	return
}

// converts a given graph G to a tex file (for visualization)
// saves to byte array, and you can do whatever you want with it (save to file, to RAM, etc.)
func texMarshal(G AdjBGraph, M *AdjMatching, offset *[2]int) (b []byte, err error) {
	b = append(b, []byte(`\documentclass{article}
\usepackage{tikz}
\begin{document}

\begin{tikzpicture}[every node/.style={circle,inner sep=1pt,fill=black}]
`)...)

	// node initialization

	for x := range G.X.Repr {
		xOff := int(x)
		if offset != nil {
			xOff += offset[0]
		}
		b = append(b, []byte(fmt.Sprintf("\t\\node (X%d) at (%d, 1) {};\n", x, xOff))...)
	}
	for y := range G.Y.Repr {
		yOff := int(y)
		if offset != nil {
			yOff += offset[1]
		}
		b = append(b, []byte(fmt.Sprintf("\t\\node (Y%d) at (%d, 0) {};\n", y, yOff))...)
	}

	// edge initialization
	for x := range G.X.Repr {
		for y := range G.Y.Repr {
			if G.Repr[x][y] > 0 {
				style := ""
				if M != nil && M.Graph.Repr[x][y] > 0 {
					style = "[red, very thick]"
				}
				b = append(b, []byte(fmt.Sprintf("\t\\draw %v (X%d) -- (Y%d)\n;", style, x, y))...)
			}
		}
	}
	b = append(b, []byte(`\end{tikzpicture}
\end{document}`)...)
	return
}

// SavePDF takes the graph and outputs in PDF form for visualization
func (G AdjBGraph) SavePDF(filename string, M *AdjMatching, offset *[2]int) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	basename := filepath.Base(filename)
	err = os.Mkdir(filename, 0644)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	err = os.Chdir(filename)
	if err != nil {
		return err
	}
	b, err := texMarshal(G, M, offset)
	if err != nil {
		_ = os.Chdir(cwd)
		return err
	}
	f, err := os.OpenFile(fmt.Sprint(basename, ".tex"),
		os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		_ = os.Chdir(cwd)
		return err
	}
	f.Write(b)
	cmd := exec.Command("pdflatex", fmt.Sprint(basename, ".tex"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		_ = os.Chdir(cwd)
		return err
	}
	f, err = os.OpenFile(fmt.Sprint(basename, ".log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Write(out)
	_ = os.Chdir(cwd)
	return err
}

// Unmarshal implements the Graph interface Unmarshal
// It opens a valid JSON file (not latex, that would be counterproductive).
// "valid" means it is an array of arrays of type uint16
func (G AdjBGraph) Unmarshal(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &G)
	return err
}

// VMap maps vertices to a boolean (whether it exists or not)
type VMap map[uint16]bool

// Keys gets the keys of a map
func (v VMap) Keys() (a []uint16) {
	var b []int
	for k := range v {
		b = append(b, int(k))
	}
	sort.Ints(b)
	for _, v := range b {
		a = append(a, uint16(v))
	}
	return
}

// AdjVertexSet is a set of vertices in a given graph G
// It maps the index of the vertex in the adjacency matrix with whether it exists in G (true/false).
//
// Ex: 3x3 adjacency matrix represents |V(G)| = 3
// var a AdjVertexSet
// a.Repr[0] => true
// a.Repr[2] => true
// a.Repr[3] => false
//
// for use in bipartite graphs, do the following:
// - for vertices in X, check with the usual a.Repr[n], where a is vertex set and n is col name
// - for vertices in Y, add |X| to n, i.e. do a[n + |X|].
//
type AdjVertexSet struct {
	Repr   VMap       // representation of a vertex set is a map
	Parent *AdjBGraph // the vertex set is for this graph
}

// Init sets the nil map to have zeroed values
func (A *AdjVertexSet) Init(parent *AdjBGraph) {
	A.Repr = make(VMap)
	A.Parent = parent
}

// Len gets the length (number of vertices) in the set
func (A AdjVertexSet) Len() uint16 {
	return uint16(len(A.Repr))
}

// Union performs a set union between two vertex sets
func (A AdjVertexSet) Union(B AdjVertexSet) (AuB AdjVertexSet) {
	AuB.Init(A.Parent)
	for k := range A.Repr {
		AuB.Repr[k] = true
	}
	for k := range B.Repr {
		AuB.Repr[k] = true
	}
	return
}

// Insec performs a set intersection between two vertex sets
func (A AdjVertexSet) Insec(B AdjVertexSet) (AnB AdjVertexSet) {
	AnB.Init(A.Parent)
	for k := range A.Repr {
		if B.Repr[k] == true {
			AnB.Repr[k] = true
		}
	}
	return
}

// Minus performs a set subtraction between two vertex sets
func (A AdjVertexSet) Minus(B AdjVertexSet) (AmB AdjVertexSet) {
	AmB.Init(A.Parent)
	for k := range A.Repr {
		if B.Repr[k] == false {
			AmB.Repr[k] = true
		}
	}
	return
}

// AdjMatching is implemented as a subgraph of AdjBGraph,
// as it uses a subset of edges of the given graph.
// The vertices can be ignored, we will only focus on the edges.
type AdjMatching struct {
	Graph  AdjBGraph  // underlying GRAPHICAL representation of the matching
	Parent *AdjBGraph // the matching is associated with this graph
}

// EmptyMatch initializes an empty match associated with graph G
func EmptyMatch(G *AdjBGraph) (M AdjMatching) {
	M.Graph = EmptyGraph(*G)
	M.Parent = G
	return
}

// Len gets the length (i.e. number of edges) in the matching.
func (M AdjMatching) Len() uint16 {
	return M.Graph.NumEdges()
}

// Set sets the graph representation of the matching. It only considers edges.
func (M *AdjMatching) Set(adjMat [][]uint16) {
	M.Graph.Set(adjMat)
	var X AdjVertexSet
	X.Init(&M.Graph)
	var Y AdjVertexSet
	Y.Init(&M.Graph)
	for i := range M.Graph.X.Repr {
		for j := range M.Graph.Repr {
			if M.Graph.Repr[i][j] > 0 {
				X.Repr[uint16(i)] = true
				Y.Repr[uint16(j)] = true
			}
		}
	}
	M.Graph.X = X
	M.Graph.Y = Y
}
