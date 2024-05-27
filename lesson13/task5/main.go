package main

import (
	format "format_v1.0.0"
	"log"
)

func main() {
	err := format.Do("patients", "result")
	if err != nil {
		log.Fatalln(err)
	}
}
