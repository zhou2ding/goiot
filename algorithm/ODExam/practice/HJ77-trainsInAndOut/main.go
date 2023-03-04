package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	n := 0
	arr := []int{}
	a := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	results := trainsOut(arr)
	sort.Strings(results)
	for _, result := range results {
		for _, idx := range result {
			fmt.Printf("%c ", idx)
		}
		fmt.Println()
	}
}

func trainsOut(trains []int) []string {
	res := make([]string, 0)
	path := ""
	stack := make([]int, 0)

	var dfs func(int, int)
	// in、out是进、出栈的次数
	dfs = func(in, out int) {
		if out == len(trains) {
			res = append(res, path)
		}
		if len(stack) != 0 {
			outTrain := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			path += strconv.Itoa(outTrain)
			dfs(in, out+1)
			stack = append(stack, outTrain)
			path = path[:len(path)-1]
		}
		if in < len(trains) {
			stack = append(stack, trains[in])
			dfs(in+1, out)
			stack = stack[:len(stack)-1]
		}
	}
	dfs(0, 0)

	return res
}
