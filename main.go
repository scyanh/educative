package main

import (
	"fmt"
	"github.com/scyanh/educative/datastructures"
)

func main() {
	//arr := []int{1, 2, 3, 4}
	arr := []int{0, 1, 2, 3}

	res := datastructures.NewListChallenge().FindProductAllElements(arr)
	fmt.Println("res=", res)


}
