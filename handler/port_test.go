package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPort(t *testing.T) {
	rr, ctx := setupRecorderAndCtx()
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/port/1", nil)
	req.SetPathValue("id", "1")

	RetrievePort(rr, req)
	res := rr.Result()

	if res.StatusCode != 200 {
		t.Errorf("Expected 200, Got %d", res.StatusCode)
	}
}

func TestPostPort(t *testing.T) {
	rr, ctx := setupRecorderAndCtx()
	body := `{"Location": "fooland"}`
	reader := strings.NewReader(body)
	req := httptest.NewRequestWithContext(ctx, http.MethodPost, "/port", reader)

	CreatePort(rr, req)
	res := rr.Result()

	if res.StatusCode != 201 {
		t.Errorf("Expected 201, Got %d", res.StatusCode)
	}
}

func TestUpdatePort(t *testing.T) {
	rr, ctx := setupRecorderAndCtx()
	body := `{"Location": "updatedLocation"}`
	reader := strings.NewReader(body)
	req := httptest.NewRequestWithContext(ctx, http.MethodPut, "/port/1", reader)
	req.SetPathValue("id", "1")

	UpdatePort(rr, req)
	res := rr.Result()

	if res.StatusCode != 204 {
		t.Errorf("Expected 204, Got %d", res.StatusCode)
	}

}

func TestDeletePort(t *testing.T) {
	rr, ctx := setupRecorderAndCtx()
	req := httptest.NewRequestWithContext(ctx, http.MethodDelete, "/port/1", nil)
	req.SetPathValue("id", "1")

	DeletePort(rr, req)
	res := rr.Result()

	if res.StatusCode != 204 {
		t.Errorf("Expected 204, Got %d", res.StatusCode)
	}
}
