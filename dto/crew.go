package dto

import "cruiseapp/model"

type CreateCrewRankRequest struct {
	Name string
}

type UpdateCrewRankRequest CreateCrewRankRequest

type CreateCrewMemberRequest struct {
	CrewRankId int64
	PersonId   int64
}

type UpdateCrewMemberRequest CreateCrewMemberRequest

// TODO add Person field
type CrewMemberResponse struct {
	Id       int64
	CrewRank model.CrewRank
}
