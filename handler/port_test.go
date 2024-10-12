package handler

import (
	"context"
	"cruiseapp/repository/factory"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
	foo := r.Context().Value("foo")
	fmt.Println(foo)
	fmt.Println("path val =", r.PathValue("id"))
}

func ctxWithMockRepoFactory(ctx context.Context) context.Context {
	var repoFactory factory.RepoFactory = MockRepoFactory{}
	return factory.CtxWithRepoFactory(ctx, repoFactory)
}

func TestGetPort(t *testing.T) {
	rr := httptest.NewRecorder()
	ctx := ctxWithMockRepoFactory(context.Background())
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/port/1", nil)
	req.SetPathValue("id", "1")

	RetrievePort(rr, req)

	res := rr.Result()

	if res.StatusCode != 200 {
		t.Errorf("expected 200 got %d", res.StatusCode)
	}

	out, _ := io.ReadAll(rr.Body)
	fmt.Println(string(out))
}