package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// main 设置一个 3s 超时的 ctx，传给每个 goroutine 第一个参数，每个 goroutine 又发起协程处理任务，
// 假设需要 5s，通过监听 ctx.Done 判断是否超时退出
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg.Add(2)
	go watch(ctx, 1)
	go watch(ctx, 2)
	wg.Wait()

	fmt.Println("finished")
}

func watch(ctx context.Context, flag int) {
	defer wg.Done()

	go func() {
		fmt.Printf("doing something flag:%d\n", flag)
		time.Sleep(5 * time.Second)
		fmt.Println("finished flag:", flag)
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("watch %d %s\n", flag, ctx.Err())
		return
	}
}
