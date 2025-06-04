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
	http.HandleFunc(dataPath, dataHandler)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("server started at port " + port)
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("data"))
}
