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
	mat := make([][]uint16, 10)
	for i := range mat {
		mat[i] = make([]uint16, 10)
		for j := range mat[i] {
			mat[i][j] = 1
		}
	}

	(&G).Set(mat)
	M := graph.MaximumMatching(G)
	fmt.Println(M)
	err := G.SavePDF("../graph/data/test", &M)
	if err != nil {
		panic(err)
	}

}
