package handler

import (
	db "cruiseapp/database"
	"cruiseapp/dto"
	"cruiseapp/model"
	"cruiseapp/repository/factory"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPort(w http.ResponseWriter, r *http.Request) {
	var p model.Port
	id := r.PathValue("id")
	conn := db.GetDb(r)
	row := conn.QueryRow(`SELECT id, location FROM port WHERE id = $1`, id)
	err := row.Scan(&p.Id, &p.Location)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusFound)
	_ = json.NewEncoder(w).Encode(p)
}

func CreatePort(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePortRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// conn := db.GetDb(r)
	// pr := port.NewPgPortRepository(conn)
	pr := factory.GetRepoFactory(r).CreatePortRepo()
	// TODO handle error
	var p model.Port
	p.Location = req.Location
	pr.Save(&p)
	w.WriteHeader(http.StatusFound)
	_ = json.NewEncoder(w).Encode(p)
}
