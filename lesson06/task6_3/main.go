package main

import "fmt"

type contract struct {
	ID     int
	Number string
	Date   string
}

func main() {
	c := contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Println("Договор №", c.Number, "от", c.Date)
}
