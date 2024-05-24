package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	Tenat    string `json:"tenat"`
}

func main() {
	str := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`
	c := contract{}

	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		log.Fatalln("error:", err)
	}
	fmt.Printf("%+v", c)
}
