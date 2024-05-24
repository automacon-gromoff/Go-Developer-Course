package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	tenat    string
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		tenat:    "Шура Балаганов",
	}

	cc := contracts{}
	cc.List = append(cc.List, c)

	res, err := xml.MarshalIndent(cc, "", "  ")
	if err != nil {
		log.Fatalln("error:", err)
	}
	fmt.Println(xml.Header, string(res))
}
