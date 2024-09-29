package dto

type CreateCrewRankRequest struct {
	Name string
}

type UpdateCrewRankRequest CreateCrewRankRequest

type CreateCrewMemberRequest struct {
	CrewRankId int64
	PersonId   int64
}

type UpdateCrewMemberRequest CreateCrewMemberRequest
