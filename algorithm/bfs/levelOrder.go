//剑指 Offer 32 - I. 从上到下打印二叉树
//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var ret []int
	queue := []*TreeNode{root}
	if root == nil {
		return ret
	}
	var node *TreeNode
	for {
		if len(queue) == 0 {
			break
		}
		node = queue[0]
		queue = queue[1:]
		ret = append(ret,node.Val)
		if node.Left != nil {
			queue = append(queue,node.Left)
		}
		if node.Right != nil {
			queue = append(queue,node.Right)
		}
	}
	return ret
}

