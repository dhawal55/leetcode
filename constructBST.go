package main

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
}

// Time Complexity: O(N ^2)
// Space Complexity: O(1)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return node
	}

	i := 0
	for i < len(inorder) && inorder[i] != preorder[0] {
		i++
	}

	node.Left = buildTree(preorder[1:i+1], inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return node
}

// Time complexity : O(N). Let's compute the solution with the help of master theorem
// 	 T(N) = 2 * T(N/2) + T(N)
// Space complexity : O(N), since we store the entire tree.
func buildTreeWithMap(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	inorderMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}

	preIndex := new(int)
	return buildTreeRecurse(preorder, inorder, inorderMap, 0, len(preorder), preIndex)
}

func buildTreeRecurse(preorder []int, inorder []int, m map[int]int, left, right int, preIndex *int) *TreeNode {
	if left == right {
		return nil
	}

	node := &TreeNode{
		Val: preorder[*preIndex],
	}
	*preIndex++

	index := m[node.Val]
	node.Left = buildTreeRecurse(preorder, inorder, m, left, index, preIndex)
	node.Right = buildTreeRecurse(preorder, inorder, m, index+1, right, preIndex)

	return node
}
