package main

import (
	"fmt"
	"github.com/scyanh/educative/selectionsort"
)

func main() {
	arr := []int{3, 2, 1, 5, 7, 2}

	arr = selectionsort.NewSelectionSort().Sort(arr)
	fmt.Println("arr=", arr)
}
