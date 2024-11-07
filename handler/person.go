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

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePersonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	repo := factory.GetRepoFactory(r).CreatePersonRepo()
	var p model.Person
	p.FirstName = req.FirstName
	p.LastName = req.LastName
	p.Email = req.Email
	p.Phone = req.Phone
	err = repo.Save(&p)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&p)
	ws.SendCreatedMsg(r, "person", p.Id)
}

func RetrievePerson(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	var p dto.PersonDetailsResponse
	repo := factory.GetRepoFactory(r).CreatePersonRepo()
	resp, err := repo.FindById(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	p = preparePersonDetails(*resp)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&p)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdatePersonRequest
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreatePersonRepo()
	p, err := repo.FindById(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	p.Id = id
	p.FirstName = req.FirstName
	p.LastName = req.LastName
	p.Email = req.Email
	p.Phone = req.Phone
	err = repo.Update(p)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	ws.SendUpdatedMsg(r, "person", p.Id)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	err := factory.GetRepoFactory(r).CreatePersonRepo().Delete(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	ws.SendDeletedMsg(r, "person", id)
}

func preparePersonDetails(p model.Person) dto.PersonDetailsResponse {
	return dto.PersonDetailsResponse{
		Id:        p.Id,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Phone:     p.Phone,
	}
}
