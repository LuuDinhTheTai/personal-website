package server

import (
	"log"
	"net/http"
)

type IServer interface {
	InitializeRoutes() error
	Run()
}

type Server struct {
	host string
	port string
}

func NewServer(host string, port string) Server {
	return Server{
		host: host,
		port: port,
	}
}

func (s Server) InitializeRoutes() error {
	fs := http.FileServer(http.Dir("web/static"))

	http.Handle("/", fs)

	return nil
}

func (s Server) Run() {
	err := http.ListenAndServe(s.host+":"+s.port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
