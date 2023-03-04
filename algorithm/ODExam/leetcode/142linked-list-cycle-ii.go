package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for {
		slow.Val = -100001
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			return nil
		}
		if slow.Val == -100001 {
			return slow
		}
		if fast.Val == -100001 {
			return fast
		}
		if fast.Next.Val == -100001 {
			return fast.Next
		}
	}
}
