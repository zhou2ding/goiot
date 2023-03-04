package leetcode

func partition(head *ListNode, x int) *ListNode {
	p := head
	dummy1 := &ListNode{Val: -1}
	dummy2 := &ListNode{Val: -1}
	p1, p2 := dummy1, dummy2
	for {
		if p == nil {
			break
		}
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
