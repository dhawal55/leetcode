func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var current []int
	var result [][]int
	leftToRight := true
	queue := []*TreeNode{root, nil}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			if leftToRight {
				current = append(current, node.Val)
			} else {
				current = append([]int{node.Val}, current...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		} else {
			result = append(result, current)
			leftToRight = !leftToRight

			current = nil
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		}
	}

	return result
}

func zigzagLevelOrderRecurse(root *TreeNode) [][]int {
	var res [][]int
	res = getLevelOrder(root, 0, res)
	return res
}

func getLevelOrder(node *TreeNode, level int, res [][]int) [][]int {
	if node == nil {
		return res
	}

	if level >= len(res) {
		res = append(res, []int{node.Val})
	} else {
		levelArr := res[level]
		if level%2 == 0 {
			levelArr = append(levelArr, node.Val)
		} else {
			levelArr = append([]int{node.Val}, levelArr...)
		}
		res[level] = levelArr
	}

	res = getLevelOrder(node.Left, level+1, res)
	res = getLevelOrder(node.Right, level+1, res)

	return res
}