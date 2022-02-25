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
	r.HandleFunc("/get-token-metadata", handlers.GetTokenMetaData(solanaClient))

	log.Fatal(http.ListenAndServe(":8080", r))
}
