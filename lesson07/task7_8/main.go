package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	s = append(s[:3], s[4:]...)
	fmt.Println(s)
}
