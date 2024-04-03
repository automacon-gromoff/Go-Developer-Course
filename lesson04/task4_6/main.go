package main

import "fmt"

var i, j = 0, 1

func main() {
	fibonacci(23)
}

func fibonacci(n int) {
	if n == 0 {
		return
	}
	fib := i
	fmt.Println(fib)
	i = j
	j += fib
	n--
	fibonacci(n)
}
