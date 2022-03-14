package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/underland/clients"
	"github.com/underland/models"
	"net/http"
)

type UploadMetadataRequest struct {
	Metadata models.CharacterMetaData
}

func UploadMetadataHandler(aClient *clients.ArweaveClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody UploadMetadataRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		txn, err := aClient.UploadMetadata(&reqBody.Metadata)
		if err != nil {
			fmt.Printf("Failed to upload arweave metadata %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"transactionId": txn.ID,
			"dataURL":       fmt.Sprintf("https://arweave.net/%v", txn.ID)})
	}
}
