package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	tenat    string
}

func main() {
	c := contract{
		Number:   1,
		Landlord: "Остап",
		tenat:    "Шура",
	}

	str, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", string(str))
}
