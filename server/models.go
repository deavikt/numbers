package server

type Server struct {
	Numbers  []int
	Port     string
	DataPath string
}

type Number struct {
	Value int `json:"value"`
}

type NumbersSum struct {
	Sum int `json:"sum"`
}
