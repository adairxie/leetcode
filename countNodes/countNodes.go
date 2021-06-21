package countNodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if left == right {
		return countNodes(root.Right) + (1 << left)
	} else {
		return countNodes(root.Left) + (1 << right)
	}
}

func countLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		level++
		root = root.Left
	}
	return level
}
