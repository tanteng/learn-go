package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// main 设置一个带有超时时间的 ctx，子协程通过监听 ctx.Done 判断是否超时
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	wg.Add(2)

	go func() {
		go watch(ctx, 1)
		go watch(ctx, 2)
	}()

	wg.Wait()

	fmt.Println("All finished")
}

func watch(ctx context.Context, flag int) {
	defer wg.Done()

	done := make(chan bool)
	go func(done chan<- bool) {
		fmt.Printf("[flag:%d]doing something\n", flag)
		time.Sleep(2 * time.Second)
		done <- true
	}(done)

	select {
	case <-ctx.Done():
		fmt.Printf("[flag:%d]ctx done,err: %s\n", flag, ctx.Err())
		return
	case <-done:
		fmt.Printf("[flag:%d]finished\n", flag)
		return
	}
}
