package model

type Port struct {
	Id       int64
	Location string
}

type CrewRank struct {
	Id   int64
	Name string
}

type CrewMember struct {
	Id         int64
	PersonId   int64
	CrewRankId int64
}
