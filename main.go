package main

import (
	"fmt"
	"github.com/scyanh/educative/sortingandsearching"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	//arr := []int{7, 8, 9, 10, 11, 12, 13, 14, 0, 3, 5, 6}

	idx := sortingandsearching.NewSASC().BinarySearch(arr, 4)
	fmt.Println("res=", idx)

}
