package clients

import (
	"context"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/tokenmeta"
	"log"
)

type SolanaClient struct {
	client *client.Client
}

func NewSolanaClient(rpcUrl string) *SolanaClient {
	return &SolanaClient{client: client.NewClient(rpcUrl)}
}

func (s *SolanaClient) GetTokenMetaData(pubKey string) (tokenmeta.Metadata, error) {
	mint := common.PublicKeyFromString(pubKey)

	metadataAccount, err := tokenmeta.GetTokenMetaPubkey(mint)
	if err != nil {
		log.Printf("failed to get metadata account, err: %v", err)
	}

	accountInfo, err := s.client.GetAccountInfo(context.Background(), metadataAccount.ToBase58())
	if err != nil {
		log.Printf("failed to get accountInfo, err: %v", err)
	}

	metadata, err := tokenmeta.MetadataDeserialize(accountInfo.Data)
	if err != nil {
		log.Printf("failed to parse metaAccount, err: %v", err)
	}

	return metadata, err
}
