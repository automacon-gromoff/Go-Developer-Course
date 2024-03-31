package main

import "fmt"

func main() {
	a := 16
	b := 3
	result := a / b
	remainder := a % b

	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T", result, remainder, result)
}
