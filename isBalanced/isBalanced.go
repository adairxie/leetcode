package isBalanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	leftH := maxDepth(root.Left)
	rightH := maxDepth(root.Right)

	if abs(leftH, rightH) > 1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftD := maxDepth(root.Left)
	rightD := maxDepth(root.Right)

	return max(leftD, rightD) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
