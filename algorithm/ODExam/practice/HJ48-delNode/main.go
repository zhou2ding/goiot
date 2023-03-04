package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			var arr []int
			input := strings.Split(sc.Text(), " ")
			for _, v := range input {
				n, _ := strconv.Atoi(v)
				arr = append(arr, n)
			}
			l := newNode(arr[1])
			for i := 3; i < len(arr)-1; i += 2 {
				l.insert(arr[i-1], arr[i])
			}
			l = l.del(arr[len(arr)-1])
			for l != nil {
				fmt.Printf("%d ", l.val)
				l = l.next
			}
			fmt.Println()
			break
		}
		break
	}
}

type listNode struct {
	val  int
	next *listNode
}

func newNode(val int) *listNode {
	return &listNode{val: val}
}

func (l *listNode) insert(val, pos int) *listNode {
	temp := l
	for {
		if temp.val != pos {
			temp = temp.next
		} else {
			node := newNode(val)
			if temp.next == nil {
				temp.next = node
			} else {
				node.next = temp.next
				temp.next = node
			}
			break
		}
	}
	return l
}

func (l *listNode) del(val int) *listNode {
	temp := l
	if temp.val == val {

		return l.next
	}
	for {
		if temp.next.val == val {
			temp.next = temp.next.next
			break
		} else {
			temp = temp.next
		}
	}
	return l
}
