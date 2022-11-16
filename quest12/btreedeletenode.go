package piscine

func FreeNode(node *TreeNode) *TreeNode {
	n := node
	node = nil
	return n
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			temp := root.Right
			FreeNode(root)
			return temp
		} else if root.Right == nil {
			temp := root.Left
			FreeNode(root)
			return temp
		}
		temp := BTreeMin(root.Right)
		root.Data = temp.Data
		root.Right = BTreeDeleteNode(root.Right, temp)
	}
	return root
}
