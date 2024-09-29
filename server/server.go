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

	router.HandleFunc("POST /crew/rank", handler.CreateCrewRank)
	router.HandleFunc("GET /crew/rank/{id}", handler.RetrieveCrewRank)
	router.HandleFunc("PUT /crew/rank/{id}", handler.UpdateCrewRank)
	router.HandleFunc("DELETE /crew/rank/{id}", handler.DeleteCrewRank)

	router.HandleFunc("POST /crew/member", handler.CreateCrewMember)
	router.HandleFunc("GET /crew/member/{id}", handler.RetrieveCrewMember)
	router.HandleFunc("PUT /crew/member/{id}", handler.UpdateCrewMember)
	router.HandleFunc("DELETE /crew/member/{id}", handler.DeleteCrewMember)

	return router
}
