package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.0.105"))
	if err != nil {
		fmt.Println("Init elastic failed, error:", err)
	}
	fmt.Println("connect to elastic success")
	p1 := Student{
		Name:    "zhangsan",
		Age:     12,
		Married: false,
	}
	put1, err := client.Index().
		Index("student").
		Type("go").
		BodyJson(p1). // 把Go的变量转换成json字符串
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexde student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Index)
}
