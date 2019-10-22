package maxDepth

import "testing"

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 2}
	l1 := &TreeNode{Val: 9}
	r1 := &TreeNode{Val: 20}
	root.Left = l1
	root.Right = r1

	rl2 := &TreeNode{Val: 15}
	rr2 := &TreeNode{Val: 7}
	r1.Left = rl2
	r1.Right = rr2

	depth := maxDepth(root)
	t.Logf("tree max depth: %d", depth)
}
