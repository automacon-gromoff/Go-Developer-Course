package main

import "fmt"

type square int

func main() {
	var s square = 34
	s += 10
	printSquare(s)
}

func printSquare(s square) {
	fmt.Printf("%d м²", s)
}
