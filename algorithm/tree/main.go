package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		t1 := &treeNode{}
		t2 := &treeNode{}
		for i := 0; i < n; i++ {
			insert1(t1, i)
			insert2(t2, i)
		}
		inorder(t1)
		inorder(t2)
	}
}

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func createNode(val int) *treeNode {
	return &treeNode{val: val}
}

func (t *treeNode) insert(n *treeNode, val int) {
	cur := n
	for cur != nil {
		if val > cur.val {
			if cur.right == nil {
				cur.right = createNode(val)
				return
			} else {
				cur = cur.right
			}
		} else {
			if cur.left == nil {
				cur.left = createNode(val)
				return
			} else {
				cur = cur.left
			}
		}
	}
}

func insert1(node *treeNode, val int) *treeNode {
	if node == nil {
		return &treeNode{val: val, left: nil, right: nil}
	}
	if node.val == val {
		return node
	}
	if node.val > val {
		node.left = insert1(node.left, val)
	}
	if node.val < val {
		node.right = insert1(node.right, val)
	}
	return node
}

func insert2(node *treeNode, val int) *treeNode {
	if node == nil {
		return &treeNode{val: val, left: nil, right: nil}
	}
	if node.val == val {
		return node
	}
	cur := node
	prev := node
	for cur != nil {
		prev = cur
		if cur.val > val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	ret := &treeNode{val: val, left: nil, right: nil}
	if prev.val > val {
		prev.left = ret
	} else {
		prev.right = ret
	}
	return node
}

func preorder(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " ")
	preorder(node.left)
	preorder(node.right)
}

func inorder(node *treeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Print(node.val, " ")
	inorder(node.right)
}

func postorder(node *treeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	inorder(node.right)
	fmt.Print(node.val, " ")
}
