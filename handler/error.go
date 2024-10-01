package handler

import (
	"cruiseapp/dto"
	"cruiseapp/repository"
	"encoding/json"
	"errors"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter) {
	var errResp dto.ErrorResponse
	var notFoundErr *repository.NotFoundError
	var actionForbiddenErr *repository.ForbiddenActionError
	if errors.As(err, &notFoundErr) {
		w.WriteHeader(http.StatusNotFound)
		errResp.Msg = notFoundErr.Error()
		_ = json.NewEncoder(w).Encode(&errResp)
	} else if errors.As(err, &actionForbiddenErr) {
		w.WriteHeader(http.StatusForbidden)
		errResp.Msg = actionForbiddenErr.Error()
		errResp.Details = actionForbiddenErr.Details
		_ = json.NewEncoder(w).Encode(&errResp)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
