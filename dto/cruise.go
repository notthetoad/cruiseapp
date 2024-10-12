package dto

import (
	"time"
)

type CreateCruiseRequest struct {
	Id             int64
	StartDate      time.Time
	EndDate        time.Time
	FromLocationId int64   `json:"FromLocation"`
	ToLocationId   int64   `json:"ToLocation"`
	CrewMembersIds []int64 `json:"CrewMembers"`
}
