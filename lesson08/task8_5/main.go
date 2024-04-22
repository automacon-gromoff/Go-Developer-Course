package main

import "fmt"

func main() {
	m := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	m["бегемот"] = 2
	fmt.Println(m)
}
