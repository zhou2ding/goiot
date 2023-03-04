package main

import (
	"bufio"
	"fmt"
	"os"
)

// A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
// n i h a o B C D E F G J K L M P Q R S T U V W X Y Z
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	key := sc.Text()
	sc.Scan()
	msg := sc.Text()
	fmt.Println(encrypt(key, msg))
}

func encrypt(key, msg string) string {
	alphabet := make([]rune, 0)
	for i := 97; i <= 122; i++ {
		alphabet = append(alphabet, rune(i))
	}

	used := make(map[rune]bool)
	head := make([]rune, 0)
	for _, c := range key {
		if !used[c] {
			head = append(head, c)
			used[c] = true
		}
	}

	keyAlphabet := make([]rune, len(head))
	copy(keyAlphabet, head)
	for _, c := range alphabet {
		if !used[c] {
			keyAlphabet = append(keyAlphabet, c)
		}
	}

	var res string
	for _, c := range msg {
		for i, c2 := range alphabet {
			if c2 == c {
				res += string(keyAlphabet[i])
			}
		}
	}
	return res
}
