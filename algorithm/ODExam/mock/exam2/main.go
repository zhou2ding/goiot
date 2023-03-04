package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var a int
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scanln(&a)
		arr = append(arr, a)
	}
	//res := distinctAndSortWithStd(arr)
	res := distinctAndSort(arr)
	for _, v := range res {
		fmt.Println(v)
	}
}

func distinctAndSortWithStd(arr []int) []int {
	m := make(map[int]bool)
	for _, v := range arr {
		m[v] = true
	}

	res := make([]int, 0)
	for k := range m {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

func distinctAndSort(arr []int) []int {
	sort.Ints(arr)
	i, j := 0, 1
	for j < len(arr) {
		if arr[i] == arr[j] {
			j++
		} else {
			i++
			arr[i] = arr[j]
		}
	}
	return arr[:i+1]
}
