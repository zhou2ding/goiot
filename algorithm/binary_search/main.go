package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input := strings.Split(sc.Text(), " ")
		arr := make([]int, 0, len(input))
		for _, v := range input {
			num, _ := strconv.Atoi(v)
			arr = append(arr, num)
		}
		var target int
		for sc.Scan() {
			target, _ = strconv.Atoi(sc.Text())
			break
		}
		fmt.Println(binarySearch(arr, target))
		fmt.Println(binarySearch2(arr, target, 0, len(arr)-1))
	}
}

// 非递归方式
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 递归方式
func binarySearch2(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := (right + left) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearch2(arr, target, mid+1, right)
	} else {
		return binarySearch2(arr, target, left, mid-1)
	}
}
