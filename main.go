package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/underland/clients"
	"github.com/underland/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	solanaClient := clients.NewSolanaClient(rpc.DevnetRPCEndpoint)

	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.HealthCheckHandler)
	r.HandleFunc("/token-metadata", handlers.TokenMetaDataHandler(solanaClient)).Methods("GET")
	r.HandleFunc("/increment-social-index", handlers.IncrementSocialIndexHandler(solanaClient)).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
