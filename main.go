package main

import (
	"fmt"
	"github.com/scyanh/educative/mergeintervals"
)

func main() {
	//arr := []int{2, 3, 3, 3, 6, 9, 9}
	//target := 11

	arr:=[][]int{
		{1,3},
		{5,7},
		{8,12},
	}

	newInterval:=[]int{4,6}
	res := mergeintervals.NewMergeIntervals().InsertInterval(arr, newInterval)
	fmt.Println("res=", res)
}
