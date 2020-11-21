package nodes

// Nodes node which specifies a source, destination, and cost between 2 edges
type Nodes struct {
	Src  int
	Dest int
	Cost int
}

// Creating custom sorting for structs source:
//	https://thenotexpert.com/golang-sorting/

// Edges e the list of edges to sort
type Edges []Nodes

// Len - the length of the list of edges
func (e Edges) Len() int { return len(e) }

// Swap - swap 2 elements in the list of edges
func (e Edges) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

// Less - find which element is less than another element in the list of edges
func (e Edges) Less(i, j int) bool { return e[i].Cost < e[j].Cost }
