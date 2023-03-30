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
		quickSort(arr, 0, len(arr)-1)
		fmt.Println(arr)
	}
}

func quickSort(arr []int, left, right int) {
	if left < right {
		key := arr[(left+right)/2]
		i, j := left, right
		for i < j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		quickSort(arr, left, i-1)
		quickSort(arr, j+1, right)
	}
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
