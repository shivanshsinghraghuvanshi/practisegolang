package treesAndgraphs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(n, min, max *TreeNode) bool {
	if n == nil {
		return true
	}
	if (min != nil && n.Value < min.Value) || (max != nil && n.Value > max.Value) {
		return false
	}
	return validate(n.Left, min, n) && validate(n.Right, n, max)
}
