package main

import (
	db "cruiseapp/database"
	"cruiseapp/repository/factory"
	"cruiseapp/server"
	"log"
	"net/http"
)

func main() {
	srv := server.NewServer()

	log.Fatal(http.ListenAndServe(srv.Addr,
		db.DbMiddleware(factory.PgRepoFactoryMiddleware(srv.Router))))
}
