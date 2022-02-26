package main

import (
	"github.com/gorilla/mux"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/underland/clients"
	"github.com/underland/handlers"
	"log"
	"net/http"
)

func main() {

	solanaClient := clients.NewSolanaClient(rpc.MainnetRPCEndpoint)

	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.HealthCheckHandler)
	r.HandleFunc("/token-metadata", handlers.TokenMetaDataHandler(solanaClient)).Methods("GET")
	r.HandleFunc("/increment-social-index", handlers.IncrementSocialIndexHandler(solanaClient)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
