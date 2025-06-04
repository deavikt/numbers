package server

import (
	"log"
	"net/http"
)

const (
	dataPath string = "/data"
	port     string = ":8080"
)

func Start() {
	http.HandleFunc(dataPath, requestHandler)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("server started at port " + port)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getHandler()
	case http.MethodPost:
		postHandler()
	case http.MethodDelete:
		deleteHandler()
	}
}

func getHandler() {
	log.Println("GET")
}

func postHandler() {
	log.Println("POST")
}

func deleteHandler() {
	log.Println("DELETE")
}
