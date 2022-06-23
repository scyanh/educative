package main

import (
	"fmt"
	"github.com/scyanh/educative/datastructures"
)

func main() {
	arr := []int{1, 2, 3, 4}
	k := 6

	n1, n2, err := datastructures.NewListChallenge().AddTwoNumbersAddToK(arr, k)
	if err!=nil{
		fmt.Println("err=",err)
	}else{
		fmt.Println("res=", n1, n2)
	}

}
