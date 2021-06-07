// 判断一个链表中是否含有闭环
// 解法：快慢指针，快指针一次走两步，慢指针一次走一步。如果两个指针能相遇，就有闭环

package main

type linkList struct {
	val  int
	next *linkList
}