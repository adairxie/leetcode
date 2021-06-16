package isValidBST

import "math"

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, min int64, max int64) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}
