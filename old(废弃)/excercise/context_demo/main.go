package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	go f2(ctx)
	defer wg.Done()
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
		fmt.Println("parent context demo")
		time.Sleep(time.Millisecond * 500)
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
		fmt.Println("son context demo")
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
