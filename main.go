package main

import (
	"fmt"
	"github.com/scyanh/educative/datastructures"
)

func main() {
	//arr := []int{1, 2, 3, 4, 5}
	arr := []int{-2,10,7,-5, 15,6}

	res, sum := datastructures.NewListChallenge().FindMaxSumSublist(arr)
	fmt.Println("res=", res, sum)

}
