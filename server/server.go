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

	router.HandleFunc("POST /port", handler.CreatePort)
	router.HandleFunc("GET /port/{id}", handler.GetPort)

	router.HandleFunc("POST /crank", handler.CreateCrewRank)
	router.HandleFunc("GET /crank/{id}", handler.RetrieveCrewRank)
	router.HandleFunc("PUT /crank/{id}", handler.UpdateCrewRank)
	router.HandleFunc("DELETE /crank/{id}", handler.DeleteCrewRank)

	return router
}
