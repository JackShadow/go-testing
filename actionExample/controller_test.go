package actionExample

import (
	"testing"
	"net/http/httptest"
	"encoding/json"
)

func TestUserHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "http://127.0.0.1:80/user?id=42", nil)
	w := httptest.NewRecorder()
	userHandler(w, r)
	user := User{}
	json.Unmarshal([]byte (w.Body.String()), &user)
	if user.Id != 42 {
		t.Errorf("Invalid user id %d expected %d", user.Id, 42)
	}

	r = httptest.NewRequest("GET", "http://127.0.0.1:80/user", nil)
	w = httptest.NewRecorder()
	userHandler(w, r)
	if w.Body.String()!="Error"{
		t.Errorf("Expected error got %s", w.Body.String())
	}

}
