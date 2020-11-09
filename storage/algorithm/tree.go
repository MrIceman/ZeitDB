package algorithm

type TreeNode struct {
	Key   int32
	Level int
	// The node above the
	Predecessor *TreeNode
	// The node below
	Successor *TreeNode
	// The node on the  same level on the right side
	RightSibling *TreeNode
	// The node on the same level on the left side
	LeftSibling *TreeNode
}

type BTree struct {
	Root    *TreeNode
	Version string
}

func (tree *BTree) merge() {

}

func (tree *BTree) split() {

}

func (tree *BTree) Append(node *TreeNode) {
	// append node and then merge or split tree
}

func (tree *TreeNode) Update(node *TreeNode) {
	// update node and then merge or split tree

}

func (tree *TreeNode) Delete(node *TreeNode) {
	// delete node and then merge or split tree

}
