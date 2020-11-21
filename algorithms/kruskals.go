package algorithms

import (
	"fmt"
	dsu "graphs/dsu"
	edge "graphs/edge"
	"sort"
)

// Kruskals - Object
//	Contains the number of Verticies in the graph,
//	and the edges which make up the graph
type Kruskals struct {
	Verticies int
	Edges     []edge.Edge
}

// Construct the minimum spanning tree for this instances graph
// The main algorithm for Kruskal's Source:
//	https://en.wikipedia.org/wiki/Kruskal%27s_algorithm
func (k Kruskals) Construct() {
	index := 0

	// Sorting Reference found here:
	//	https://golangdocs.com/golang-sort-package
	sort.Sort(edge.Edges(k.Edges))

	subsets := make([]dsu.Subset, k.Verticies)
	for i := 0; i < k.Verticies; i++ {
		subsets[i] = dsu.Subset{Parent: i, Children: 0}
	}

	result := make([]edge.Edge, k.Verticies-1)
	index = 0
	i := 0

	for index < k.Verticies-1 {
		if i >= len(k.Edges) {
			break
		}

		next := k.Edges[i]
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
