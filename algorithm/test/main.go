package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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
