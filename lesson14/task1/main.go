package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("Привет из дочерней горутины!")
		ch <- struct{}{}
	}()
	<-ch
}
