package fastandslowpointers

// HasCycle determine if the LinkedList has a cycle
// Return true if the LinkedList has a cycle, false otherwise
// complexity: O(n)
// space: O(1)
func (fastAndSlowPointers) HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
