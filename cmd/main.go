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
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bachittle/graph-proj/graph"
	"os"
	"strings"
)

func main() {
	fmt.Println(`Match Maker
Written by: Bailey Chittle
----------------------------

Please enter a bipartite graph as an adjacency matrix of size |X|x|Y|.
It parses in JSON format, so input is as JSON. 
ex:
>>[
>>[1, 1],
>>[1, 1]
>>]
>>!
this creates K2,2.

End your input with the character: !
to exit, press Ctrl+C`)
	var G graph.AdjBGraph
	for {
		var mat [][]uint16
		reader := bufio.NewReader(os.Stdin)
		/*
			for i := range mat {
				mat[i] = make([]uint16, 10)
				for j := range mat[i] {
					mat[i][j] = 1
				}
			}
		*/
		var buf []byte
		for {
			fmt.Print(">>")
			line, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if strings.TrimSpace(line) == "!" {
				break
			}
			buf = append(buf, []byte(fmt.Sprintln(line))...)
		}
		err := json.Unmarshal(buf, &mat)
		if err != nil {
			fmt.Println("json unmarshal error", err)
			os.Exit(1)
		}

		(&G).Set(mat)
		M := graph.MaximumMatching(G)
		fmt.Println("maximum matching:", M)
		fmt.Print("print to pdf? (y/n):")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if strings.ToLower(line[:1]) == "y" {
			fmt.Println("give a new directory name to save tex and pdf files. If not a path, will save in save directory as executable. ")
			fmt.Print(">>")
			line, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if len(line) == 0 {
				fmt.Println("no filename given, aborting...")
			} else {
				line = strings.TrimSpace(line)
				fmt.Println("saving file to", line)
				err := G.SavePDF(line, &M, nil)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
