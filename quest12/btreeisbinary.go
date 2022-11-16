package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil || root.Left == nil || root.Right == nil {
		return true
	}
	if root.Left.Data > root.Data || root.Right.Data < root.Data {
		return false
	}
	if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	}
	return true
}
