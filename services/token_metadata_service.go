package services

import (
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/tokenmeta"
	"log"
)

func GetMetadataAccount(pubKey string) (common.PublicKey, error) {
	mint := common.PublicKeyFromString(pubKey)

	metadataAccount, err := tokenmeta.GetTokenMetaPubkey(mint)
	if err != nil {
		log.Printf("failed to get metadata account, err: %v", err)
		return common.PublicKey{}, err
	}
	return metadataAccount, nil
}
