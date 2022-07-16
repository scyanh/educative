package main

import (
	"fmt"
	"github.com/scyanh/educative/subsets"
)

func main() {
	//arr := []int{2, 3, 3, 3, 6, 9, 9}

	//target := 11

	/*arr:=[][]int{
		{1,3},
		{5,7},
		{8,12},
	}*/

	//arr := []int{3, 1, 3}
	arr := []int{1, 3, 3, 5}
	res:= subsets.NewSubset().FindSubsetsWithoutDuplicates(arr)

	fmt.Println("res=", res)
}
