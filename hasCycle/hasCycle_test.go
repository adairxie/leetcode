package hasCycle

import "testing"

func TestHasCycle(t *testing.T) {
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
	e4.Next = e2

	cycle := hasCycle(head)
	t.Logf("hasCycle: %t\n", cycle)
}
