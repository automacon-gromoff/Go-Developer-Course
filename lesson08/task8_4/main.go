package main

import "fmt"

func main() {
	m := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	m["выдра"] = struct{}{}
	fmt.Println(m)
}
