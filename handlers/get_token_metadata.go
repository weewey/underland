package handlers

import (
	"encoding/json"
	"github.com/underland/clients"
	"net/http"
)

func GetTokenMetaData(client *clients.SolanaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metadata, err := client.GetTokenMetaData("GphF2vTuzhwhLWBWWvD8y5QLCPp1aQC5EnzrWsnbiWPx")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		respBody, err := json.Marshal(metadata)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respBody)
	}
}
