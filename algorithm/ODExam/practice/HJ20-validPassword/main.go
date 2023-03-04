package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	for {
		n, _ := fmt.Scan(&s)
		if n == 0 {
			break
		} else {
			if isValid1(s) && isValid2(s) && isValid3(s) {
				fmt.Println("OK")
			} else {
				fmt.Println("NG")
			}
		}
	}
}

func isValid1(s string) bool {
	return len(s) > 8
}

func isValid2(s string) bool {
	var (
		upper int
		lower int
		digit int
		other int
	)
	for _, c := range s {
		if string(c) != " " && string(c) != "\r" && string(c) != "\n" && string(c) != "\r\n" {
			if unicode.IsUpper(c) {
				upper = 1
			} else if unicode.IsLower(c) {
				lower = 1
			} else if unicode.IsDigit(c) {
				digit = 1
			} else {
				other = 1
			}
		}
	}
	if upper+lower+digit+other < 3 {
		return false
	}
	return true
}

func isValid3(s string) bool {
	for i := 0; i < len(s)-5; i++ {
		for j := i + 3; j < len(s)-2; j++ {
			if s[j] == s[i] && s[j+1] == s[i+1] && s[j+2] == s[i+2] {
				return false
			}
		}
	}
	return true
}
