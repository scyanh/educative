package bfs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBFS_LevelOrderTraversal(t *testing.T) {
	tbfs := NewBFS()
	root := tbfs.NewTreeNode(12)
	root.Left = tbfs.NewTreeNode(7)
	root.Right = tbfs.NewTreeNode(1)
	root.Left.Left = tbfs.NewTreeNode(9)
	root.Right.Left = tbfs.NewTreeNode(10)
	root.Right.Right = tbfs.NewTreeNode(5)

	res := tbfs.LevelOrderTraversal(root)
	expected := [][]int{
		{12},
		{7, 1},
		{9, 10, 5},
	}

	require.Equal(t, expected, res)
}
