package dto

type CreatePortRequest struct {
	Location string `json:"location"`
}

type PortDetailsResponse struct {
	Id       int64
	Location string
}
