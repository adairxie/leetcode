package dfs

import (
	"leetcode/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) (res []*TreeNode) {
	if root == nil {
		return res
	}

	stack := stack.New()
	stack.Push(root)

	for !stack.Empty() {
		top := stack.Peek()
		res = append(res, top.(*TreeNode))
		if top.(*TreeNode).Right != nil {
			res = append(res, top.(*TreeNode).Right)
		}
		if top.(*TreeNode).Left != nil {
			res = append(res, top.(*TreeNode).Left)
		}
	}
	return res
}
