package algorithms

import "fmt"

// INF Constant for Infinite weight edges
var INF int = 9999999

// Prims The struct for Prims without a union find structure
type Prims struct {
	Edges     int
	Verticies int
	Graph     [][]int
	Seen      []bool
}

// Construct the MST
func (p Prims) Construct() {
	p.Seen[0] = true
	p.Edges = 0

	// For each edge in the graph
	for p.Edges < p.Verticies-1 {
		min := INF
		src, dest := 0, 0
		// For each vertex
		for i := 0; i < p.Verticies; i++ {
			if p.Seen[i] {
				// Check if we have been to that node
				//	if not, check if the node connects to this value
				// 		update:
				//			The min cost to the cost of this connection
				//			The src node to the parent of this node
				//			The destination node to this node
				for j := 0; j < p.Verticies; j++ {
					if !p.Seen[j] && p.Graph[i][j] != 0 {
						if min > p.Graph[i][j] {
							src, dest, min = i, j, p.Graph[i][j]
						}
					}
				}
			}
		}

		fmt.Printf("%d - %d | %d\n", src, dest, p.Graph[src][dest])
		p.Seen[dest] = true
		p.Edges++
	}
}
