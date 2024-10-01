package main

import (
	"cruiseapp/server"
	"log"
)

func main() {
	srv := server.NewServer()

	log.Fatal(srv.ListenAndServe())
}
