package main

import (
	"fmt"
	"github.com/scyanh/educative/recursion"
)

func main() {
	//arr := []int{3, 2, 1, 5, 7, 2}

	res := recursion.NewRecursionPower().Power(2, 6)
	fmt.Println("res=", res)
}
