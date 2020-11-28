package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 定义一个带超时时间的 ctx，传入每个 goroutine 第一个参数，实现超时退出协程的效果
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	wg.Add(2)
	go watch(ctx, 1)
	go watch(ctx, 2)
	wg.Wait()

	fmt.Println("finished")
}

func watch(ctx context.Context, flag int) {
	defer wg.Done()

	fmt.Printf("doing something flag:%d\n", flag)
	time.Sleep(5 * time.Second)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("watch %d %s\n", flag, ctx.Err())
			return
		case <-time.After(2 * time.Second):
			fmt.Println("2 second")
		}
	}
}
