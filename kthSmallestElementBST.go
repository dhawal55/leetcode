package main

func kthSmallest(root *TreeNode, k int) int {
	var result []int
	result = inorderTraversal(root, result, k)

	if k > len(result) {
		return 0
	}

	return result[k-1]
}

func inorderTraversal(node *TreeNode, result []int, k int) []int {
	if len(result) < k && node != nil {
		result = inorderTraversal(node.Left, result, k)
		result = append(result, node.Val)
		result = inorderTraversal(node.Right, result, k)
	}

	return result
}

func kthSmallestonRecurse(root *TreeNode, k int) int {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			temp := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			k--
			if k == 0 {
				return temp.Val
			}

			root = temp.Right
		}
	}

	return -1
}
