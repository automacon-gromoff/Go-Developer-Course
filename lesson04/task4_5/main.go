package main

import "fmt"

func main() {
	defer fmt.Println("завершение работы")
	hello()
}

func hello() {
	fmt.Println("Hello, Go!")
}
