package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
)

func mySarama() {
	config := sarama.NewConfig()
	// 设置配置
	config.Producer.RequiredAcks = sarama.WaitForAll          // 等待leader和followerACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个分区
	config.Producer.Return.Successes = true                   // 成功发送的消息将在success channel返回
	// 构造消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is s test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, error:", err)
		return
	}
	fmt.Println("连接kafka成功")
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, error:", err)
		return
	}
	fmt.Printf("pid:%v, offser:%v\n", pid, offset)
}

func myTail() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开，切割文件时使用
		Follow:    true,                                 // 跟随文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个位置开始读
		MustExist: false,                                // 文件不存在时不报错
		Poll:      true,                                 //
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, error:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen, filename:", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line: ", line)
	}
}
func main() {
	// MySarama()
}
