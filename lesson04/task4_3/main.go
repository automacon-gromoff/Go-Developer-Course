package main

import "fmt"

func main() {
	hello(func() { fmt.Println("Hello, Go!") })
}

func hello(f func()) {
	f()
}
