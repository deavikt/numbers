package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Numbers  []int
	Port     string
	DataPath string
}

func (srv *Server) Start() {
	http.HandleFunc("/data", srv.requestHandler)

	err := http.ListenAndServe(srv.Port, nil)

	if err != nil {
		log.Println("server startup error: ", err)
	}
}

func (srv *Server) requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		srv.handleGET(w)
	case http.MethodPost:
		srv.handlePOST(r)
	case http.MethodDelete:
		srv.handleDELETE()
	default:
		http.Error(w, "This request method isn't supported", http.StatusMethodNotAllowed)
	}
}

func (srv *Server) handleGET(w http.ResponseWriter) {
	sum := NumbersSum{
		Sum: srv.getNumbersSum(),
	}

	raw, err := json.Marshal(sum)

	if err != nil {
		log.Println(err)
	}

	w.Write(raw)
}

func (srv *Server) handlePOST(r *http.Request) {
	var number Number

	log.Println("POST request")

	raw, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
	}

	if err := json.Unmarshal(raw, &number); err != nil {
		log.Println(err)
	} else {
		srv.addNumber(number.Value)
		log.Println("number was successfully added")
	}
}

func (srv *Server) handleDELETE() {
	srv.deleteNumbers()
}
