package isSameTree

import "testing"

func TestIsSameTree(t *testing.T) {

	p := TreeNode{Val:1}
	p2 := TreeNode{Val:2}
	p3 := TreeNode{Val:3}
	p.Left = &p2
	p.Right = &p3

	q := TreeNode{Val:1}
	q2 := TreeNode{Val:2}
	q3 := TreeNode{Val:3}
	q4 := TreeNode{Val:4}
	q.Left = &q2
	q.Right = &q3
	q3.Right = & q4
	res := isSameTree(&p, &q)
	t.Logf("res:%t\n", res)
}