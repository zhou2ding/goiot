package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var num1, num2 string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		num1 = sc.Text()
		for sc.Scan() {
			num2 = sc.Text()
			break
		}
		fmt.Println(multi(num1, num2))
	}
}

func multi(num1, num2 string) string {
	var total int64
	for i := len(num1) - 1; i >= 0; i-- {
		tmp1 := int64(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			tmp2 := int64(num2[j] - '0')
			left := tmp1 * int64(math.Pow(10, float64(len(num1)-1-i)))
			right := tmp2 * int64(math.Pow(10, float64(len(num2)-1-j)))
			total += left * right
		}
	}
	result := strconv.FormatInt(total, 64)
	return result
}
