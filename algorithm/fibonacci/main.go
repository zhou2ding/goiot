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
		fmt.Println(genFibArr(n))
	}
}

func genFibArr(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, fib(i))
	}
	return arr
}

func fib(n int) int {
	if n < 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
