package algorithms

import (
	"fmt"
	"graphs/dsu"
	node "graphs/nodes"
)

// Boruvkas object
type Boruvkas struct {
	Verticies int
	Graph     [][]int
}

// Construct c
func (b Boruvkas) Construct() {

	// 1 - Initialize a Forest
	forest := make([]node.Nodes, b.Verticies-1)
	subset := make([]dsu.Subset, b.Verticies)
	edges := makeGraph(b.Graph, b.Verticies)
	cheapest := make([]int, b.Verticies)

	// Set all trees in the forest to single components
	for i := 0; i < b.Verticies; i++ {
		subset[i] = dsu.Subset{Parent: i, Children: 0}
		cheapest[i] = -1
	}

	numOfForests := b.Verticies
	index := 0

	// While we have more than 1 forest
	for numOfForests > 1 {
		// Reinitialize cheapest array
		for i := range cheapest {
			cheapest[i] = -1
		}

		// Traverse and update the cheapest for each edge
		for i := range edges {
			u := dsu.Find(subset, edges[i].Src)
			v := dsu.Find(subset, edges[i].Dest)

			// If edge u and v are in the same component already
			if u == v {
				continue
			}

			cheapU := cheapest[u]
			cheapV := cheapest[v]

			// Check if this node is cheaper than previous nodes
			if cheapU == -1 || edges[cheapU].Cost > edges[i].Cost {
				cheapest[u] = i
			}

			// Check if there are cheaper nodes for this edge
			if cheapV == -1 || edges[cheapV].Cost > edges[i].Cost {
				cheapest[v] = i
			}
		}

		// for each vertex in the graph
		for i := 0; i < b.Verticies; i++ {

			// If there is a path to the node at i
			if cheapest[i] != -1 {
				u := dsu.Find(subset, edges[cheapest[i]].Src)
				v := dsu.Find(subset, edges[cheapest[i]].Dest)

				// check if the nodes are already in the same component
				if u == v {
					continue
				}

				// Add the edges to the forest
				forest[index] = edges[cheapest[i]]

				// Merge the 2 components
				dsu.Union(subset, u, v)

				// Decrease the number of forests
				numOfForests--
				index++
			}
		}
	}

	for _, anEdge := range forest {
		fmt.Printf("%d - %d | %d\n", anEdge.Src, anEdge.Dest, anEdge.Cost)
	}

}
