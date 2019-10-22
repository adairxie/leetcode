package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(1+maxDepth(root.Left), 1+maxDepth(root.Right))
}
