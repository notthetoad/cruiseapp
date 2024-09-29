package handler

import (
	"cruiseapp/dto"
	"cruiseapp/model"
	"cruiseapp/repository/factory"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateCrewRank(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCrewRankRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	crf := factory.GetRepoFactory(r).CreateCrewRankRepo()
	cr, err = crf.FindById(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusFound)
	_ = json.NewEncoder(w).Encode(&cr)
}

func UpdateCrewRank(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCrewRankRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	repo := factory.GetRepoFactory(r).CreateCrewRankRepo()
	var cr model.CrewRank
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cr.Id = int64(id)
	cr.Name = req.Name
	err = repo.Update(&cr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteCrewRank(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	crf := factory.GetRepoFactory(r).CreateCrewRankRepo()
	err = crf.Delete(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func CreateCrewMember(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCrewMemberRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	repo := factory.GetRepoFactory(r).CreateCrewMemberRepo()
	var cm model.CrewMember
	cm.CrewRankId = req.CrewRankId
	cm.PersonId = req.PersonId
	err = repo.Save(&cm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusFound)
	_ = json.NewEncoder(w).Encode(&cm)
}

func RetrieveCrewMember(w http.ResponseWriter, r *http.Request) {}

func UpdateCrewMember(w http.ResponseWriter, r *http.Request) {}

func DeleteCrewMember(w http.ResponseWriter, r *http.Request) {}
