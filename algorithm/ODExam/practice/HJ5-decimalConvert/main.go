package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func main() {
	var a string
	fmt.Scanln(&a)
	fmt.Println(hexToDec(a))
}

func hexToDec(s string) string {
	num := 0
	top := len(s) - 2 - 1
	for i := 2; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			num += (int(s[i]) - 55) * int(math.Pow(16, float64(top)))
		} else {
			num += (int(s[i] - '0')) * int(math.Pow(16, float64(top)))
		}
		top--
	}
	return strconv.Itoa(num)
}
