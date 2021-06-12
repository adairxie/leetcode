package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	var first *ListNode = head
	for i := n; i > 0; i-- {
		first = first.Next
	}

	if first == nil {
		head = head.Next
		return head
	}

	pre := head
	for first.Next != nil {
		pre = pre.Next
		first = first.Next
	}

	pre.Next = pre.Next.Next

	return head
}

func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	var result = &ListNode{}
	result.Next = head
	cur := result

	var pre *ListNode
	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}

	pre.Next = pre.Next.Next

	return result.Next
}
