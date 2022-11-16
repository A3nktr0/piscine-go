package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
