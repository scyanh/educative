package main

import (
	"fmt"
	"github.com/scyanh/educative/mergeintervals"
)

func main() {
	//arr := []int{2, 3, 3, 3, 6, 9, 9}
	//target := 11

	arr:=[][]int{
		{1,4},
		{2,6},
		{3,5},
	}
	res := mergeintervals.NewMergeIntervals().MergeAllOverlappingIntervals(arr)
	fmt.Println("res=", res)
}
