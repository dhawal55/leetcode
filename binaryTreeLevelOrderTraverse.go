/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root, nil}
	var current []int
	var result [][]int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			current = append(current, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		} else {
			result = append(result, current)
			current = nil

			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		}
	}

	return result
}

func levelOrderRecurse(root *TreeNode) [][]int {
	return getLevelOrder(root, nil)
}

func getLevelOrder(node *TreeNode, result [][]int) [][]int {
	if root == nil {
		return nil
	}

	result = append(result, []int{node.Val})
	result = getLevelOrder(node.Left, result)
	result = getLevelOrder(node.Right, result)

	return result
}