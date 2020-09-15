package main

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
	var levelArr []int
	var result [][]int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			levelArr = append(levelArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		} else {
			result = append(result, levelArr)
			levelArr = nil

			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		}
	}

	return result
}

func levelOrderRecurse(root *TreeNode) [][]int {
	return getLevelOrder(root, 0, nil)
}

func getLevelOrder(node *TreeNode, level int, result [][]int) [][]int {
	if root == nil {
		return nil
	}

	if level >= len(result) {
		result = append(result, []int{node.Val})
	} else {
		levelArr := result[level]
		levelArr = append(levelArr, node.Val)
		result[level] = levelArr
	}

	result = getLevelOrder(node.Left, level+1, result)
	result = getLevelOrder(node.Right, level+1, result)

	return result
}
