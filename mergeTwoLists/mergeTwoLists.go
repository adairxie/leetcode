package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var cur *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		head, cur, l1 = l1, l1, l1.Next
	} else {
		head, cur, l2 = l2, l2, l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next, cur, l1 = l1, l1, l1.Next
		} else {
			cur.Next, cur, l2 = l2, l2, l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return head
}

func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{}
	cur := result

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return result.Next
}
