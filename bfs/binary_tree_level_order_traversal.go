package bfs

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func (bFS) NewTreeNode(val int) *treeNode {
	tn := &treeNode{
		Val: val,
	}
	return tn
}

func (bFS) Traverse(root *treeNode) [][]int {
	var res [][]int



	return res
}
