package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	for head != nil && fast != nil && fast.Next != nil {
		if head == fast {
			return true
		}
		fast = fast.Next.Next
		head = head.Next
	}

	return false
}
