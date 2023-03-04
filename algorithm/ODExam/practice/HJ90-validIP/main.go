package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scanln(&s)
	if isValidIP(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isValidIP(s string) bool {
	arr := strings.Split(s, ".")
	if len(arr) != 4 {
		return false
	}
	for _, n := range arr {
		if n == "" {
			return false
		}
		for _, c := range n {
			if !unicode.IsDigit(c) {
				return false
			}
		}
		if n[0] == '0' && len(n) > 1 {
			return false
		}
		num, err := strconv.Atoi(n)
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}
