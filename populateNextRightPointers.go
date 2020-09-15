package main

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root, nil}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			if len(queue) > 0 {
				node.Next = queue[0]
			} else {
				node.Next = nil
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		} else {
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		}
	}

	return root
}
