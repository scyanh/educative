package main

import (
	"fmt"
	"github.com/scyanh/educative/recursion"
)

func main() {
	//arr := []int{3, 2, 1, 5, 7, 2}

	res := recursion.NewRecursiveFactorial().Factorial(5)
	fmt.Println("res=", res)
}
