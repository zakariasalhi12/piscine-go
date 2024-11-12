package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	for current.Right != nil {
		current = current.Right
	}

	return current
}
