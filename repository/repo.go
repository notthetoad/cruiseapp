package repository

import "fmt"

type NotFoundError struct {
	Id int64
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Object with id %d not found", e.Id)
}

func NewNotFoundError(id int64) *NotFoundError {
	return &NotFoundError{
		Id: id,
	}
}

type ForbiddenActionError struct {
	Id      int64
	Action  string
	Details string
}

func (e *ForbiddenActionError) Error() string {
	return fmt.Sprintf("Action forbidden: trying to %s object with id: %d", e.Action, e.Id)
	// return fmt.Sprintf("Action forbidden: trying to %s object with id %d %s", e.Action, e.Id, e.Details)
}

func (e *ForbiddenActionError) WithDetails(details string) *ForbiddenActionError {
	return &ForbiddenActionError{
		Id:      e.Id,
		Action:  e.Action,
		Details: details,
	}
}

func NewForbiddenActionError(id int64, action string) *ForbiddenActionError {
	return &ForbiddenActionError{
		Id:     id,
		Action: action,
	}
}
