package main

import (
	"fmt"
	"time"
)

// 10个任务，每个任务按指定时间周期执行
func job(input interface{}) {
	fmt.Printf("working on %v", input)
}

func main() {
	jobChan := make(chan func(interface{}))
	var (
		input    int
		duration time.Duration
	)
	for i := 0; i < 10; i++ {
		go work(jobChan, input, duration)
	}
	for i := 0; i < 10; i++ {
		jobChan <- job
	}
}

func work(jobChan chan func(interface{}), input interface{}, duration time.Duration) {
	t := time.Tick(duration)
	for range t {
		j := <-jobChan
		j(input)
	}
}

// 一个表有三列：人名、课程、分数，查找每个课程最高分数的人名
var sqlString = `select course,max(score) 
				from table 
				group by course e 
				join 
				select name, score 
				from table d 
				on e.max(score) = d.score`
