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
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	wg.Add(2)
	go watch(ctx, 1)
	go watch(ctx, 2)
	wg.Wait()

	fmt.Println("All finished")
}

func watch(ctx context.Context, flag int) {
	defer wg.Done()

	done := make(chan bool)
	go func(done chan<- bool) {
		fmt.Printf("[flag:%d]doing something\n", flag)
		time.Sleep(6 * time.Second)
		done <- true
	}(done)

	select {
	case <-ctx.Done():
		fmt.Printf("[flag:%d]ctx done,err: %s\n", flag, ctx.Err())
		return
	case <-time.After(8 * time.Second):
		fmt.Printf("[flag:%d]timeout\n", flag)
		return
	case <-done:
		fmt.Printf("[flag:%d]finished\n", flag)
		return
	}
}
