package main

import (
	"net/http"
	"time"

	"github.com/juanmiguelar/character_generator/entrypoints"
)

func main() {
	router := entrypoints.CreateRoutes()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}