package main

import "fmt"

func main() {
	ch := make(chan string)
	go write(ch)
	res := read(ch)
	fmt.Println(res)
}

func write(ch chan<- string) {
	ch <- "Привет, строковый канал!"
}

func read(ch <-chan string) string {
	return <-ch
}
