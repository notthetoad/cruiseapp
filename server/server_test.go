package server

import (
	"cruiseapp/dto"
	"cruiseapp/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
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

var currId int64
var want = "fooland"

func TestCreatePortIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping create integration test")
	}
	ts := createTestServer()
	defer ts.Close()

	req := dto.CreatePortRequest{
		Location: want,
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

	if res.StatusCode != 201 {
		t.Errorf("Expected 201, Got %d", res.StatusCode)
	}
	if !(port.Id > 0) {
		t.Errorf("Id cannot be 0, Got: %d", port.Id)
	}
	if port.Location != want {
		t.Errorf("Expected %s, Got %s", want, port.Location)
	}
	currId = port.Id
}

func TestRetrievePortIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping retrieve integration test")
	}

	ts := createTestServer()
	defer ts.Close()

	id := strconv.Itoa(int(currId))
	res, err := http.Get(ts.URL + "/port/" + id)
	if err != nil {
		t.Error(err)
	}
	var p model.Port
	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if p.Id != currId {
		t.Errorf("Incorrect id. Want %d, Got %d", currId, p.Id)
	}

	if p.Location != want {
		t.Errorf("Retrieved port incorrect. Want %s, Got %s", want, p.Location)
	}
}

func TestUpdatePortIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping update integration test")
	}

	ts := createTestServer()
	defer ts.Close()

	body := dto.CreatePortRequest{
		Location: "updated",
	}
	bodyReader, err := prepBody(body)
	if err != nil {
		t.Error(err)
	}

	id := strconv.Itoa(int(currId))
	req := httptest.NewRequest(http.MethodPut, ts.URL+"/port/"+id, bodyReader)
	// TODO finish test using the same methodology as in delete test

}

func TestDeletePortIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping delete integration test")
	}

	ts := createTestServer()
	defer ts.Close()

	id := strconv.Itoa(int(currId))
	req, err := http.NewRequest(http.MethodDelete, ts.URL+"/port/"+id, nil)
	if err != nil {
		t.Error(err)
	}
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
	// TODO finish asserts
}
