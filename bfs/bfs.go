package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		var current []int
		var nqueue []*TreeNode
		for i := 0; i < len(queue); i++ {
			current = append(current, queue[i].Val)
			if queue[i].Left != nil {
				nqueue = append(nqueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nqueue = append(nqueue, queue[i].Right)
			}
		}
		queue = nqueue
		result = append(result, current)
	}
	return result
}
