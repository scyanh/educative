package main

import (
	"github.com/scyanh/educative/concurrency"
)

func main() {
	//arr := []int{2, 3, 3, 3, 6, 9, 9}
	concurrency.NewQuitChannel().Racing()
}
