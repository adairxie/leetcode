package addTwoNumbers
 
type ListNode struct {
     Val int
     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var header, cur *ListNode
	var sum, incr int

	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + incr
		if sum >= 10 {
			incr = 1
		}else {
			incr = 0
		}

		v := sum - 10 * incr
		if cur == nil {
			cur = new(ListNode)
			if header == nil {
				header = cur
			}
		} else {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
		cur.Val = v
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if cur == nil {
			cur = new(ListNode)
			if header == nil {
				header = cur
			}
		}else {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
		sum = l1.Val + incr
		if sum >= 10 {
			incr = 1
		}else {
			incr = 0
		}

		v := sum - 10 * incr
		cur.Val = v
		l1 = l1.Next
	}

	for l2 != nil {
		if cur == nil {
			cur = new(ListNode)
			if header == nil {
				header = cur
			}
		} else {
			cur.Next = new(ListNode)
			cur = cur.Next
		}

		sum = l2.Val + incr
		if sum >= 10 {
			incr = 1
		}else{
			incr = 0 
		}
		v := sum - 10 * incr
		cur.Val = v
		l2 = l2.Next
	}

	if incr == 1 {
		cur.Next = new(ListNode)
		cur.Next.Val = incr
	}

	return header
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	rs := &ListNode{
		Val: 0,
		Next: nil,
	}
	
	cur := rs
	var incr int

	for l1 != nil || l2 != nil || incr > 0 {
		var t1, t2 int
		if l1 != nil {
			t1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			t2 = l2.Val
			l2 = l2.Next
		}

		v := t1 + t2 + incr
		if v >= 10 {
			v = v % 10
			incr = 1
		}else {
			incr = 0
		}

		tmp := &ListNode{
			Val: v,
			Next: nil,
		}

		cur.Next = tmp
		cur = cur.Next
	}

	return rs.Next
}
// declare variable with var is more faster than :=
func AddTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	var rs = &ListNode{
		Val: 0,
		Next: nil,
	}
	
	var cur = rs
	var incr int

	for l1 != nil || l2 != nil || incr > 0 {
		var t1, t2 int
		if l1 != nil {
			t1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			t2 = l2.Val
			l2 = l2.Next
		}

		var v = t1 + t2 + incr
		if v >= 10 {
			v = v % 10
			incr = 1
		}else {
			incr = 0
		}

		var tmp = &ListNode{
			Val: v,
			Next: nil,
		}

		cur.Next = tmp
		cur = cur.Next
	}

	return rs.Next
}