package fastandslowpointers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFastAndSlowPointers_HasCycle(t *testing.T) {
	head := NewFastAndSlowPointers().NewListNode(1)
	head.Next =NewFastAndSlowPointers().NewListNode(2)
	head.Next.Next =NewFastAndSlowPointers().NewListNode(3)
	head.Next.Next.Next =NewFastAndSlowPointers().NewListNode(4)
	head.Next.Next.Next.Next =NewFastAndSlowPointers().NewListNode(5)
	head.Next.Next.Next.Next.Next =NewFastAndSlowPointers().NewListNode(6)

	res := NewFastAndSlowPointers().HasCycle(head)
	expected := false
	require.Equal(t, expected, res)

	head.Next.Next.Next.Next.Next.Next = head.Next.Next
	res2 := NewFastAndSlowPointers().HasCycle(head)
	expected2 := true
	require.Equal(t, expected2, res2)
}
