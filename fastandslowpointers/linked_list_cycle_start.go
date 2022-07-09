package fastandslowpointers

// FindCycleStart N1->N2->N3->N4->N5->N6--->to N3
// output = 3
func (fastAndSlowPointers) FindCycleStart(head *ListNode) int {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			length := findCycleStartCalculateLength(slow)
			return findCycleStartCalculateStart(head, length).Val
		}
	}

	return 0
}

func findCycleStartCalculateLength(slow *ListNode) int {
	current := slow
	length := 0

	for {
		current = current.Next
		length++
		if current == slow {
			break
		}
	}
	return length
}

// findCycleStartCalculateStart move pointer two length times ahead
// move by one each pointers until the meet, then that will be the start
func findCycleStartCalculateStart(slow *ListNode, length int) *ListNode {
	pointer1 := slow
	pointer2 := slow

	for i := 0; i < length; i++ {
		pointer2 = pointer2.Next
	}

	for pointer1 != pointer2 {
		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
	}

	return pointer1
}
