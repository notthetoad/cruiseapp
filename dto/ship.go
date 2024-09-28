package dto

import "cruiseapp/model"

type CreateShipModelRequest struct {
	Name string
}

type UpdateShipModelRequest CreateShipModelRequest

type DeleteShipModelRequest int64

type CreateShipRequest struct {
	Name         string
	SerialNumber string
	ShipModelId  int64
}

type ShipDetailsResponse struct {
	Id           int64
	Name         string
	SerialNumber string
	ShipModel    model.ShipModel
}

type UpdateShipRequest CreateShipRequest

type DeleteShipRequest int64
