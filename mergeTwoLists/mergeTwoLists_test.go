package mergeTwoLists

import (
	"fmt"
	"testing"
)

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
	}

	l12 := &ListNode{
		Val: 2,
	}

	l13 := &ListNode{
		Val: 4,
	}

	l1.Next = l12
	l12.Next = l13

	l2 := &ListNode{
		Val: 1,
	}

	l21 := &ListNode{
		Val: 3,
	}

	l22 := &ListNode{
		Val: 4,
	}

	l2.Next = l21
	l21.Next = l22

	res := mergeTwoListsV2(l1, l2)
	printList(res)
}
