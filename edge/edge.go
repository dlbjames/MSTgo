package edge

// Edge - Object
// 	a single edge in the graph which contains the source and destination nodes
//	and the cost betweeb them
type Edge struct {
	Src  int
	Dest int
	Cost int
}

// Creating custom sorting for structs source:
//	https://thenotexpert.com/golang-sorting/

// Edges - The list of edges to sort
type Edges []Edge

// Len - the length of the list of edges
func (e Edges) Len() int { return len(e) }

// Swap - swap 2 elements in the list of edges
func (e Edges) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

// Less - find which element is less than another element in the list of edges
func (e Edges) Less(i, j int) bool { return e[i].Cost < e[j].Cost }

// MakeEdges - makes the list of edges from the adjacency matrix
func MakeEdges(graph [][]int, numOfVerticies int) []Edge {
	index := 0
	edges := make([]Edge, 0)
	for i := 0; i < numOfVerticies; i++ {
		for j := 0; j < numOfVerticies; j++ {
			// If no nodes were allowed to be 0
			//		use this: i*k.Verticies + j for each index
			// 0 means there is no edge between 2 nodes
			if graph[i][j] != 0 {
				edges = append(edges, Edge{Src: i, Dest: j, Cost: graph[i][j]})
				index++
			}
		}
	}

	return edges
}
