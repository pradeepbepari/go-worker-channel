package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeepbepari/jsonplaceholder/routes"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}
func (s *Server) Run() error {
	server := make(chan *mux.Router)
	go func(ch chan *mux.Router) {
		ch <- mux.NewRouter()
	}(server)
	router := <-server
	routes.RegisteredRoutes(router)
	log.Println("Port Running on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
