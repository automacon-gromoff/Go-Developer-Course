package main

import (
	"fmt"
)

func main() {
	fruitMarket("яблоки")
	fruitMarket("галоперидол")
}

func fruitMarket(s string) {
	m := map[string]int{
		"апельсин": 5,
		"яблоки":   3,
		"сливы":    1,
		"груши":    0,
	}
	num, ok := m[s]
	if !ok {
		fmt.Println("фрукт", s, "отсутствует")
		return
	}
	fmt.Println(s, "-", num)
}
