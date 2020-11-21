package dsu

// NOTE
// The section was the most difficult part
// to implement at the time of this writing
// As such there is heavy inspiration from here:
//	https://cp-algorithms.com/data_structures/disjoint_set_union.html

// Subset - Object
//	Contains the Parent and the number of Children for a subset
type Subset struct {
	Parent, Children int
}

// Find - returns a nodes parent / source node
func Find(subsets []Subset, i int) int {
	if subsets[i].Parent != i {
		subsets[i].Parent = Find(subsets, subsets[i].Parent)
	}

	return subsets[i].Parent
}

// Union - merges 2 subsets into 1 subset
func Union(subsets []Subset, x, y int) {
	xroot := Find(subsets, x)
	yroot := Find(subsets, y)

	if subsets[xroot].Children < subsets[yroot].Children {
		subsets[xroot].Parent = yroot
	} else if subsets[xroot].Children > subsets[yroot].Children {
		subsets[yroot].Parent = xroot
	} else {
		subsets[yroot].Parent = xroot
		subsets[xroot].Children++
	}
}
