package server

import "log"

func (srv *Server) addNumber(number int) {
	srv.Numbers = append(srv.Numbers, number)
	log.Println("number was successfully added")
}

func (srv *Server) deleteNumbers() {
	srv.Numbers = nil
	log.Println("numbers were successfully deleted")
}

func (srv *Server) getNumbersSum() int {
	sum := 0

	for _, number := range srv.Numbers {
		sum += number
	}

	log.Println("numbers sum = ", sum)

	return sum
}
