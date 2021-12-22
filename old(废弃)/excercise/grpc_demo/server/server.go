package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type User struct {

}

func (u *User) SayHello(input string, output *string) error {
	*output =  "hello, " + input
	return nil
}

func main() {
	usr := new(User)
	//把usr对象注册到rpc服务中
	_ = rpc.Register(usr)
	//把usr提供的服务注册到HTTP协议上
	rpc.HandleHTTP()
	//监听tcp连接
	listener,err := net.Listen("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	//service接收监听器传入的http连接
	_ = http.Serve(listener, nil)
}

