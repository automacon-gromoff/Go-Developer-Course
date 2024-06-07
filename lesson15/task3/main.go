package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	var c int
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			c++
		}()
	}
	wg.Wait()
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(c)
}
