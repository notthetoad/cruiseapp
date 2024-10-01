package handler

import (
	"cruiseapp/dto"
	"cruiseapp/handler/util"
	"cruiseapp/model"
	"cruiseapp/repository/factory"
	"encoding/json"
	"log"
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&p)

}

func RetrievePerson(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	var p dto.PersonResponse
	repo := factory.GetRepoFactory(r).CreatePersonRepo()
	resp, err := repo.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
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
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	p.Id = id
	p.FirstName = req.FirstName
	p.LastName = req.LastName
	p.Email = req.Email
	p.Phone = req.Phone
	repo.Update(p)
	w.WriteHeader(http.StatusNoContent)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	err := factory.GetRepoFactory(r).CreatePersonRepo().Delete(id)
	if err != nil {
		// log.Println(err)
		// w.WriteHeader(http.StatusInternalServerError)
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func preparePersonDetails(p model.Person) dto.PersonResponse {
	return dto.PersonResponse{
		Id:        p.Id,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Phone:     p.Phone,
	}
}
