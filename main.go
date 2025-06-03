package main

import (
	"encoding/json"
	"log"
)

type Number struct {
	value int
}

func main() {
	number := Number{
		value: 1000,
	}

	convertToJson(number)
}

func convertToJson(number Number) {
	data, err := json.Marshal(number)

	if err != nil {
		log.Println(err)
	}

	log.Printf("%s", data)
}
