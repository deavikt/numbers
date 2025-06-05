package main

import (
	"numbers/server"
)

func main() {
	srv := server.Server{
		Numbers:  nil,
		Port:     ":8081",
		DataPath: "/data",
	}

	srv.Start()
}
