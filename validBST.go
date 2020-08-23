func isValidBST(root *TreeNode) bool {
	return traverse(root, nil, nil)
}

func traverse(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && !(node.Val > *min) {
		return false
	}

	if max != nil && !(node.Val < *max) {
		return false
	}

	return traverse(node.Left, min, &node.Val) && traverse(node.Right, &node.Val, max)
}

func isValidBSTNonrecurse(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode
	var prev int
	var prevSet bool

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if prevSet && root.Val <= prev {
				return false
			}

			prev = root.Val
			prevSet = true
			root = root.Right
		}
	}

	return true
}