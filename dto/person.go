package dto

type CreatePersonRequest struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UpdatePersonRequest CreatePersonRequest

type PersonDetailsResponse struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
