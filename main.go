package main

import (
	"fmt"
	"github.com/scyanh/educative/fastandslowpointers"
)

func main() {
	//arr := []int{2, 3, 3, 3, 6, 9, 9}
	//target := 11

	head:=fastandslowpointers.NewFastAndSlowPointers().NewListNode(1)
	head.Next=fastandslowpointers.NewFastAndSlowPointers().NewListNode(2)
	head.Next.Next=fastandslowpointers.NewFastAndSlowPointers().NewListNode(3)
	head.Next.Next.Next=fastandslowpointers.NewFastAndSlowPointers().NewListNode(4)
	head.Next.Next.Next.Next=fastandslowpointers.NewFastAndSlowPointers().NewListNode(5)
	head.Next.Next.Next.Next.Next=fastandslowpointers.NewFastAndSlowPointers().NewListNode(6)
	head.Next.Next.Next.Next.Next.Next=head.Next.Next
	res := fastandslowpointers.NewFastAndSlowPointers().FindCycleStart(head)
	fmt.Println("res=", res)

}
