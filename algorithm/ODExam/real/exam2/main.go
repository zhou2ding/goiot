package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	sc := bufio.NewScanner(os.Stdin)
	inputs := make([]string, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		inputs = append(inputs, sc.Text())
	}

	services := make([][]int, len(inputs))
	for i := range services {
		services[i] = make([]int, 0)
	}

	for i, v := range inputs {
		line := strings.Split(v, " ")
		for _, s := range line {
			num, _ := strconv.Atoi(s)
			services[i] = append(services[i], num)
		}
	}
	var k int
	fmt.Scanln(&k)
	fmt.Println(getMaxWaitTime(services, k-1))
}

func getMaxWaitTime(services [][]int, k int) int {
	waitTime := 0
	for i := 0; i < len(services[k]); i++ {
		if i != k && services[k][i] != 0 {
			waitTime = max(waitTime, getMaxWaitTime(services, i))
		}
	}
	return services[k][k] + waitTime
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
