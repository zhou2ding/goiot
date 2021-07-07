package main

import (
	"fmt"
)

func main() {
	num1 := []int{1,3,5,7}
	num2 := []int{2,4,6,8}
	list1 := makeList(num1)
	list2 := makeList(num2)
	printList(list1)
	printList(list2)
	ret := merge(list1,list2)
	printList(ret)
}

func printList(ret *List) {
	for ret.Next != nil {
		fmt.Print(ret.Val," ")
		ret = ret.Next
	}
	fmt.Println(ret.Val)
}

func makeList(nums []int) *List {
	if len(nums) ==  0 {
		return nil
	}
	res := &List{
		Val:nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &List{Val:nums[i],}
		temp = temp.Next
	}

	return  res
}

type List struct {
	Val  int
	Next *List
}

func merge(List1, List2 *List) (ret *List) {
	if List1 == nil {
		return List2
	}
	if List2 == nil {
		return List1
	}
	if List1.Val <= List2.Val {
		List1.Next = merge(List1.Next,List2)
		return List1
	} else {
		List2.Next = merge(List2.Next,List1)
		return List2
	}
}
