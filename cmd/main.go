// --------------------------
// MATH-3150 final assignment
// --------------------------
// Author: Bailey Chittle
// Date: December 6, 2020
//
// This program implements option 4 of the final project.
//   It calculates the maximum matching of a given bipartite graph.
//   This is inputted either as a file, or through the command line.
//
package main

import (
	"fmt"
	"github.com/bachittle/graph-proj/graph"
)

func main() {
	var G graph.AdjBGraph
	(&G).Set([][]uint16{{1, 1}, {1, 1}})
	M := graph.EmptyMatch(&G)
	U := G.X
	fmt.Println("G:", G)
	fmt.Println("M:", M)
	vc, path := graph.AugmentingPath(&G, &M, &U)
	fmt.Println(vc, path)
}
