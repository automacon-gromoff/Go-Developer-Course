package main

import "fmt"

func main() {
	m := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	delete(m, "бегемот")
	fmt.Println(m)
}
