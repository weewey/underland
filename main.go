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
	arweavePvtKey := os.Getenv("ARWEAVE_PVT_KEY")
	arweaveClient := clients.NewArweaveClient(arweavePvtKey, "https://arweave.net")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	solanaPvtKey := os.Getenv("SOLANA_PVT_KEY")
	solanaClient := clients.NewSolanaClient(rpc.DevnetRPCEndpoint, solanaPvtKey)

	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.HealthCheckHandler)

	r.HandleFunc("/token-metadata",
		handlers.TokenMetaDataHandler(solanaClient, arweaveClient)).Methods("GET")

	r.HandleFunc("/increment-social",
		handlers.IncrementSocialHandler(solanaClient, arweaveClient)).Methods("POST")

	r.HandleFunc("/upload-metadata",
		handlers.UploadMetadataHandler(arweaveClient)).Methods("POST")

	r.HandleFunc("/update-metadata",
		handlers.UpdateMetadataHandler(solanaClient)).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
