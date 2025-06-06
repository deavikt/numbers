package server

import "net/http"

type Server struct {
	Numbers    []int
	Mux        *http.ServeMux
	FileServer http.Handler
	Port       string
	DataPath   string
}

type Number struct {
	Value int `json:"value"`
}

type NumbersSum struct {
	Sum int `json:"sum"`
}
