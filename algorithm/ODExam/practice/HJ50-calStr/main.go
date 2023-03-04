package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scanln(&s)
	s = strings.ReplaceAll(s, "{", "(")
	s = strings.ReplaceAll(s, "[", "(")
	s = strings.ReplaceAll(s, "}", ")")
	s = strings.ReplaceAll(s, "]", ")")
	fmt.Println(cal(s))
}

func cal(s string) int {
	num := 0
	stack := make([]int, 0)
	opt := '+'
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '(' {
			cnt := 1
			j := i
			for cnt > 0 {
				j++
				if s[j] == '(' {
					cnt++
				}
				if s[j] == ')' {
					cnt--
				}
			}
			num = cal(s[i+1 : j])
			i = j
		}
		if !unicode.IsDigit(rune(s[i])) || i == len(s)-1 {
			switch opt {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			opt = rune(s[i])
			num = 0
		}
	}
	sum := 0
	for _, n := range stack {
		sum += n
	}
	return sum
}
