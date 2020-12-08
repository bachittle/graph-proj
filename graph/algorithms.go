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
		for x := range S.Minus(marked).Repr {
			// x => vertex in S subset X
			fmt.Printf("N(%v) = {", x)
			for y := range G.Y.Repr {
				// consider the neighbors of x
				// y => vertices which are a neighbour of x in Y
				// such that x,y is not an element of M
				//
				// if x,y is an edge in G | if x,y is NOT an edge in M
				//       V                    V
				if G.Repr[x][y] > 0 && M.Graph.Repr[x][y] == 0 {
					fmt.Printf("%v, ", y)
					// check if y is saturated in M by any other edges
					saturated := false
					// for all weX
					for w := range M.Graph.X.Repr {
						// w => vertices which are a neighbour of y in X, not including x
						if w != x && G.Repr[w][y] == 1 {
							// y is saturated in M.
							fmt.Println("saturated")
							saturated = true
							// include y in T and w in S
							T.Repr[y] = true
							S.Repr[w] = true

						}
					}
					if !saturated {
						// report an M-augmenting path
					}
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
