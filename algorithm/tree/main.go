package main

import "fmt"

func main() {
	t1 := &treeNode{val: 4}
	insert1(t1, 2)
	insert1(t1, 5)
	insert1(t1, 1)
	insert1(t1, 3)
	inorder(t1)

	t2 := &treeNode{val: 4}
	insert2(t2, 2)
	insert2(t2, 5)
	insert2(t2, 1)
	insert2(t2, 3)
	inorder(t2)
}

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
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
	inorder(node.left)
	inorder(node.right)
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
