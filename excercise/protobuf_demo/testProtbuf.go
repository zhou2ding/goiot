package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"protobufDemo/pbf"
)

func main() {
	phone := []int64{1,3,8,10}
	name := "zhangsan"
	helloData := &pbf.Hello{
		Name:&name,
		Phone:phone,
		Addr: "河北",
	}
	fmt.Println(*helloData.Name)
	fmt.Println(helloData.GetName())
	ret,_ := proto.Marshal(helloData)
	fmt.Println(ret)
	helloData2 := pbf.Hello{}
	_ = proto.Unmarshal(ret,&helloData2)
	fmt.Println(helloData2.Addr)
}
