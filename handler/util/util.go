package util

import (
	"log"
	"net/http"
	"strconv"
)

func ParseIdFromRequest(r *http.Request) int64 {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error parsing id from request: %v", err)
		return -1
	}

	return int64(id)
}
