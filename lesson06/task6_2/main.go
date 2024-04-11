package main

import "fmt"

type contract struct {
	ID     int
	Number string
	Date   string
}

func main() {
	c1 := contract{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v", c1)
}
