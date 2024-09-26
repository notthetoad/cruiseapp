package repository

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "Not found in database"
}
