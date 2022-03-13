package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/underland/clients"
	"github.com/underland/services"
	"net/http"
	"strings"
)

func TokenMetaDataHandler(sClient *clients.SolanaClient, aClient *clients.ArweaveClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pubKey := r.URL.Query().Get("pubKey")
		metadataAccount, err := services.GetMetadataAccount(pubKey)
		if err != nil {
			fmt.Printf("Failed to fetch solana metadata account %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		solanaMetadata, err := sClient.GetTokenMetaData(metadataAccount)
		if err != nil {
			fmt.Printf("Failed to fetch solana metadata %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		arweaveDataURI := strings.Split(solanaMetadata.Data.Uri, "/")
		arweaveTxId := arweaveDataURI[len(arweaveDataURI)-1]

		arweaveMetadata, err := aClient.GetMetaData(arweaveTxId)
		if err != nil {
			fmt.Printf("Failed to fetch arweave metadata %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(arweaveMetadata)
	}
}
