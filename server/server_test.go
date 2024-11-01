package server

import (
	"cruiseapp/dto"
	"cruiseapp/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createTestServer() *httptest.Server {
	return httptest.NewServer(wrapMiddleware(newRouter()))
}

func prepBody(v any) (io.Reader, error) {
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(&v)
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(sb.String())

	return reader, nil

}

func TestCreatePortIntegration(t *testing.T) {
	expected := "fooland"
	ts := createTestServer()
	_ = ts
	defer ts.Close()

	req := dto.CreatePortRequest{
		Location: expected,
	}
	bodyReader, err := prepBody(req)
	if err != nil {
		t.Error(err)
	}

	res, err := http.Post(ts.URL+"/port", "application/json", bodyReader)
	if err != nil {
		t.Error(err)
	}

	var port model.Port
	_ = json.NewDecoder(res.Body).Decode(&port)
	defer res.Body.Close()

	fmt.Println(port)
	if res.StatusCode != 201 {
		t.Errorf("Expected 201, Got %d", res.StatusCode)
	}
	if !(port.Id > 0) {
		t.Errorf("Id cannot be 0, Got: %d", port.Id)
	}
	if port.Location != expected {
		t.Errorf("Expected %s, Got %s", expected, port.Location)
	}
}
