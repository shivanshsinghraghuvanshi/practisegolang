package treesAndgraphs

import "math"

var max = math.MinInt32

func maxPathSum(root *TreeNode) int {
	maxGain(root)
	return max
}

func maxGain(n *TreeNode) int {
	if n == nil {
		return 0
	}

	l := maximum(maxGain(n.Left), 0)
	r := maximum(maxGain(n.Right), 0)

	cv := n.Value + l + r

	max = maximum(max, cv)

	return n.Value + maximum(l, r)
}

func maximum(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
