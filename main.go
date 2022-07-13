package main

import (
	"fmt"
	"github.com/scyanh/educative/bfs"
)

func main() {
	//arr := []int{2, 3, 3, 3, 6, 9, 9}
	//target := 11

	/*arr:=[][]int{
		{1,3},
		{5,7},
		{8,12},
	}*/

	tbfs := bfs.NewBFS()
	root := tbfs.NewTreeNode(12)
	root.Left = tbfs.NewTreeNode(7)
	root.Right = tbfs.NewTreeNode(1)
	root.Left.Left = tbfs.NewTreeNode(9)
	root.Left.Left = tbfs.NewTreeNode(10)
	root.Left.Right = tbfs.NewTreeNode(5)

	res := tbfs.Traverse(root)
	fmt.Println("res=", res)
}
