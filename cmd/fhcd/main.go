package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	exitOk   = 0
	exitErr  = 1
	certFile = "./dist/127.0.0.1/cert.pem"
	keyFile  = "./dist/127.0.0.1/key.pem"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitErr)
	}
}

func run() error {
	r := chi.NewRouter()
	r.Use(middleware.BasicAuth("restricted area", map[string]string{
		"patrick": "patrick",
	}))
	r.Use(middleware.Logger)

	srv := &server{
		router: r,
	}
	srv.routes()

	return http.ListenAndServeTLS(":8080", certFile, keyFile, srv)
}
