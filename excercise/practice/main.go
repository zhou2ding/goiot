package main

import (
	"crypto/md5"
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
	// fmt.Println(f2()) // 6，defer中x++是加的f2的参数x
	// fmt.Println(f3()) // 5
	// fmt.Println(f4()) // 5，defer中x++是加的defer函数的参数x
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
	// x := 1
	// y := 2
	// defer calc("AA", x, calc("A", x, y))
	// x = 10
	// defer calc("BB", x, calc("B", x, y))
	// y = 20

	// rand.Seed(time.Now().Unix())
	// t := time.Tick(time.Second)
	// for _ = range t {
	// 	fmt.Println(rand.Intn(10000))
	// }
	// m := sync.Map{}
	// m.Store("key", 0)
	// m.Store("key1", 1)
	// m.Store("key2", 2)
	// m.Store("key3", 3)
	// m.Store("key4", 4)
	// m.Store("key5", 5)
	// m.Range(func(k, v interface{}) bool {
	// 	fmt.Printf("key:%v,value:%v\n", k, v)
	// 	return true
	// })
	fmt.Println(getmd5("zhangsan"))

}

func calc(index string, a, b int) int {
	//结果：
	//A 1 2 3
	//B 10 2 12
	//BB 10 12 22
	//AA 1 3 4
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func getmd5(s string) string {
	tmp := md5.Sum([]byte(s))
	ret := string(tmp[:])
	return ret
}
