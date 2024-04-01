package main

import "fmt"

func main() {
	hello()()
}

func hello() func() {
	return func() { fmt.Println("Hello, Go!") }
}
