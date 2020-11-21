package algorithms

import (
	"fmt"
	dsu "graphs/dsu"
	node "graphs/nodes"
	"sort"
)

// Kruskals object
type Kruskals struct {
	Verticies int
	Graph     [][]int
}

// MakeGraph - makes the list of edges from the adjacency matrix
func makeGraph(graph [][]int, numOfVerticies int) []node.Nodes {
	index := 0
	edges := make([]node.Nodes, 0)
	for i := 0; i < numOfVerticies; i++ {
		for j := 0; j < numOfVerticies; j++ {
			// If no nodes were allowed to be 0
			//		use this: i*k.Verticies + j for each index
			// 0 means there is no edge between 2 nodes
			if graph[i][j] != 0 {
				edges = append(edges, node.Nodes{Src: i, Dest: j, Cost: graph[i][j]})
				index++
			}
		}
	}

	return edges
}

// Construct the minimum spanning tree for this instances graph
// The main algorithm for Kruskal's Source:
//	https://en.wikipedia.org/wiki/Kruskal%27s_algorithm
func (k Kruskals) Construct() {
	edges := makeGraph(k.Graph, k.Verticies)
	index := 0

	// Sorting Reference found here:
	//	https://golangdocs.com/golang-sort-package
	sort.Sort(node.Edges(edges))

	subsets := make([]dsu.Subset, k.Verticies)
	for i := 0; i < k.Verticies; i++ {
		subsets[i] = dsu.Subset{Parent: i, Children: 0}
	}

	result := make([]node.Nodes, k.Verticies-1)
	index = 0
	i := 0

	for index < k.Verticies-1 {
		if i >= len(edges) {
			break
		}

		next := edges[i]
		x := dsu.Find(subsets, next.Src)
		y := dsu.Find(subsets, next.Dest)

		if x != y {
			result[index] = next
			dsu.Union(subsets, x, y)
			index++
		}

		i++
	}

	for _, anEdge := range result {
		fmt.Printf("%d - %d | %d\n", anEdge.Src, anEdge.Dest, anEdge.Cost)
	}
}

// Sort all the edges
// Take the edge with the lowest weight and add it to the spanning tree.
//	If adding the edge create a cycle, remove the added edge.
//		This is the DSU Structures
//	Keep adding the edges until we reach all the verticies
