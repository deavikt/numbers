package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Create(
	numbers []int,
	mux *http.ServeMux,
	port, homePath, dataPath string) Server {

	return Server{
		Numbers:  numbers,
		Mux:      mux,
		Port:     port,
		HomePath: homePath,
		DataPath: dataPath,
	}
}

func (srv *Server) Start() {
	srv.setUpMux()

	err := http.ListenAndServe(srv.Port, srv.Mux)

	if err != nil {
		log.Println("server startup error: ", err)
	}
}

func (srv *Server) setUpMux() {
	fileServerHandler := http.FileServer(http.Dir("./static/"))

	srv.Mux.Handle(srv.HomePath, http.StripPrefix(srv.HomePath, fileServerHandler))
	srv.Mux.HandleFunc(srv.DataPath, srv.requestHandler)
}

func (srv *Server) requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			log.Println("GET request")
			srv.postNumbersSumJSON(w)
		}
	case http.MethodPost:
		{
			log.Println("POST request")
			srv.getNumberJSON(r)
		}
	case http.MethodDelete:
		{
			log.Println("DELETE request")
			srv.deleteNumbers()
		}
	default:
		http.Error(w, "This request method isn't supported", http.StatusMethodNotAllowed)
	}
}

func (srv *Server) postNumbersSumJSON(w http.ResponseWriter) {
	sum := NumbersSum{
		Sum: srv.getNumbersSum(),
	}

	raw, err := json.Marshal(&sum)

	if err != nil {
		log.Println(err)
	}

	w.Write(raw)
}

func (srv *Server) getNumberJSON(r *http.Request) {
	var number Number

	raw, readingErr := io.ReadAll(r.Body)

	if readingErr != nil {
		log.Println(readingErr)
	} else {
		marshallingErr := json.Unmarshal(raw, &number)

		if marshallingErr != nil {
			log.Println(marshallingErr)
		} else {
			srv.addNumber(number.Value)
		}
	}
}
