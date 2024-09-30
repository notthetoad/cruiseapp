package dto

type CreatePortRequest struct {
	Location string `json:"location"`
}

type PortResponse struct {
	Id       int64
	Location string
}
