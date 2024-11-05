package server

import (
	dm "cruiseapp/database/middleware"
	"cruiseapp/handler"
	fm "cruiseapp/repository/factory/middleware"
	m "cruiseapp/server/middleware"
	"cruiseapp/ws"
	"net/http"
	"time"
)

func NewServer() http.Server {
	router := http.NewServeMux()

	router.HandleFunc("POST /port", handler.CreatePort)
	router.HandleFunc("GET /port/{id}", handler.RetrievePort)
	router.HandleFunc("PUT /port/{id}", handler.UpdatePort)
	router.HandleFunc("DELETE /port/{id}", handler.DeletePort)

	router.HandleFunc("POST /ship/model", handler.CreateShipModel)
	router.HandleFunc("GET /ship/model/{id}", handler.RetrieveShipModel)
	router.HandleFunc("PUT /ship/model/{id}", handler.UpdateShipModel)
	router.HandleFunc("DELETE /ship/model/{id}", handler.DeleteShipModel)

	router.HandleFunc("POST /ship", handler.CreateShip)
	router.HandleFunc("GET /ship/{id}", handler.RetrieveShip)
	router.HandleFunc("PUT /ship/{id}", handler.UpdateShip)
	router.HandleFunc("DELETE /ship/{id}", handler.DeleteShip)

	router.HandleFunc("POST /crew/rank", handler.CreateCrewRank)
	router.HandleFunc("GET /crew/rank/{id}", handler.RetrieveCrewRank)
	router.HandleFunc("PUT /crew/rank/{id}", handler.UpdateCrewRank)
	router.HandleFunc("DELETE /crew/rank/{id}", handler.DeleteCrewRank)

	router.HandleFunc("POST /crew/member", handler.CreateCrewMember)
	router.HandleFunc("GET /crew/member/{id}", handler.RetrieveCrewMember)
	router.HandleFunc("PUT /crew/member/{id}", handler.UpdateCrewMember)
	router.HandleFunc("DELETE /crew/member/{id}", handler.DeleteCrewMember)

	router.HandleFunc("POST /person", handler.CreatePerson)
	router.HandleFunc("GET /person/{id}", handler.RetrievePerson)
	router.HandleFunc("PUT /person/{id}", handler.UpdatePerson)
	router.HandleFunc("DELETE /person/{id}", handler.DeletePerson)

	router.HandleFunc("POST /cruise", handler.CreateCruise)
	router.HandleFunc("GET /cruise/{id}", handler.RetrieveCruise)
	router.HandleFunc("PUT /cruise/{id}", handler.UpdateCruise)
	router.HandleFunc("DELETE /cruise/{id}", handler.DeleteCruise)

	router.HandleFunc("GET /stats", handler.StatisticsHandler)

	hub := ws.NewHub()
	go hub.Run()

	handler := m.ChainMiddleware(
		dm.DbMiddleware,
		fm.PgRepoFactoryMiddleware,
		ws.WsHubMiddleware(hub),
	)(router)

	return http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
