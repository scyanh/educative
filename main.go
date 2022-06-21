package main

import (
	"fmt"
	"github.com/scyanh/educative/datastructures"
)

func main() {
	arr := []int{1, 2, 5, 6}
	arr2 := []int{3, 4}

	res := datastructures.NewListChallenge().MergeTwoSortedList(arr, arr2)
	fmt.Println("res=", res)
}
