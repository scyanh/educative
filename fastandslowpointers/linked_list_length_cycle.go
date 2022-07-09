package fastandslowpointers

// LengthOfCycle finds the length of the cycle
// if there is no cycle, return 0
// complexity: O(n)
// space: O(1)
func (fastAndSlowPointers) LengthOfCycle(head *ListNode) int {
	if head == nil {
		return 0
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return calculateCycleLength(slow)
		}
	}

	return 0
}

func calculateCycleLength(slow *ListNode) int {
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
