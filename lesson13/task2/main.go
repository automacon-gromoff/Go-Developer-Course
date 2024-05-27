package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	tenat    string
}

func main() {
	c := contract{
		Number:   2,
		Landlord: "Остап",
		tenat:    "Шура",
	}

	str, err := json.Marshal(c)
	if err != nil {
		log.Fatalln("error:", err)
	}
	fmt.Printf("%+v", string(str))
}
