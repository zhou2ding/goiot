package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{Val: -10001}
	for _, list := range lists {
		mergeTwoLists2(dummy, list)
	}
	return dummy.Next
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	dummy := &ListNode{Val: -1}
	p := dummy
	for {
		if p1 == nil || p2 == nil {
			break
		}
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
