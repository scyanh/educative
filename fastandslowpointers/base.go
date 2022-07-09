package fastandslowpointers

type fastAndSlowPointers struct {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (fastAndSlowPointers) NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func NewFastAndSlowPointers() fastAndSlowPointers {
	return fastAndSlowPointers{}
}
