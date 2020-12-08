package graph

import (
	"fmt"
)

// AugmentingPath implements the algorithm outlined in a graph theory textbook:
// 		Douglas B. West: Introduction to Graph Theory (2nd edition)
// 		Algorithm 3.2.1
//
// Input: a bipartite graph, a matching M in the graph,
// 		and a vertex set U of M-unsaturated vertices.
//
// Output: A vertex cover of the result of the algorithm,
// 		and an augmenting path if the graph has one.
//
func AugmentingPath(G *AdjBGraph, M *AdjMatching, U *AdjVertexSet) (
	vertexCover AdjVertexSet,
	augPath *AdjBGraph,
) {
	// initialize S and T. S=U and T=empty
	S := *U
	var T AdjVertexSet
	T.Init(G)

	// mark arrays with a vertex set
	var marked AdjVertexSet

	// continue in a loop while U - marked != 0
	for S.Minus(marked).Len() != 0 {
		fmt.Println(S.Minus(marked).Len())
		// iterate through unmarked vertices
		for i := range S.Minus(marked).Repr {
			// i => vertex in S subset X
			// consider the neighbors of xeS
			fmt.Printf("N(%v) = {", i)
			for j := range G.Repr {
				if G.Repr[i][j] > 0 {
					fmt.Printf("%v, ", j)
				}
			}
			fmt.Println("}")
		}
		break
	}
	// return Tu(X-S) as minimum cover
	vertexCover = T.Union(G.X.Minus(S))
	return
}
