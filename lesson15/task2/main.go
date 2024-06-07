package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var j int32
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&j, 1)
		}()
	}
	wg.Wait()
}
