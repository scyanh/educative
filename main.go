package main

import (
	"fmt"
	"github.com/scyanh/educative/slidingwindow"
)

func main() {
	//arr := []int{1, 2, 3, 4, 5}
	//arr := []int{7, 8, 9, 10, 11, 12, 13, 14, 0, 3, 5, 6}

	//arr:=[]string{"abc", "bca", "cab", "cba", "xyz", "yzx"}

	//str:="00110011"
	//str:="clementisacap"
	//input := []int{2, 3, 1, 2, 4, 3}
	input:=[]string{"A","B","C","A","C"}
	//input:=[]string{"A", "B", "C", "B", "B", "C"}
	//k := 2
	//input := "aabccbbaab"

	res := slidingwindow.NewSlidingWindowChallenge().MaxFruitsIntoBasket(input)
	fmt.Println("res=", res)

}
