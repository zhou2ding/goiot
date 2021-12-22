package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9090/xxx?name=dq&age=18")
	if err != nil {
		fmt.Printf("get url failed, error:%v\n", err)
		return
	}
	defer resp.Body.Close()
	// 自定义请求

	// 比较麻烦
	// var data [1029]byte
	// n, _ := resp.Body.Read(data[:])
	// fmt.Println(data[:n])
	// resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body failed, error:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
