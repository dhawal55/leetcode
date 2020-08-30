func copyRandomList(head *Node) *Node {
	listMapping := make(map[*Node]*Node)

	var newCurrent, prev *Node
	current := head
	for current != nil {
		if node, ok := listMapping[current]; ok {
			newCurrent = node
		} else {
			newCurrent = &Node{}
			listMapping[current] = newCurrent
		}
		newCurrent.Val = current.Val

		if current.Random != nil {
			if node, ok := listMapping[current.Random]; ok {
				newCurrent.Random = node
			} else {
				newCurrent.Random = &Node{}
				listMapping[current.Random] = newCurrent.Random
			}
		}

		if prev != nil {
			prev.Next = newCurrent
		}

		prev = newCurrent
		current = current.Next
	}

	return listMapping[head]
}