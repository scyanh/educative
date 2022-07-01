package main

import (
	"fmt"
	"github.com/scyanh/educative/sortingandsearching"
)

func main() {
	//arr := []int{1, 2, 3, 4, 5}
	arr := []int{1, 21, 3, 14, 5, 60, 7, 6}

	n1, n2 := sortingandsearching.NewSASC().FindTwoNumbersAddToK(arr, 81)
	fmt.Println("res=", n1, n2)

}
