package main

import (
	"fmt"
	"github.com/scyanh/educative/twopointers"
)

func main() {
	arr := []int{2, 3, 3, 3, 6, 9, 9}
	//target := 11

	res := twopointers.NewTwoPointers().RemoveDuplicates(arr)
	fmt.Println("res=", res)

}
