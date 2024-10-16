package handler

import (
	"cruiseapp/dto"
	"cruiseapp/repository"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type BadRequestError struct {
	InvalidField string
	Details      string
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("Bad request (invalid field: %s, details: %s)", e.InvalidField, e.Details)
}

func (e *BadRequestError) WithDetails(details string) BadRequestError {
	return BadRequestError{
		InvalidField: e.InvalidField,
		Details:      details,
	}
}

func NewBadRequestError(invalidField string) error {
	return &BadRequestError{
		InvalidField: invalidField,
	}
}

func HandleError(err error, w http.ResponseWriter) {
	var (
		errResp            dto.ErrorResponse
		errNotFound        *repository.NotFoundError
		errActionForbidden *repository.ForbiddenActionError
		errBadRequest      *BadRequestError
	)

	if errors.As(err, &errNotFound) {
		w.WriteHeader(http.StatusNotFound)
		errResp.Msg = errNotFound.Error()
		_ = json.NewEncoder(w).Encode(&errResp)
		return
	}
	if errors.As(err, &errActionForbidden) {
		w.WriteHeader(http.StatusForbidden)
		errResp.Msg = errActionForbidden.Error()
		errResp.Details = errActionForbidden.Details
		_ = json.NewEncoder(w).Encode(&errResp)
		return
	}
	if errors.As(err, &errBadRequest) {
		w.WriteHeader(http.StatusBadRequest)
		errResp.Msg = "Bad request"
		errResp.Details = errBadRequest.InvalidField
		_ = json.NewEncoder(w).Encode(&errResp)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
