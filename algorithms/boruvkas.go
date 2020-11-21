package algorithms

import (
	"fmt"
	"graphs/dsu"
	edge "graphs/edge"
)

// Boruvkas - Object
//	Contains the number of verticies in the graph
//	and the edges which make up the graph
type Boruvkas struct {
	Verticies int
	Edges     []edge.Edge
}

// Construct the MST for Boruvkas Algorithm
func (b Boruvkas) Construct() {

	// 1 - Initialize a Forest
	forest := make([]edge.Edge, b.Verticies-1)
	subset := make([]dsu.Subset, b.Verticies)
	cheapest := make([]int, b.Verticies)

	// Set all trees in the forest to single components
	for i := 0; i < b.Verticies; i++ {
		subset[i] = dsu.Subset{Parent: i, Children: 0}
		cheapest[i] = -1
	}

	numOfForests := b.Verticies
	index := 0
	pass := 0

	// While we have more than 1 forest
	for numOfForests > 1 {
		if pass > len(b.Edges) {
			break
		}

		// Reinitialize cheapest array
		for i := range cheapest {
			cheapest[i] = -1
		}

		// Traverse and update the cheapest for each edge
		for i := range b.Edges {
			u := dsu.Find(subset, b.Edges[i].Src)
			v := dsu.Find(subset, b.Edges[i].Dest)

			// If edge u and v are in the same component already
			if u == v {
				continue
			}

			cheapU := cheapest[u]
			cheapV := cheapest[v]

			// Check if this node is cheaper than previous nodes
			if cheapU == -1 || b.Edges[cheapU].Cost > b.Edges[i].Cost {
				cheapest[u] = i
			}

			// Check if there are cheaper nodes for this edge
			if cheapV == -1 || b.Edges[cheapV].Cost > b.Edges[i].Cost {
				cheapest[v] = i
			}
		}

		// for each vertex in the graph
		for i := 0; i < b.Verticies; i++ {

			// If there is a path to the node at i
			if cheapest[i] != -1 {
				u := dsu.Find(subset, b.Edges[cheapest[i]].Src)
				v := dsu.Find(subset, b.Edges[cheapest[i]].Dest)

				// check if the nodes are already in the same component
				if u == v {
					continue
				}

				// Add the edges to the forest
				forest[index] = b.Edges[cheapest[i]]

				// Merge the 2 components
				dsu.Union(subset, u, v)

				// Decrease the number of forests
				numOfForests--
				index++
			}
		}
		pass++
	}

	for _, anEdge := range forest {
		fmt.Printf("%d - %d | %d\n", anEdge.Src, anEdge.Dest, anEdge.Cost)
	}
}
