package handler

import (
	"cruiseapp/dto"
	"cruiseapp/model"
	"cruiseapp/repository/factory"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateShipModel(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateShipModelRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := factory.GetRepoFactory(r).CreateShipModelRepo()
	var sm model.ShipModel
	sm.Name = req.Name
	repo.Save(&sm)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(sm); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

// TODO validate if request fields not empty
// TODO validate ship id
func CreateShip(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateShipRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := factory.GetRepoFactory(r).CreateShipRepo()
	var s model.Ship
	s.Name = req.Name
	s.SerialNumber = req.SerialNumber
	s.ShipModelId = req.ShipModelId
	err = repo.Save(&s)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(&s); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
