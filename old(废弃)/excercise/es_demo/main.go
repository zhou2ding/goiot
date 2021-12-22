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
		BodyJson(p1). // 必须是json或能转成json的数据，如json字符串、map[string]interface{}、结构体
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexde student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Index)
}
