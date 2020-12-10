package graph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

// testing the maximum matching algorithm with different graph inputs
func TestMaximumMatching(t *testing.T) {
	b, err := ioutil.ReadFile("data/testGraphs.json")
	if err != nil {
		t.Fatal(err)
	}
	var tests map[string][][]uint16
	err = json.Unmarshal(b, &tests)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range tests {
		t.Log("testing", k)
		var G AdjBGraph
		(&G).Set(v)
		M := MaximumMatching(G)
		t.Log("M:", M)
		if k == "textbook_ex" {
			err = G.SavePDF(fmt.Sprint("./data/", k), &M, &[2]int{0, 3})
		} else {
			err = G.SavePDF(fmt.Sprint("./data/", k), &M, &[2]int{0, 0})
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TesAugPath(t *testing.T) {
	var mat [4][4]uint16
	mat[1][0] = 1
	mat[2][1] = 1
	mat[3][2] = 1

	mat2 := mat
	mat2[0][0] = 1
	mat2[1][1] = 1
	mat2[2][2] = 1
	mat2[3][3] = 1

	var smat [][]uint16
	for i := range mat {
		smat = append(smat, mat[i][:])
	}

	var smat2 [][]uint16
	for i := range mat2 {
		smat2 = append(smat2, mat2[i][:])
	}

	var G AdjBGraph
	err := G.Set(smat2)
	if err != nil {
		t.Fatal(err)
	}

	M := EmptyMatch(&G)
	M.Set(smat)
	fmt.Println("M:", M)
	U := G.X.Minus(M.Graph.X)
	//G.SavePDF("data/pathTest", &M, nil)
	_, p, _ := AugmentingPath(&G, &M, &U)
	fmt.Println(p)
}
