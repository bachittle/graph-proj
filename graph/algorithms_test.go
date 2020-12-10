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
		var G AdjBGraph
		(&G).Set(v)
		M := MaximumMatching(G)
		t.Log("M:", M)
		if k == "textbook_ex" {
			err = G.SavePDF(fmt.Sprint("./data/", k), &M, [2]int{0, 3})
		} else {
			err = G.SavePDF(fmt.Sprint("./data/", k), &M, [2]int{0, 0})
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}
