package model

import "time"

type Port struct {
	Id       int64
	Location string
}

type ShipModel struct {
	Id   int64
	Name string
}

type Ship struct {
	Id           int64
	Name         string
	SerialNumber string
	ShipModelId  int64
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

type Person struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type Cruise struct {
	Id           int64
	StartDate    time.Time
	EndDate      time.Time
	FromLocation Port
	ToLocation   Port
	Crew         []*CrewMember
	Passengers   []*Person
}
