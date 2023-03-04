package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		s := sc.Text()
		fmt.Println(sort(s))
	}
}

func sort(s string) string {
	others := make([][1]rune, len(s))
	letters := make(map[rune][]bool) // keys转成小写后的字符，value切片是该字符大小写的顺序，true小写，false大写

	for i, c := range s {
		if unicode.IsLower(c) {
			letters[c] = append(letters[c], true)
		} else if unicode.IsUpper(c) {
			letters[c+32] = append(letters[c+32], false)
		} else {
			others[i] = [1]rune{c}
		}
	}

	var sortedLetters string
	for i := 97; i <= 122; i++ {
		ul := letters[rune(i)]
		for _, isLower := range ul {
			if isLower {
				sortedLetters += string(rune(i))
			} else {
				sortedLetters += string(rune(i - 32))
			}
		}
	}

	for idx, c := range others {
		if c[0] != 0 {
			if idx <= len(sortedLetters) {
				sortedLetters = sortedLetters[:idx] + string(c[0]) + sortedLetters[idx:]
			} else {
				sortedLetters += string(c[0])
			}
		}
	}

	return sortedLetters
}
