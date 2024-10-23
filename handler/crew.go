package handler

import (
	"cruiseapp/dto"
	"cruiseapp/handler/util"
	"cruiseapp/model"
	"cruiseapp/repository/factory"
	"encoding/json"
	"net/http"
)

func CreateCrewRank(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCrewRankRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		HandleError(err, w)
		return
	}
	repo := factory.GetRepoFactory(r).CreateCrewRankRepo()
	var cr model.CrewRank
	cr.Name = req.Name
	repo.Save(&cr)
	w.WriteHeader(http.StatusFound)
	_ = json.NewEncoder(w).Encode(&cr)
}

func RetrieveCrewRank(w http.ResponseWriter, r *http.Request) {
	var cr *model.CrewRank
	id := util.ParseIdFromRequest(r)
	crf := factory.GetRepoFactory(r).CreateCrewRankRepo()
	cr, err := crf.FindById(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusFound)
	_ = json.NewEncoder(w).Encode(&cr)
}

func UpdateCrewRank(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCrewRankRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		HandleError(err, w)
		return
	}
	repo := factory.GetRepoFactory(r).CreateCrewRankRepo()
	var cr model.CrewRank
	id := util.ParseIdFromRequest(r)
	if err != nil {
		HandleError(err, w)
		return
	}
	cr.Id = id
	cr.Name = req.Name
	err = repo.Update(&cr)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteCrewRank(w http.ResponseWriter, r *http.Request) {
	id := util.ParseIdFromRequest(r)
	crf := factory.GetRepoFactory(r).CreateCrewRankRepo()
	err := crf.Delete(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func CreateCrewMember(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCrewMemberRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		HandleError(err, w)
		return
	}
	repo := factory.GetRepoFactory(r).CreateCrewMemberRepo()
	var cm model.CrewMember
	cm.CrewRankId = req.CrewRankId
	cm.PersonId = req.PersonId
	err = repo.Save(&cm)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusFound)
	_ = json.NewEncoder(w).Encode(&cm)
}

func RetrieveCrewMember(w http.ResponseWriter, r *http.Request) {
	repo := factory.GetRepoFactory(r).CreateCrewMemberRepo()
	id := util.ParseIdFromRequest(r)
	cm, err := repo.FindById(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	cr, err := factory.GetRepoFactory(r).CreateCrewRankRepo().FindById(cm.CrewRankId)
	if err != nil {
		HandleError(err, w)
		return
	}
	p, err := factory.GetRepoFactory(r).CreatePersonRepo().FindById(cm.PersonId)
	resp := prepareCrewMemberDetailsResp(*cm, *cr, *p)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		HandleError(err, w)
	}
}

func UpdateCrewMember(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCrewMemberRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		HandleError(err, w)
		return
	}
	id := util.ParseIdFromRequest(r)
	repo := factory.GetRepoFactory(r).CreateCrewMemberRepo()
	var cm model.CrewMember
	cm.Id = id
	cm.CrewRankId = req.CrewRankId
	cm.PersonId = req.PersonId
	err = repo.Update(&cm)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteCrewMember(w http.ResponseWriter, r *http.Request) {
	repo := factory.GetRepoFactory(r).CreateCrewMemberRepo()
	id := util.ParseIdFromRequest(r)
	err := repo.Delete(id)
	if err != nil {
		HandleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// TODO add Person field
func prepareCrewMemberDetailsResp(cm model.CrewMember, cr model.CrewRank, p model.Person) dto.CrewMemberResponse {
	return dto.CrewMemberResponse{
		Id:       cm.Id,
		CrewRank: cr,
		Person:   p,
	}
}
