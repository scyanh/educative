package main

import (
	"fmt"
	"github.com/scyanh/educative/insertionsort"
)

func main() {
	arr := []int{3, 2, 1, 5, 7, 2}

	arr = insertionsort.NewInsertionSort().Sort(arr)
	fmt.Println("arr=", arr)
}
