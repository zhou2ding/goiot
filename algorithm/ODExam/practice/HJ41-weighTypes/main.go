package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n       int
		weights []int
		counts  []int
	)
	for {
		nn, _ := fmt.Scan(&n)
		if nn == 0 {
			break
		} else {
			sc := bufio.NewScanner(os.Stdin)
			sc.Scan()
			line2 := strings.Split(sc.Text(), " ")
			for _, v := range line2 {
				weight, _ := strconv.Atoi(v)
				weights = append(weights, weight)
			}
			sc.Scan()
			line3 := strings.Split(sc.Text(), " ")
			for _, v := range line3 {
				count, _ := strconv.Atoi(v)
				counts = append(counts, count)
			}
			fmt.Println(getWeightTypes(weights, counts, n))
		}
	}
}

func getWeightTypes(weights, counts []int, n int) int {
	m := make(map[int]bool)
	m[0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < counts[i]; j++ {
			var weighted []int
			for k := range m {
				weighted = append(weighted, k)
			}
			for _, weight := range weighted {
				m[weight+weights[i]] = true
			}
		}
	}

	return len(m)
}
