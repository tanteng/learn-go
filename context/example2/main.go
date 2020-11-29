package main

import (
	"context"
	"fmt"
	"time"
)

// 定义一个 ctx，主动取消协程，防止协程泄露
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}

	time.Sleep(time.Second * 80)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				time.Sleep(2 * time.Second)
				fmt.Println("canceled")
				return
			case ch <- n:
				fmt.Println("gen:", n)
				n++
				time.Sleep(time.Second)
			}
		}
	}()
	return ch
}
