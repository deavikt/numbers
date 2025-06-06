package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (srv *Server) Start() {
	srv.Mux.HandleFunc(srv.DataPath, srv.requestHandler)
	srv.Mux.Handle("/", http.StripPrefix("/", srv.FileServer))

	err := http.ListenAndServe(srv.Port, srv.Mux)

	if err != nil {
		log.Println("server startup error: ", err)
	}
}

func (srv *Server) requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		srv.handleGET(w)
	case http.MethodPost:
		srv.handlePOST(w, r)
	case http.MethodDelete:
		srv.handleDELETE()
	default:
		http.Error(w, "This request method isn't supported", http.StatusMethodNotAllowed)
	}
}

func (srv *Server) handleGET(w http.ResponseWriter) {
	log.Println("GET request")

	sum := NumbersSum{
		Sum: srv.getNumbersSum(),
	}

	raw, err := json.Marshal(&sum)

	if err != nil {
		log.Println(err)
	}

	w.Write(raw)
}

func (srv *Server) handlePOST(w http.ResponseWriter, r *http.Request) {
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
	}

	w.WriteHeader(http.StatusOK)
}

func (srv *Server) handleDELETE() {
	log.Println("DELETE request")
	srv.deleteNumbers()
}
