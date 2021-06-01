package main

import (
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	//与服务端创建连接
	client,err := rpc.DialHTTP("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = client.Close()
	}()
	//调用server端提供的服务，三个参数分别是提供服务的方法，方法的入参，方法的返回
	var output string
	tk := time.Tick(time.Second)
	for i := range tk {
		err = client.Call("User.SayHello", "张三" + i.Format("2006-01-02 15:04:05"),&output)
		fmt.Println("结果为：",output)
	}
}
