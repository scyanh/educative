package fastandslowpointers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFastAndSlowPointers_LengthOfCycle(t *testing.T) {
	head := NewFastAndSlowPointers().NewListNode(1)
	head.Next =NewFastAndSlowPointers().NewListNode(2)
	head.Next.Next =NewFastAndSlowPointers().NewListNode(3)
	head.Next.Next.Next =NewFastAndSlowPointers().NewListNode(4)
	head.Next.Next.Next.Next =NewFastAndSlowPointers().NewListNode(5)
	head.Next.Next.Next.Next.Next =NewFastAndSlowPointers().NewListNode(6)

	head.Next.Next.Next.Next.Next.Next = head.Next.Next
	res := NewFastAndSlowPointers().HasCycle(head)
	expected := true
	require.Equal(t, expected, res)

	lengthExpected := 4
	require.Equal(t, lengthExpected, NewFastAndSlowPointers().LengthOfCycle(head))
}
