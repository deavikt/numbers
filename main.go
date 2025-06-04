package main

import (
	"encoding/json"
	"log"
	"numbers/server"
)

type Number struct {
	Value int `json:"value"`
}

func main() {
	number := Number{
		Value: 1000,
	}

	ConvertToJson(number)

	server.Start()
}

func ConvertToJson(number Number) {
	data, err := json.Marshal(number)

	if err != nil {
		log.Println(err)
	}

	log.Printf("%s", data)
}
