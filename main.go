package main

import (
	"fmt"
	"github.com/scyanh/educative/datastructures"
)

func removeEven(arr []int) []int {
	var res []int

	for _, val := range arr {
		if val%2 != 0 {
			res = append(res, val)
		}
	}

	return res
}

func main() {
	//arr := []int{1, 2, 3, 4, 5}
	arr := []int{-2, 10, 7, -5, 15, 6}
	//arr := []int{1, -2, 3, 4}
	//arr := []int{5, 2, 3}
	//arr2 := []int{1, 2, 6, 7}
	//concurrency.NewBufferedChan().SimpleBuffered()
	//fmt.Println(datastructures.NewListChallenge().RemoveEven(arr))
	//k := 5

	fmt.Println(datastructures.NewListChallenge().FindMaxSumSublist(arr))
	fmt.Println(datastructures.NewListChallenge().FindMaxSumSublist2(arr))

}
