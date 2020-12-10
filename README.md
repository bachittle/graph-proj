# graph-proj
MATH-3150 final project : Maximum Matchings

This final project implements option 4 given in my graph theory class: to implement an augmenting path algorithm, and to apply it to get a maximum matching of a bipartite graph. 

# important places:
- cmd/main.go     - where the interactive program is, try running that by calling 'go run main.go' or 'go build'
- graph/graph.go  - some graph data structures implemented in Go. The main ones are AdjBGraph (Adjacency Matrix Bipartite Graphs), 
AdjVertexSet (a vertex set of an adjacency matrix), and AdjMatching (matchings of AdjBGraph's)
- graph/algorithms.go - where the main algorithms (AugmentingPath, MaximumMatching) are implemented. This is where the meat is. 

# installing:
you need go to build from source.

go to the cmd folder, then type 'go build' to get it built on your system!

# some sample visuals

This program also creates visuals so you can easily see what is going on
It outputs in LaTeX and as PDF

Here are some of the main examples I tested:
- textbook example:
![textbook](https://i.imgur.com/aPzX2in.png)

- interesting small example that required using an augmenting path with length larger than 1
![weird_case](https://i.imgur.com/H8USvmW.png)

- big example: K10,10
![k10,10](https://i.imgur.com/bv2cezm.png)

