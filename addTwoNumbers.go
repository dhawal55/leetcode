// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Example:
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 4,
		},
	}

	result := addTwoNumbers(l1, l2)

	var out int
	var tensUnit int
	for result != nil {
		tensUnit = tensUnit * 10
		if tensUnit == 0 {
			tensUnit = 1
		}

		out = (tensUnit * result.Val) + out
		result = result.Next
	}

	fmt.Println(out)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	head := &ListNode{}
	current := head
	var carry int
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		temp := &ListNode{
			Val: sum % 10,
		}

		current.Next = temp
		current = temp
	}

	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}

	return head.Next
}
