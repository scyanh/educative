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

func (bFS) LevelOrderTraversal(root *treeNode) [][]int {
	var res [][]int

	if root == nil {
		return [][]int{}
	}

	queue := []*treeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		// get the current size of the queue, this indicates the total number of nodes
		// that are part of current level
		sz := len(queue)
		level := []int{}

		for i := 0; i < sz; i++ {
			// remove a node
			node := queue[0]
			queue = queue[1:]

			// visit the node, collecting it into the output array
			level = append(level, node.Val)

			// insert children of the node into the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
