package inorderTraversal

import "testing"

func TestInorderTraversal(t *testing.T) {
	v1 := &TreeNode{
		Val: 1,
	}

	v2 := &TreeNode{
		Val: 2,
	}

	v3 := &TreeNode{
		Val: 3,
	}

	v1.Right = v2
	v2.Left = v3

	res := inorderTraversal(v1)

	t.Logf("res:%v", res)
}
