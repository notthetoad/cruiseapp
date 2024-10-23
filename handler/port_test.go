package handler

import (
	"fmt"
	"io"
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

	out, _ := io.ReadAll(rr.Body)
	fmt.Println(string(out))
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
