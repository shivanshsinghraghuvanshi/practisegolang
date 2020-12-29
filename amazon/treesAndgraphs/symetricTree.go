package treesAndgraphs

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return (t1.Value == t2.Value) && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}
