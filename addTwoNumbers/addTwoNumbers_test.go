package addTwoNumbers

import (
	"fmt"
)

func ExampleAddTwoNumbers2() {
	l11 := &ListNode{
		Val: 2,
	}
	/*
	l12 := &ListNode{
		Val: 4,
	}
	
	l13 := &ListNode{
		Val: 3,
	}
	l14 := &ListNode{
		Val: 1,
	}*/
	
	//l11.Next = l12
	//l12.Next = l13
	//l13.Next = l14

	l21 := &ListNode{
		Val: 7, 
	}
	/*
	l22 := &ListNode{
		Val: 6,
	}
	l23 := &ListNode {
		Val: 4,
	}
	l21.Next = l22
	l22.Next = l23
	*/

	res := AddTwoNumbers2(l11, l21)
	for res != nil {
		if res.Next != nil {
			fmt.Printf("%d->", res.Val)
		}else{
			fmt.Printf("%d", res.Val)
		}
		res = res.Next
	}

	// Output:
	// 9
}