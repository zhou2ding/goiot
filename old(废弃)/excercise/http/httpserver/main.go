package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func homeHandlerfunc(w http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadFile("./xxx.html")
	if err != nil {
		fmt.Printf("read html failed, error:%v\n", err)
	}
	w.Write(str)
}

func xxxHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都在URL上(query param)，请求体中是没有数据的
	// fmt.Println(r.URL)
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Printf("name:%v, age:%v\n", name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/home", homeHandlerfunc)
	http.HandleFunc("/xxx", xxxHandlerFunc)

	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Printf("start http server failed, error:%v\n", err)
		return
	}
}
