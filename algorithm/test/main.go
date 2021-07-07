package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 主程等待子协程指定时间，超时后就不等了

var wg sync.WaitGroup

func main() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	wg.Add(1)
	go f(ctx)
	wg.Wait()
	fmt.Println(time.Since(start))
}
func f(ctx context.Context) {
	fmt.Println("子协程...")
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	}
}

