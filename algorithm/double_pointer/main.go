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
		arrStr := strings.Split(sc.Text(), " ")
		var arr []int
		for _, v := range arrStr {
			num, _ := strconv.Atoi(v)
			arr = append(arr, num)
		}
		fmt.Println(maxArea(arr))
	}
}
