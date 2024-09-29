package dto

type CreatePersonRequest struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UpdatePersonRequest CreatePersonRequest

type PersonResponse struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
