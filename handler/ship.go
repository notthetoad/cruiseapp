package handler

import (
	"cruiseapp/dto"
	"cruiseapp/handler/util"
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

func RetrieveShipModel(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	sm, err := factory.GetRepoFactory(r).CreateShipModelRepo().FindById(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(&sm); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func UpdateShipModel(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateShipModelRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := factory.GetRepoFactory(r).CreateShipModelRepo()
	var sm model.ShipModel
	id := util.ParseIdFromRequest(r)
	sm.Id = id
	sm.Name = req.Name
	repo.Update(&sm)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteShipModel(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreateShipModelRepo()
	err := repo.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
		return
	}
	sm, err := factory.GetRepoFactory(r).CreateShipModelRepo().FindById(s.ShipModelId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := prepareShipDetailsResp(s, sm)
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func RetrieveShip(w http.ResponseWriter, r *http.Request) {
	repo := factory.GetRepoFactory(r).CreateShipRepo()
	var s model.Ship
	id := util.ParseIdFromRequest(r)
	s, err := repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sm, err := factory.GetRepoFactory(r).CreateShipModelRepo().FindById(s.ShipModelId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := prepareShipDetailsResp(s, sm)
	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateShip(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateShipRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreateShipRepo()
	var s model.Ship
	s.Id = id
	s.Name = req.Name
	s.SerialNumber = req.SerialNumber
	s.ShipModelId = req.ShipModelId
	err = repo.Update(&s)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteShip(w http.ResponseWriter, r *http.Request) {
	repo := factory.GetRepoFactory(r).CreateShipRepo()
	id := util.ParseIdFromRequest(r)
	err := repo.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func prepareShipDetailsResp(s model.Ship, sm model.ShipModel) dto.ShipDetailsResponse {
	return dto.ShipDetailsResponse{
		Id:           s.Id,
		Name:         s.Name,
		SerialNumber: s.SerialNumber,
		ShipModel:    sm,
	}
}
