package main

import (
	"encoding/json"
	"fmt"
)

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getPerson() (p *person) {
	data := `{"name":"zhangsan","age":18}`
	err := json.Unmarshal([]byte(data), &p)
	if err != nil {
		fmt.Println("unmarshal failed, error:", err)
		return
	}
	return
}
func main() {
	// fmt.Println(f1()) // 5
	// fmt.Println(f2()) // 6
	// fmt.Println(f3()) // 5
	// fmt.Println(f4()) // 5
	fmt.Printf("%#v\n", getPerson())
}
