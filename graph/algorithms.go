package graph

import (
//"fmt"
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
	vc [2]AdjVertexSet, // vertices in S, vertices in T
	augPath *AdjBGraph,
) {
	augPath = new(AdjBGraph)
	augPath.Repr = make([][]uint16, G.X.Len())
	for i := range augPath.Repr {
		augPath.Repr[i] = make([]uint16, G.Y.Len())
	}
	augPath.X.Init(G)
	augPath.Y.Init(G)
	// initialize S and T
	// S=U and T=empty

	// S => vertices in G.X
	S := *U
	// T => vertices in G.Y
	var T AdjVertexSet
	T.Init(G)

	// mark arrays with a vertex set
	var marked AdjVertexSet
	marked.Init(G)

	// continue in a loop while U - marked != 0
	for S.Minus(marked).Len() != 0 {
		// fmt.Println(S.Minus(marked).Len())
		// iterate through unmarked vertices
		for _, x := range S.Minus(marked).Repr.Keys() {
			//fmt.Println("x:", x)
			// x => vertex in S subset X
			for _, y := range G.Y.Repr.Keys() {
				// consider the neighbors of x
				// y => vertices which are a neighbour of x in Y
				// such that x,y is not an element of M
				//
				// if x,y is an edge in G | if x,y is NOT an edge in M
				//       V                    V
				if G.Repr[x][y] > 0 && M.Graph.Repr[x][y] == 0 {
					//fmt.Printf("%v: ", y)
					// check if y is saturated in M by any other edges
					saturated := false
					// for all weX
					for _, w := range G.X.Repr.Keys() {
						// w => vertices which are a neighbour of y in X, not including x
						//fmt.Println("w:", w)
						if w != x && M.Graph.Repr[w][y] == 1 {
							// y is saturated in M.
							//fmt.Println(y, "is saturated")
							// include y in T and w in S
							T.Repr[y] = true
							S.Repr[w] = true
							saturated = true
						}
					}
					if !saturated {
						//fmt.Println(y, "is NOT saturated")
						// report an M-augmenting path
						vc[0] = S
						vc[1] = T
						if marked.Len() == 0 {
							// set the path to be the X value
							augPath.X.Repr[x] = true
							augPath.Y.Repr[y] = true
							augPath.Repr[x][y] = 1
						} else {
							// use marked vertices?
							panic("implement this")
						}
						return
					}
				}
			}
			// mark x and iterate
			marked.Repr[x] = true
			// fmt.Println("}")
		}
		break
	}
	// return Tu(X-S) as minimum cover
	vc[0] = G.X.Minus(S)
	vc[1] = T
	augPath = nil
	return
}

// MaximumMatching uses the AugmentingPath algorithm to get the maximum
// matching in G, and returns it.
func MaximumMatching(G AdjBGraph) AdjMatching {
	M := EmptyMatch(&G)
	U := G.X
	//fmt.Println("G:", G)
	//fmt.Println("M:", M)
	_, augPath := AugmentingPath(&G, &M, &U)
	for augPath != nil {
		for _, x := range augPath.X.Repr.Keys() {
			for _, y := range augPath.Y.Repr.Keys() {
				M.Graph.Repr[x][y] = 1
				M.Graph.Y.Repr[y] = true
			}
			M.Graph.X.Repr[x] = true
		}
		U = U.Minus(M.Graph.X)
		//fmt.Println("U:", U)
		//fmt.Println("M:", M, "len:", M.Len())
		//fmt.Println("vc:", vc, "len:", vc[0].Len(), "+", vc[1].Len())
		//fmt.Println("path:", augPath)
		_, augPath = AugmentingPath(&G, &M, &U)
	}
	return M
}
