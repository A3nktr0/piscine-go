package piscine

func ChangeLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 0 {
		f(root.Data)
	} else {
		ChangeLevel(root.Left, level-1, f)
		ChangeLevel(root.Right, level-1, f)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	for i := 0; i < BTreeLevelCount(root); i++ {
		ChangeLevel(root, i, f)
	}
}
