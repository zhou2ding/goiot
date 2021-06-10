package main

import "fmt"

var keyboard = map[byte]string{
	50: "abc",
	51: "def",
	52: "ghi",
	53: "jkl",
	54: "mno",
	55: "pqrs",
	56: "tuv",
	57: "wxyz",
}
var ret []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ret = []string{}	// 全局变量，leetcode，要手动重新初始化一下
	dfs(digits, 0, "")
	return ret
}

func dfs(digits string, index int, combination string) {
	if index == len(digits) {
		ret = append(ret, combination)
		return
	}
	digit := digits[index]
	letters := keyboard[digit]
	letterLen := len(letters)
	for i := 0; i < letterLen; i++ {
		tmp := combination + string(letters[i])
		dfs(digits,index+1,tmp)
	}
}

func main() {
	fmt.Println(letterCombinations("2"))
}
