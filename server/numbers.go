package server

import "log"

var numbers []Number = nil

type Number struct {
	Value int `json:"value"`
}

type NumbersSum struct {
	Sum int `json:"sum"`
}

func addNumber(number int) {
	numbers = append(numbers, convertIntToNumber(number))
	log.Println("number was succussfully added")
}

func deleteNumbers() {
	numbers = nil
	log.Println("numbers were succussfully deleted")
}

func getNumbersSum() int {
	sum := 0

	for number := range numbers {
		sum += number
	}

	log.Println("numbers sum = ", sum)

	return sum
}

func convertIntToNumber(number int) Number {
	return Number{Value: number}
}
