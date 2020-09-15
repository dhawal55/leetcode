package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	current := head

	for i := 0; i < n; i++ {
		current = current.Next
		if current == nil {
			return head.Next
		}
	}

	for current.Next != nil {
		current = current.Next
		prev = prev.Next
	}

	prev.Next = prev.Next.Next

	return head
}
