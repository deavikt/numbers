package server

import (
	"log"
	"net/http"
	"strconv"
)

const (
	dataPath string = "/data"
	port     string = ":8080"
)

func Start() {
	http.HandleFunc(dataPath, requestHandler)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Println("server startup error: ", err)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGET()
	case http.MethodPost:
		handlePOST(r)
	case http.MethodDelete:
		handleDELETE()
	default:
		http.Error(w, "This request method isn't supported", http.StatusMethodNotAllowed)
	}
}

func handleGET() {
	log.Println("GET request")
	getNumbersSum()
}

func handlePOST(r *http.Request) {
	log.Println("POST request")

	err := r.ParseForm()

	if err != nil {
		log.Println(err)
	}

	stringNumber := r.FormValue("value")
	intNumber, err := strconv.Atoi(stringNumber)

	if err != nil {
		log.Println("incorrect input")
	} else {
		addNumber(intNumber)
	}
}

func handleDELETE() {
	log.Println("DELETE request")
	deleteNumbers()
}
