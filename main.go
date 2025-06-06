package main

import (
	"net/http"
	"numbers/server"
)

func main() {
	srv := server.Server{
		Numbers:    nil,
		Mux:        http.NewServeMux(),
		FileServer: http.FileServer(http.Dir("./ui/")),
		Port:       ":8081",
		DataPath:   "/data",
	}

	srv.Start()
}
