package main

import (
	"fmt"
	"github.com/scyanh/educative/datastructures"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}

	res := datastructures.NewListChallenge().ArrangeMaxMinElements(arr)
	fmt.Println("res=", res)

}
