package main

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	wg   sync.WaitGroup
	work chan func()
}

func NewPool(workers int) *Pool {
	p := &Pool{
		wg:   sync.WaitGroup{},
		work: make(chan func()),
	}
	p.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
					p.wg.Done()
				}
			}()
			for fn := range p.work {
				fn()
			}
			p.wg.Done()
		}()
	}
	return p
}

func (p *Pool) Add(fn func()) {
	p.work <- fn
}

func (p *Pool) Run() {
	close(p.work)
	p.wg.Wait()
}

func parseTask(i int) func() {
	return func() {
		time.Sleep(time.Second)
		fmt.Println("finish parse ", i)
	}
}

func main() {
	p := NewPool(20)
	for i := 0; i < 100; i++ {
		p.Add(parseTask(i))
	}
	p.Run()
}

// 一个表有三列：人名、课程、分数，查找每个课程最高分数的人名
var sqlString = `select course,max(score) 
				from table 
				group by course e 
				join 
				select name, score 
				from table d 
				on e.max(score) = d.score`
