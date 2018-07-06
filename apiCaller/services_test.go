package apiCaller

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"errors"
	"fmt"
)

func TestApiCaller(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, `{
  "result": "ok",
  "data": {
    "user_id": 1,
    "rating": 42
  }
}`)
	}))
	defer ts.Close()

	user := User{id: 1}
	err := ApiCaller(&user, ts.URL)
	if err != nil {
		t.Error("Ecspect no error", err)
	}

	if user.rating != 42 {
		t.Error("User rating has not updated", err)
	}
}

func TestApiCallerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, `Code 500`)
	}))
	defer ts.Close()

	user := User{id: 1}
	err := ApiCaller(&user, ts.URL)
	if err == nil {
		t.Error("Ecspect no error", err)
	}

}

func TestApiCallerNoServer(t *testing.T) {
	user := User{id: 1}
	err := ApiCaller(&user, "/")
	if err == nil {
		t.Error("Ecspect no error", err)
	}

}

func TestAdiCallerErrorResponse(t *testing.T) {

	user := User{id: 1}
	err := updateUser(&user, errReader(0))
	if err == nil {
		t.Error("Ecspect error", err)
	}

}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}
func (errReader) Close() error {
	return nil
}
