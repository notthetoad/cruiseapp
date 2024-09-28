package repository

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "Not found in database"
}

type ForbiddenActionError struct{}

func (e *ForbiddenActionError) Error() string {
	return "Forbidden action"
}
