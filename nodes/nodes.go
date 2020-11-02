package nodes

// Nodes node
type Nodes struct {
	Src  int
	Dest int
	Cost int
}

// Creating custom sorting for structs source:
//	https://thenotexpert.com/golang-sorting/

// Edges e
type Edges []Nodes

// Len l
func (e Edges) Len() int { return len(e) }

// Swap s
func (e Edges) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

// Less l
func (e Edges) Less(i, j int) bool { return e[i].Cost < e[j].Cost }
