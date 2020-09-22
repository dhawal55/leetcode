/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Time Complexity: O(n)
// Space Complexity: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	evenHead := head.Next

	odd := head
	even := head.Next

	for odd != nil && even != nil {
		odd.Next = even.Next
		if odd.Next != nil {
			even.Next = odd.Next.Next
			odd = odd.Next
		}

		even = even.Next
	}

	odd.Next = evenHead
	return head
}