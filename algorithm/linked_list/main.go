package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		l := linkedList{}
		l.createTail(n)
	}
}

type linkedList struct {
	val  int
	next *linkedList
}

// 尾插法
func (l *linkedList) createTail(n int) *linkedList {
	temp := l
	for i := 0; i < n; i++ {
		node := &linkedList{val: i}
		temp.next = node
		temp = node
	}
	return l
}

// 头插法（总是插到首节点的下一个）
func (l *linkedList) createHead(n int) *linkedList {
	for i := 0; i < n; i++ {
		node := &linkedList{val: i}
		node.next = l.next
		l.next = node
	}
	return l
}

func (l *linkedList) insert(pos, val int) *linkedList {
	temp := l
	for i := 0; i < pos-1; i++ {
		temp = temp.next
	}
	node := &linkedList{val: val}
	node.next = temp.next
	temp.next = node
	return l
}

func (l *linkedList) del(pos int) *linkedList {
	temp := l
	for i := 0; i < pos-1; i++ {
		temp = temp.next
	}
	delNode := temp.next
	temp.next = delNode.next
	delNode = nil
	return l
}
