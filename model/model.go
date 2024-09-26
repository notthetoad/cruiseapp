package model

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
