package main

import (
	"github.com/gorilla/mux"
	"github.com/underland/handlers"
)

func initializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.HealthCheckHandler)
	r.HandleFunc("/increment-social-index", handlers.IncrementSocialIndexHandler)
	return r
}
