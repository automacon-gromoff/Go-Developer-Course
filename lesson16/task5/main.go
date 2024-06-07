package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	stop := 0
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				fmt.Println("сложные вычисления горутины:", i)
				if stop == 1 {
					break
				}
				time.Sleep(time.Second)
			}
			fmt.Println("stop горутина:", i)
		}()
	}
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		stop = 1
	}()
	wg.Wait()
}
