package main

import "fmt"

func main() {
	s := make([]int, 0, 10)
	s = append(s, 4, 1, 8, 9)
	fmt.Println(s)
}
