package main

import (
	"fmt"
	"github.com/scyanh/educative/twopointers"
)

func main() {
	arr := []int{2,5,9,11}
	target := 11

	res := twopointers.NewTwoPointers().PairWithTargetSum(arr, target)
	fmt.Println("res=", res)

}
