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

// TODO refactor
// TODO bad request error response
func HandleError(err error, w http.ResponseWriter) {
	var errResp dto.ErrorResponse
	var notFoundErr *repository.NotFoundError
	var actionForbiddenErr *repository.ForbiddenActionError
	var badRequestErr *BadRequestError
	if errors.As(err, &notFoundErr) {
		w.WriteHeader(http.StatusNotFound)
		errResp.Msg = notFoundErr.Error()
		_ = json.NewEncoder(w).Encode(&errResp)
	} else if errors.As(err, &actionForbiddenErr) {
		w.WriteHeader(http.StatusForbidden)
		errResp.Msg = actionForbiddenErr.Error()
		errResp.Details = actionForbiddenErr.Details
		_ = json.NewEncoder(w).Encode(&errResp)
	} else if errors.As(err, &badRequestErr) {
		w.WriteHeader(http.StatusBadRequest)
		errResp.Msg = "Bad request"
		errResp.Details = badRequestErr.InvalidField
		_ = json.NewEncoder(w).Encode(&errResp)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
