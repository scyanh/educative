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
	arr := []int{1, 2, 3, 4, 5}
	//arr2 := []int{1, 2, 6, 7}
	//concurrency.NewBufferedChan().SimpleBuffered()
	//fmt.Println(datastructures.NewListChallenge().RemoveEven(arr))
	k := 5
	fmt.Println(datastructures.NewListChallenge().TwoSum(arr, k))
	fmt.Println(datastructures.NewListChallenge().TwoSum2(arr, k))

}
