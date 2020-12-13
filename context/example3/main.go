package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	buffer := make(chan int, 3)

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				close(buffer)
				return
			case buffer <- rand.Intn(90):
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	for j := range buffer {
		fmt.Println(j)
	}

	fmt.Println("aa")
}
