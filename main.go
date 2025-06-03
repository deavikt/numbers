package main

import "log"

type Number struct {
	value int
}

func main() {
	number := Number{
		value: 1000,
	}

	log.Println(number.value)
}
