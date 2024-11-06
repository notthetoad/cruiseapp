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

const (
	ERR_STATUS_CODE = "Invalid status code"
	ERR_LOCATION    = "Invalid Location field"
	ERR_ID          = "Invalid id"
	WANT            = "fooland"
	UPDATED         = "updated"
)

func skipIntegragionTestMsg(method string) string {
	return fmt.Sprintf("Skipping %s integration test", method)
}

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

func TestCreatePortIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip(skipIntegragionTestMsg("CREATE"))
	}
	ts := createTestServer()
	defer ts.Close()

	req := dto.CreatePortRequest{
		Location: WANT,
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

	if res.StatusCode != http.StatusCreated {
		t.Errorf("%s. Want %d, Got %d", ERR_STATUS_CODE, http.StatusCreated, res.StatusCode)
	}
	if !(port.Id > 0) {
		t.Errorf("Id cannot be 0, Got: %d", port.Id)
	}
	if port.Location != WANT {
		t.Errorf("%s. Want %s, Got %s", ERR_LOCATION, WANT, port.Location)
	}
	currId = port.Id
}

func TestRetrievePortIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip(skipIntegragionTestMsg("RETRIEVE"))
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
		t.Errorf("%s. Want %d, Got %d", ERR_ID, currId, p.Id)
	}

	if p.Location != WANT {
		t.Errorf("%s. Want %s, Got %s", ERR_LOCATION, WANT, p.Location)
	}
}

func TestUpdatePortIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip(skipIntegragionTestMsg("UPDATE"))
	}

	tests := []struct {
		loc     string
		paramId string
		want    int
	}{
		{"newLocation", "10", 204},
		{"newLocation", "999", 404},
	}

	ts := createTestServer()
	defer ts.Close()

	var c http.Client

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt)
		t.Run(testname, func(t *testing.T) {
			body := dto.CreatePortRequest{
				Location: tt.loc,
			}

			bodyReader, err := prepBody(body)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(http.MethodPut, ts.URL+"/port/"+tt.paramId, bodyReader)
			if err != nil {
				t.Fatalf("cannot prepare http request: %v", err)
			}

			res, err := c.Do(req)
			if err != nil {
				t.Fatalf("client unable to send request: %v", err)
			}

			if res.StatusCode != tt.want {
				t.Errorf("%s; want %d, got %d", ERR_STATUS_CODE, tt.want, res.StatusCode)
			}
		})

	}

	body := dto.CreatePortRequest{
		Location: UPDATED,
	}
	bodyReader, err := prepBody(body)
	if err != nil {
		t.Error(err)
	}

	id := strconv.Itoa(int(currId))
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/port/"+id, bodyReader)
	if err != nil {
		t.Error(err)
	}
	c = http.Client{}
	res, err := c.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 204 {
		t.Errorf("%s. Want %d, Got %d", ERR_STATUS_CODE, http.StatusNoContent, res.StatusCode)
	}

	var port model.Port
	resp, err := http.Get(ts.URL + "/port/" + id)
	if err != nil {
		t.Error(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&port)
	if err != nil {
		t.Error(err)
	}

	if port.Id != currId {
		t.Errorf("%s. Want %d, Got %d", ERR_ID, currId, port.Id)
	}

	if port.Location != UPDATED {
		t.Errorf("%s. Want %s, Got %s", ERR_LOCATION, UPDATED, port.Location)
	}

}

func TestDeletePortIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip(skipIntegragionTestMsg("DELETE"))
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

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("%s. Want %d, Got %d", ERR_STATUS_CODE, http.StatusNoContent, res.StatusCode)
	}
}
