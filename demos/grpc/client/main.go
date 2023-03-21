package main

import (
	"bufio"
	"client/pb"
	"context"
	"flag"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:2023", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	// 通过TLS/SSL建立安装连接
	//creds, _ := credentials.NewClientTLSFromFile("../certFile.cert", "")
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	runSayHello(c, ctx)
	runLotsOfReplies(c, ctx)
	runLotsOfRequests(c, ctx)
	runBidiHello(c, ctx)
}

func runSayHello(c pb.GreeterClient, ctx context.Context) {

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		s := status.Convert(err)        // 将err转为status
		for _, d := range s.Details() { // 获取details
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("Quota failure: %s\n", info)
			default:
				fmt.Printf("Unexpected type: %s\n", info)
			}
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Greeting: %s\n", r.GetReply())
}

func runLotsOfReplies(c pb.GreeterClient, ctx context.Context) {
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// 接收服务端返回的流式数据，当收到io.EOF或错误时退出
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("got reply: %q\n", res.GetReply())
	}
}

func runLotsOfRequests(c pb.GreeterClient, ctx context.Context) {
	stream, err := c.LotsOfRequests(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	names := []string{"七米", "q1mi", "沙河娜扎"}
	for _, n := range names {
		// 发送流式数据
		err = stream.Send(&pb.HelloRequest{Name: n})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("got reply: %v", res.GetReply())
}

func runBidiHello(c pb.GreeterClient, ctx context.Context) {
	stream, err := c.BidiHello(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("AI：%s\n", in.GetReply())
		}
	}()
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		if err = stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			fmt.Println(err)
			return
		}
	}
	stream.CloseSend()
}

// unaryCallWithMetadata 普通RPC调用客户端metadata操作
func unaryCallWithMetadata(c pb.GreeterClient, name string) {
	fmt.Println("--- UnarySayHello client---")
	// 创建metadata
	md := metadata.Pairs(
		"token", "app-test-q1mi",
		"request_id", "1234567",
	)

	// 基于metadata创建context.
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 添加一些metadata到context
	md2 := metadata.New(map[string]string{
		"k2": "v2",
		"K2": "v2.1"}, // 所有的键将自动转换为小写，因此“k2”和“K2”将是相同的键，它们的值将合并到相同的列表中。因为 type MD map[string][]string
	)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Join(md2))

	// 添加一些metadata到context，此函数也可用来创建metadata
	ctx = metadata.AppendToOutgoingContext(ctx, "k3", "v3")

	// RPC调用
	var header, trailer metadata.MD
	r, err := c.SayHello(
		ctx,
		&pb.HelloRequest{Name: name},
		grpc.Header(&header),   // 接收服务端发来的header
		grpc.Trailer(&trailer), // 接收服务端发来的trailer
	)
	if err != nil {
		log.Printf("failed to call SayHello: %v", err)
		return
	}
	// 从header中取location
	if t, ok := header["location"]; ok {
		fmt.Printf("location from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("location expected but doesn't exist in header")
		return
	}
	// 获取响应结果
	fmt.Printf("got response: %s\n", r.Reply)
	// 从trailer中取timestamp
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}

// bidirectionalWithMetadata 流式RPC调用客户端metadata操作，metadata的其他操作参考普通调用的函数，次数不赘述
func bidirectionalWithMetadata(c pb.GreeterClient, name string) {
	// 创建metadata和context.
	md := metadata.Pairs("token", "app-test-q1mi")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 使用带有metadata的context执行RPC调用.
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("failed to call BidiHello: %v\n", err)
	}

	go func() {
		// 当header到达时读取header.
		header, err := stream.Header()
		if err != nil {
			log.Fatalf("failed to get header from stream: %v", err)
		}
		// 从返回响应的header中读取数据.
		if l, ok := header["location"]; ok {
			fmt.Printf("location from header:\n")
			for i, e := range l {
				fmt.Printf(" %d. %s\n", i, e)
			}
		} else {
			log.Println("location expected but doesn't exist in header")
			return
		}

		// 发送所有的请求数据到server.
		for i := 0; i < 5; i++ {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// 读取所有的响应.
	var rpcStatus error
	fmt.Printf("got response:\n")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", r.Reply)
	}
	if rpcStatus != io.EOF {
		log.Printf("failed to finish server streaming: %v", rpcStatus)
		return
	}

	// 当RPC结束时读取trailer
	trailer := stream.Trailer()
	// 从返回响应的trailer中读取metadata.
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}
