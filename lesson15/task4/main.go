package main

import (
	"fmt"
	"sync"
)

var startOnce sync.Once

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			startOnce.Do(start)
			fmt.Println("горутина", i)
		}()
	}
	wg.Wait()
}

func start() {
	fmt.Println("starting...")
}
