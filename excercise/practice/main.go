package main

import (
	"context"
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
	fmt.Printf("%#v", p)
	return
}
func main() {
	// fmt.Println(f1()) // 5
	// fmt.Println(f2()) // 6
	// fmt.Println(f3()) // 5
	// fmt.Println(f4()) // 5
	// getPerson()
	// var p1 = &person{}
	// var p2 = new(person)
	// var p3 = make([]*person, 0, 10)
	// var tmp = &person{}
	// p3 = append(p3, tmp)
	// data1 := `{"name":"lisi","age":90}`
	// data2 := `{"name":"wangwu","age":12}`
	// data3 := `{"name":"zhangliu","age":66}`

	// json.Unmarshal([]byte(data1), p1)
	// json.Unmarshal([]byte(data2), p2)
	// json.Unmarshal([]byte(data3), p3[0])

	// fmt.Printf("%#v\n", p1)
	// fmt.Printf("%#v\n", p2)
	// fmt.Printf("%#v\n", p3[0])
	// m1 := make(map[string]interface{}, 10)
	// dataM := `{"name":"lisi","age":90}`
	// m1["data"] = dataM
	// b, _ := json.Marshal(m1)
	// for k, v := range m1 {
	// 	fmt.Printf("key:%v, value:%v\n", k, v)
	// }
	// fmt.Printf("%s\n", b)
}

func filter(ctx *context.Context) {

}
