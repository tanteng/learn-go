package main

import (
	"context"
	"fmt"
	"time"
)

// 主函数设置一个 ctx，传入每个 goroutine 第一个参数，可以实现超时退出协程的效果
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
		fmt.Println("end_func")
	}()
	return ch
}
