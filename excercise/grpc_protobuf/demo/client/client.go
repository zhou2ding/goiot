package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	hello "grpc_protobuf"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	var waitIndex uint64
	cliApi, _ := api.NewClient(config)
	//获取符合筛选条件的所有服务实体。
	//前三个参数分别是配置文件中的name、tag，和"是否只保留通过筛选的服务"
	services,_,_ := cliApi.Health().Service("sayHello","sayhello",true,&api.QueryOptions{
		WaitIndex: waitIndex,
	})
	//services[0]还可以获得Node、Checks的具体信息
	url := fmt.Sprintf("%v:%v",services[0].Service.Address,services[0].Service.Port)

	//WithInsecure：跳过证书的验证
	conn,err := grpc.Dial(url,grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client := hello.NewTestServiceClient(conn)
	resp, err := client.SayHello(context.Background(),&hello.HelloReq{
		Input: "张三",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Output)
}
