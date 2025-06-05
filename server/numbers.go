package server

import "log"

type Number struct {
	Value int `json:"value"`
}

type NumbersSum struct {
	Sum int `json:"sum"`
}

func (srv *Server) addNumber(number int) {
	srv.Numbers = append(srv.Numbers, number)
	log.Println("number was succussfully added")
}

func (srv *Server) deleteNumbers() {
	srv.Numbers = nil
	log.Println("numbers were succussfully deleted")
}

func (srv *Server) getNumbersSum() int {
	sum := 0

	for number := range srv.Numbers {
		sum += number
	}

	log.Println("numbers sum = ", sum)

	return sum
}
