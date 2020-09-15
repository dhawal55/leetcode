package main

type NodeWrapper struct {
	Node *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result := &NodeWrapper{}
	inorderTraverse(root, p, q, result)
	return result.Node
}

func inorderTraverse(node, p, q *TreeNode, result *NodeWrapper) bool {
	if node == nil || result.Node != nil {
		return false
	}

	var left, right, mid int
	if inorderTraverse(node.Left, p, q, result) {
		left = 1
	}

	if inorderTraverse(node.Right, p, q, result) {
		right = 1
	}

	if node == p || node == q {
		mid = 1
	}

	if mid+left+right >= 2 {
		result.Node = node
	}

	return left+right+mid > 0
}
