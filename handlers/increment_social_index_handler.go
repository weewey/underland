package handlers

import (
	"encoding/json"
	"github.com/underland/clients"
	"net/http"
)

type IncrementSocialIndexRequest struct {
	PubKey string
}

func IncrementSocialIndexHandler(client *clients.SolanaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody IncrementSocialIndexRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		metadata, err := client.GetTokenMetaData(reqBody.PubKey)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Get the social index from the metadata
		// increment it by 1
		// send an update to increment it
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(metadata)
	}
}
