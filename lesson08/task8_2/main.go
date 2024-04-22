package main

import "fmt"

func main() {
	m := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	for _, v := range []string{"слон", "бегемот", "осьминог"} {
		a, ok := m[v]
		fmt.Printf(":Животное: %s, количество: %d (есть в карте: %v)\n", v, a, ok)
	}
}
