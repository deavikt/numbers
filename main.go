package main

import (
	"net/http"
	"numbers/server"
)

const (
	port     = ":8000"
	homePath = "/"
	dataPath = "/data"
)

func main() {
	srv := server.Create(nil, http.NewServeMux(), port, homePath, dataPath)
	srv.Start()
}
