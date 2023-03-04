package main

import (
	"fmt"
	"sort"
	"unicode"
)

func main() {
	var (
		n int
		s string
	)
	for {
		nn, _ := fmt.Scan(&n)
		if nn == 0 {
			break
		} else {
			for i := 0; i < n; i++ {
				fmt.Scan(&s)
				fmt.Printf("%d\n", getMaxBeauty(s))
			}
		}
	}
}

func getMaxBeauty(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		unicode.ToLower(c)
		m[c]++
	}

	sortedKeys := make([]int, 0)
	for _, cnt := range m {
		sortedKeys = append(sortedKeys, cnt)
	}
	sort.Ints(sortedKeys)

	maxBeauty := 0
	beauty := 26
	for i := len(sortedKeys) - 1; i >= 0; i-- {
		maxBeauty += sortedKeys[i] * beauty
		beauty--
	}
	return maxBeauty
}
