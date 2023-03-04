package leetcode

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
}
