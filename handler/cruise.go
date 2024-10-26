package handler

import (
	"cruiseapp/dto"
	"cruiseapp/handler/util"
	"cruiseapp/model"
	"cruiseapp/repository"
	"cruiseapp/repository/factory"
	"encoding/json"
	"errors"
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
	if err != nil {
		HandleError(err, w)
		return
	}
	c.Crew = crew

	personRepo := factory.GetRepoFactory(r).CreatePersonRepo()
	var passengers []*model.Person
	passengers, err = personRepo.FindAllByIds(req.PassengersIds)
	if err != nil {
		HandleError(err, w)
		return
	}
	c.Passengers = passengers

	repo := factory.GetRepoFactory(r).CreateCruiseRepo()
	c.StartDate = req.StartDate
	c.EndDate = req.EndDate
	err = repo.Save(&c)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&c)
}

func RetrieveCruise(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreateCruiseRepo()
	c, err := repo.FindById(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := prepareCruiseDetailsResponse(*c)
	_ = json.NewEncoder(w).Encode(&resp)
}

func UpdateCruise(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCruiseRequest
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreateCruiseRepo()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		HandleError(err, w)
		return
	}
	var c model.Cruise
	c.Id = id
	c.StartDate = req.StartDate
	c.EndDate = req.EndDate
	c.FromLocation.Id = req.FromLocationId
	c.ToLocation.Id = req.ToLocationId
	err = repo.Update(&c)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteCruise(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreateCruiseRepo()
	err := repo.Delete(id)
	if err != nil {
		validateFK(&err, "cruise crew member")
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func prepareCruiseDetailsResponse(c model.Cruise) dto.CruiseDetailsResponse {
	duration := c.EndDate.Sub(c.StartDate)

	return dto.CruiseDetailsResponse{
		Id:           c.Id,
		StartDate:    c.StartDate,
		EndDate:      c.EndDate,
		Duration:     duration.String(),
		FromLocation: c.FromLocation,
		ToLocation:   c.ToLocation,
		CrewMembers:  c.Crew,
	}
}

func validateFK(err *error, fieldName string) {
	var notFoundErr *repository.NotFoundError
	if errors.As(*err, &notFoundErr) {
		*err = NewBadRequestError(fieldName)
	}
}
