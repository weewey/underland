package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/underland/clients"
	"github.com/underland/services"
	"net/http"
)

type UpdateMetadataDataURIRequest struct {
	PubKey  string `json:"pubKey"`
	DataURI string `json:"dataURI"`
}

func UpdateMetadataHandler(sClient *clients.SolanaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody UpdateMetadataDataURIRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		metadataAccount, err := services.GetMetadataAccount(reqBody.PubKey)
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
		solanaMetadata.Data.Uri = reqBody.DataURI
		txnId, err := sClient.UpdateTokenMetaData(metadataAccount, &solanaMetadata.Data)
		if err != nil {
			fmt.Printf("Failed to update token metadata %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"transactionId": txnId})
	}
}
