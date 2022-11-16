package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	lh := BTreeLevelCount(root.Left)
	rh := BTreeLevelCount(root.Right)
	if lh >= rh {
		return lh + 1
	} else {
		return rh + 1
	}
}
