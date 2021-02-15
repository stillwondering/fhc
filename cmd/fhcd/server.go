package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	router *chi.Mux
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "welcome")
	}
}
