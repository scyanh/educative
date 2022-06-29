package main

import (
	"fmt"
	"github.com/scyanh/educative/datastructures"
)

func main() {
	//arr := []int{1, 2, 3, 4}
	arr := []int{0, 1, 2, 3}

	res := datastructures.NewListChallenge().RotateToRightK(arr, 1)
	fmt.Println("res=", res)


}
