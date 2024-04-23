package main

import "fmt"

func main() {
	checkFood("яблоко")
	checkFood("помидор")
	checkFood("мухомор")
}

func checkFood(s string) {
	switch s {
	case "груша", "яблоко", "апельсин":
		fmt.Println("это фрукт")
	case "тыква", "огурец", "помидор":
		fmt.Println("это овощ")
	default:
		fmt.Println("что-то странное...")
	}
}
