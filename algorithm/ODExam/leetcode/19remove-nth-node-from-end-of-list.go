package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	rm := findNthFromEnd(dummy, n+1)
	rm.Next = rm.Next.Next
	return dummy.Next
}

func findNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
