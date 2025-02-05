package handler

import (
	"cruiseapp/dto"
	"cruiseapp/handler/util"
	"cruiseapp/model"
	"cruiseapp/repository/factory"
	"cruiseapp/ws"
	"encoding/json"
	"net/http"
)

func CreatePort(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePortRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	repo := factory.GetRepoFactory(r).CreatePortRepo()
	var p model.Port
	p.Location = req.Location
	repo.Save(&p)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&p)
	ws.SendCreatedMsg(r, "port", p.Id)
}

func RetrievePort(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreatePortRepo()
	port, err := repo.FindById(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	resp := preparePortResponse(*port)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&resp)

}

func UpdatePort(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePortRequest
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreatePortRepo()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		HandleError(err, w)
		return
	}
	var p model.Port
	p.Id = id
	p.Location = req.Location
	err = repo.Update(&p)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	ws.SendUpdatedMsg(r, "port", p.Id)
}

func DeletePort(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreatePortRepo()
	err := repo.Delete(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	ws.SendDeletedMsg(r, "port", id)
}

func preparePortResponse(p model.Port) dto.PortDetailsResponse {
	return dto.PortDetailsResponse{
		Id:       p.Id,
		Location: p.Location,
	}
}
