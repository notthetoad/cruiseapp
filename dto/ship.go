package dto

type CreateShipModelRequest struct {
	Name string
}

type CreateShipRequest struct {
	Name         string
	SerialNumber string
	ShipModelId  int64
}
