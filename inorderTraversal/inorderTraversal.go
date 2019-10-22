package inorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	var stack []*TreeNode = []*TreeNode{}
	var node *TreeNode = root
	var res []int = []int{}

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) > 0 {
			last := stack[len(stack)-1]
			node = last.Right
			stack = stack[:len(stack)-1]
			res = append(res, last.Val)
		}
	}

	return res
}
