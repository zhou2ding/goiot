package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输出第一个数数字：")
		var a int
		for sc.Scan() {
			a, _ = strconv.Atoi(sc.Text())
			break
		}
		fmt.Print("请输出第二个数数字：")
		var b int
		for sc.Scan() {
			b, _ = strconv.Atoi(sc.Text())
			break
		}
		fmt.Printf("%v 和 %v 的和为：%v\n", a, b, a+b)
	}
}
