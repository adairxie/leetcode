package removeNthFromEnd

import "testing"

func printList(head *ListNode) {
	for head != nil {
		print(head.Val)
		head = head.Next
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	
	e1 := &ListNode{
		Val: 2,
	}
	
	e2 := &ListNode{
		Val: 3,
	}

	e3 := &ListNode{
		Val: 4,
	}

	e4 := &ListNode{
		Val: 5,
	}
	

	head.Next = e1
	e1.Next = e2
	e2.Next = e3
	e3.Next = e4

	res := removeNthFromEnd(head, 5)
	printList(res)
}
