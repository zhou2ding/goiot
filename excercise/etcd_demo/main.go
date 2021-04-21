package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func myWatch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, error:", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// 派一个哨兵，一直监视“zhangsan”这个key的变化（新增、修改、删除等）
	ch := cli.Watch(context.Background(), "zhangsan")
	// 尝试从通道取值（监视的信息）
	for wresp := range ch {
		for _, env := range wresp.Events {
			fmt.Printf("Type:%v, key:%s, value:%s\n", env.Type, env.Kv.Key, env.Kv.Value)
		}
	}
}

func myEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, error:", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "zhangsan", "18") //put操作
	if err != nil {
		fmt.Println("put to etcd failed, error:", err)
		return
	}
	cancel()
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "zhangsan") //get操作
	if err != nil {
		fmt.Println("get from etcd failed, error:", err)
	}
	cancel()
	// 从resp的KV中取值
	for _, ev := range resp.Kvs {
		fmt.Printf("key:%s, value:%s\n", ev.Key, ev.Value)
	}
}

func main() {
}
