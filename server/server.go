package server

import (
	"cruiseapp/handler"
	"fmt"
	"net/http"
)

type Server struct {
	Addr   string
	Router *http.ServeMux
}

func NewServer() Server {
	return Server{
		Addr:   ":8080",
		Router: Router(),
	}
}

func (srv *Server) Run() {
	fmt.Println("starting server")
	http.ListenAndServe(srv.Addr, srv.Router)
}

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /port/{id}", handler.GetPort)
	router.HandleFunc("POST /port", handler.CreatePort)

	router.HandleFunc("POST /ship/model", handler.CreateShipModel)
	router.HandleFunc("POST /ship", handler.CreateShip)
	return router
}
