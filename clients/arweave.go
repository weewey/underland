package clients

import (
	"encoding/json"
	"fmt"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/underland/models"
	"log"
)

type ArweaveClient struct {
	wallet *goar.Wallet
	client *goar.Client
}

func NewArweaveClient(arweavePvtKey string, arweaveUrl string) *ArweaveClient {
	wallet, err := goar.NewWallet([]byte(arweavePvtKey), arweaveUrl)
	if err != nil {
		log.Println(err.Error())
	}
	client := goar.NewClient(arweaveUrl)
	return &ArweaveClient{
		wallet: wallet,
		client: client,
	}
}

func (a *ArweaveClient) UploadMetadata(data *models.CharacterMetaData) (types.Transaction, error) {
	metadataJson, _ := json.Marshal(data)
	tx, err := a.wallet.SendData(metadataJson, []types.Tag{})
	if err != nil {
		fmt.Printf("Error uploading data to arweave %s", err.Error())
		return types.Transaction{}, err
	}
	return tx, nil
}

func (a ArweaveClient) GetMetaData(tx string) (models.CharacterMetaData, error) {
	txn, err := a.client.GetTransactionDataByGateway(tx)
	if err != nil {
		log.Printf("Error fetching %s metadata: %s", tx, err.Error())
		return models.CharacterMetaData{}, err
	}
	var metadata models.CharacterMetaData

	err = json.Unmarshal(txn, &metadata)
	if err != nil {
		log.Printf("Error unmarshaling %s metadata: %s", tx, err.Error())
		return models.CharacterMetaData{}, err
	}
	return metadata, nil
}
