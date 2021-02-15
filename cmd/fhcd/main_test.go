package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestHandleIndex(t *testing.T) {
	srv := &server{
		router: chi.NewRouter(),
	}
	srv.routes()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("status should be ok, got %d", w.Code)
	}

	expected := []byte("welcome")
	actual := w.Body.Bytes()
	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("expected body %s, got %s", string(expected), string(actual))
	}
}
