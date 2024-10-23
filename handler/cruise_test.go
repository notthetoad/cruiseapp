package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetCruise(t *testing.T) {
	rr, ctx := setupRecorderAndCtx()
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/cruise/1", nil)
	req.SetPathValue("id", "1")

	RetrieveCruise(rr, req)

	res := rr.Result()

	if res.StatusCode != 200 {
		t.Errorf("Expected 200, Got %d", res.StatusCode)
	}
}

func TestPostCruise(t *testing.T) {
	rr, ctx := setupRecorderAndCtx()
	body := `
{
    "EndDate": "2025-01-01T00:00:00Z",
    "StartDate": "2024-01-01T00:00:00Z",
    "CrewMembers": [
        16,
        28
    ],
    "FromLocation": 2,
    "ToLocation": 2
}`
	reader := strings.NewReader(body)
	req := httptest.NewRequestWithContext(ctx, http.MethodPost, "/cruise", reader)

	CreateCruise(rr, req)

	res := rr.Result()

	if res.StatusCode != 201 {
		t.Errorf("Expected 201, Got %d", res.StatusCode)
	}
}

func TestUpdateCruise(t *testing.T) {
	rr, ctx := setupRecorderAndCtx()
	body := `
{
    "EndDate": "2025-01-01T00:00:00Z",
    "StartDate": "2024-01-01T00:00:00Z",
    "CrewMembers": [
        16,
        28
    ],
    "FromLocation": 4,
    "ToLocation": 6
}`
	reader := strings.NewReader(body)
	req := httptest.NewRequestWithContext(ctx, http.MethodPut, "/cruise/1", reader)
	req.SetPathValue("id", "1")

	UpdateCruise(rr, req)

	res := rr.Result()

	if res.StatusCode != 204 {
		t.Errorf("Expected 204, Got %d", res.StatusCode)
	}
}
