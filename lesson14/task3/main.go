package main

import "fmt"

func main() {
	ch := make(chan string, 4)
	ch <- "Привет"
	ch <- "буферизованный канал!"
	close(ch)
	for s := range ch {
		fmt.Println(s)
	}
}
