package handler

import (
	"cruiseapp/dto"
	"cruiseapp/handler/util"
	"cruiseapp/model"
	"cruiseapp/repository"
	"cruiseapp/repository/factory"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func CreateCruise(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCruiseRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	var c model.Cruise
	portRepo := factory.GetRepoFactory(r).CreatePortRepo()
	fromLocation, err := portRepo.FindById(req.FromLocationId)
	if err != nil {
		validateFK(&err, "FromLocation")
		HandleError(err, w)
		return
	}
	c.FromLocation = *fromLocation
	toLocation, err := portRepo.FindById(req.ToLocationId)
	if err != nil {
		validateFK(&err, "ToLocation")
		HandleError(err, w)
		return
	}
	c.ToLocation = *toLocation

	crewRepo := factory.GetRepoFactory(r).CreateCrewMemberRepo()
	var crew []*model.CrewMember
	crew, err = crewRepo.FindAllByIds(req.CrewMembersIds)
	c.Crew = crew
	fmt.Println(crew)

	repo := factory.GetRepoFactory(r).CreateCruiseRepo()
	c.StartDate = req.StartDate
	c.EndDate = req.EndDate
	err = repo.Save(&c)
	fmt.Println(err)
	_ = json.NewEncoder(w).Encode(&c)
}

func RetrieveCruise(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreateCruiseRepo()
	resp, err := repo.FindById(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&resp)
}

func validateFK(err *error, fieldName string) {
	var notFoundErr *repository.NotFoundError
	if errors.As(*err, &notFoundErr) {
		*err = NewBadRequestError(fieldName)
	}
}
