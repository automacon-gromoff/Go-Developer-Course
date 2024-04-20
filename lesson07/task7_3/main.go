package main

import "fmt"

func main() {
	s := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	s[2] = "персик"
	fmt.Println(s)
}
