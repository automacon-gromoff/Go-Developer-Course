package main

import "fmt"

const (
	firstNum = iota
	secondNum
)

var count = 0

func main() {
	fibonacci(firstNum, secondNum, 23)
}

func fibonacci(i int, j int, n int) {
	if count == 23 {
		return
	}
	count++
	fmt.Println(i)
	fibonacci(j, i+j, n)
}
