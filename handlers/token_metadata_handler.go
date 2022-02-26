package handlers

import (
	"encoding/json"
	"github.com/underland/clients"
	"net/http"
)

func TokenMetaDataHandler(client *clients.SolanaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pubKey := r.URL.Query().Get("pubKey")
		metadata, err := client.GetTokenMetaData(pubKey)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(metadata)
	}
}
