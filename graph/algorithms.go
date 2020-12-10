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
	order []uint16, // path order, ex: [2, 0, 0, 1]
	//                                  ^X ^Y ^X ^Y
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
	var S AdjVertexSet
	S.Init(G)
	for k := range U.Repr {
		S.Repr[k] = true
	}
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
		xQueue := S.Repr.Keys()
		for len(xQueue) > 0 {
			x := xQueue[0]
			xQueue = xQueue[1:]
			//for _, x := range S.Minus(marked).Repr.Keys() {
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
							//fmt.Println("y =", y, "is saturated")
							// include y in T and w in S
							T.Repr[y] = true
							if G.X.Minus(*U).Repr[w] == true {
								// if element is not in queue, add it
								xQueue = append(xQueue, w)
							}
							S.Repr[w] = true
							saturated = true
						}
					}
					if !saturated {
						//fmt.Println("y =", y, "is NOT saturated")
						// report an M-augmenting path
						vc[0] = S
						vc[1] = T
						//fmt.Println("G:", G.X)
						//fmt.Println("U:", U)
						//n := x
						if G.X.Minus(*U).Repr[x] == false {
							// augmented path is only length 1
							// set the path to be the X value
							augPath.X.Repr[x] = true
							augPath.Y.Repr[y] = true
							augPath.Repr[x][y] = 1
							order = append(order, x)
							order = append(order, y)
						} else {
							// augmented path has length larger than 1
							// follow the augmented path
							//fmt.Println("marked:", marked)
							//fmt.Println("S:", S)
							//fmt.Println("T:", T)
							n := marked.Insec(*U)
							//fmt.Println("n:", n)
							if n.Len() != 1 {
								panic("implement this 2.0")
							}
							// m cycles between x and y values to find the augmented path
							var m uint16
							for k := range n.Repr {
								m = k
							}
							order = append(order, m)
							buildPath := func(isX bool, v AdjVertexSet) (uint16, bool) {
								var m2 uint16
								found := false
								var w AdjVertexSet
								if isX {
									w = v.Minus(augPath.Y)
								} else {
									w = v.Minus(augPath.X)
								}
								for m2 = range w.Repr {
									//fmt.Println("m", m)
									//fmt.Println("m2", m2)
									if isX {
										if G.Repr[m][m2] > 0 {
											// path exists
											augPath.X.Repr[m] = true
											augPath.Y.Repr[m2] = true
											augPath.Repr[m][m2] = 1
											found = true
											break
										}
									} else {
										//fmt.Println("in path?")
										if M.Graph.Repr[m2][m] > 0 {
											// path exists
											augPath.X.Repr[m2] = true
											augPath.Y.Repr[m] = true
											augPath.Repr[m2][m] = 1
											found = true
											break
										}
									}
								}
								if found {
									order = append(order, m2)
								}
								return m2, found
							}
							count := 0
							for {
								//fmt.Println("count", count%2 == 0)
								var found bool
								m, found = buildPath(count%2 == 0, T)
								if !found {
									m, found = buildPath(count%2 == 0, G.Y)
									//fmt.Println(m, found)
									if !found {
										panic("could not build augmented path")
									}
									//fmt.Println("augPath", augPath)
									//fmt.Println("order", order)
									break
								}
								count++
							}
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
	_, augPath, o := AugmentingPath(&G, &M, &U)
	for augPath != nil {
		/*
			for _, x := range augPath.X.Repr.Keys() {
				for _, y := range augPath.Y.Repr.Keys() {
					M.Graph.Repr[x][y] = 1
					M.Graph.Y.Repr[y] = true
				}
				M.Graph.X.Repr[x] = true
			}
		*/
		for i := 0; i < len(o)-1; i++ {
			//fmt.Println(o[i], o[i+1])
			if i%2 == 0 {
				M.Graph.Repr[o[i]][o[i+1]] = 1
				M.Graph.X.Repr[o[i]] = true
				M.Graph.Y.Repr[o[i+1]] = true
			} else {
				M.Graph.Repr[o[i+1]][o[i]] = 0
				M.Graph.X.Repr[o[i+1]] = false
				M.Graph.Y.Repr[o[i]] = false
			}
		}
		U = U.Minus(M.Graph.X)
		//fmt.Println("U:", U)
		//fmt.Println("M:", M, "len:", M.Len())
		//fmt.Println("vc:", vc, "len:", vc[0].Len(), "+", vc[1].Len())
		//fmt.Println("path:", augPath)
		_, augPath, o = AugmentingPath(&G, &M, &U)
	}
	return M
}
