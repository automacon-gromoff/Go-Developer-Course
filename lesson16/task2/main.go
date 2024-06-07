package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("отстановка горутины", i, ctx.Err())
					return
				default:
					fmt.Println("работа горутины", i)
					time.Sleep(400 * time.Millisecond)
				}
			}
		}()
	}
	wg.Wait()
}
