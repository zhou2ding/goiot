package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "grpc_protobuf"
	"net"
)

type User struct {}

func (s *User) SayHello(ctx context.Context, input *hello.HelloReq) (output *hello.HelloResp,err error){
	output = &hello.HelloResp{
		Output: "hello, " + input.Input,
	}
	return
}

func main() {
	server := grpc.NewServer()
	hello.RegisterTestServiceServer(server,&User{})
	listener,err := net.Listen("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = server.Serve(listener)
}
