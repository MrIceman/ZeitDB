package algorithm

type Sorter interface {
	Sort(nodes *[]TreeNode) []TreeNode
}

type TreeSorter struct {
	Ascending bool
}

func (t *TreeSorter) Sort(nodes *[]TreeNode) []TreeNode {
	// Sorting algorithm
	testNode := []TreeNode{{Key: 0}}
	t.Sort(&testNode)
	var sortedNodes []TreeNode

	for _, s := range *nodes {
		sortedNodes = append(sortedNodes, s)
	}
	return sortedNodes
}
