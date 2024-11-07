package dto

import (
	"cruiseapp/model"
	"time"
)

type CreateCruiseRequest struct {
	Id             int64
	StartDate      time.Time
	EndDate        time.Time
	FromLocationId int64   `json:"FromLocation"`
	ToLocationId   int64   `json:"ToLocation"`
	CrewMembersIds []int64 `json:"CrewMembers"`
	PassengersIds  []int64 `json:"Passengers"`
}

type CruiseDetailsResponse struct {
	Id           int64
	StartDate    time.Time
	EndDate      time.Time
	Duration     string
	FromLocation model.Port
	ToLocation   model.Port
	CrewMembers  []*model.CrewMember
	Passengers   []*model.Person
}
