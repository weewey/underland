package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/portto/solana-go-sdk/program/metaplex/tokenmeta"
	"github.com/underland/clients"
	"github.com/underland/models"
	"github.com/underland/services"
	"net/http"
	"strings"
)

type IncrementSocialIndexRequest struct {
	PubKey            string
	IncrementSocialBy int
}

func IncrementSocialHandler(sClient *clients.SolanaClient, aClient *clients.ArweaveClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody IncrementSocialIndexRequest
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

		arweaveDataURI := strings.Split(solanaMetadata.Data.Uri, "/")
		arweaveTxId := arweaveDataURI[len(arweaveDataURI)-1]

		arweaveMetadata, err := aClient.GetMetaData(arweaveTxId)
		if err != nil {
			fmt.Printf("Failed to fetch arweave metadata %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		IncrementSocial(reqBody.IncrementSocialBy, &arweaveMetadata)

		txn, err := aClient.UploadMetadata(&arweaveMetadata)
		if err != nil {
			fmt.Printf("Failed to upload arweave metadata %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		UpdateMetadataURI(&solanaMetadata.Data, fmt.Sprintf("https://arweave.net/%s", txn.ID))
		txId, err := sClient.UpdateTokenMetaData(metadataAccount, &solanaMetadata.Data)
		if err != nil {
			fmt.Printf("Failed to update solana metadata %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"transactionId": txId})
	}
}

func UpdateMetadataURI(tokenMetaData *tokenmeta.Data, dataURI string) *tokenmeta.Data {
	tokenMetaData.Uri = dataURI
	return tokenMetaData
}

func IncrementSocial(incrementCount int, metadata *models.CharacterMetaData) *models.CharacterMetaData {
	var scores []int
	for _, attribute := range metadata.Attributes {
		if attribute.TraitType == models.SOCIAL_INDEX {
			continue
		}
		if attribute.TraitType == models.SOCIAL {
			attribute.Value += incrementCount
		}
		scores = append(scores, attribute.Value)
	}

	var socialIndex = 0
	for _, score := range scores {
		socialIndex += 20 * score
	}

	metadata.SocialIndex = socialIndex
	for _, attribute := range metadata.Attributes {
		if attribute.TraitType == models.SOCIAL_INDEX {
			attribute.Value = socialIndex
		}
	}
	return metadata
}
