package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	head = swapPairs(head)

	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{Next: head}
	current := head.Next
	head = current

	for current != nil {
		temp := current.Next
		current.Next = prev.Next
		prev.Next.Next = temp
		prev.Next = current

		if temp != nil {
			prev = current.Next
			current = temp.Next
		} else {
			current = nil
		}

	}

	return head
}
