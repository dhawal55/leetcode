package main

func inorderTraversalRecursive(root *TreeNode) []int {
	var result []int

	if root == nil {
		return nil
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result = append(result, root.Val)
			root = root.Right
		}
	}

	return result
}
