package clients

import (
	"context"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/tokenmeta"
	"github.com/portto/solana-go-sdk/types"
	"log"
)

type SolanaClient struct {
	client   *client.Client
	feePayer types.Account
}

func NewSolanaClient(rpcUrl string, txsFeePayer string) *SolanaClient {
	feePayer, _ := types.AccountFromBase58(txsFeePayer)
	return &SolanaClient{
		client:   client.NewClient(rpcUrl),
		feePayer: feePayer,
	}
}

func (s *SolanaClient) GetTokenMetaData(metadataAccount common.PublicKey) (tokenmeta.Metadata, error) {
	accountInfo, err := s.client.GetAccountInfo(context.Background(), metadataAccount.ToBase58())
	if err != nil {
		log.Printf("failed to get accountInfo, err: %v", err)
		return tokenmeta.Metadata{}, err
	}

	metadata, err := tokenmeta.MetadataDeserialize(accountInfo.Data)
	if err != nil {
		log.Printf("failed to parse meta account, err: %v", err)
		return tokenmeta.Metadata{}, err
	}

	return metadata, err
}

func (s *SolanaClient) UpdateTokenMetaData(metadataAccountPubKey common.PublicKey, updatedMetaData *tokenmeta.Data) (string, error) {
	recentBlockhashResponse, err := s.client.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{s.feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        s.feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				tokenmeta.UpdateMetadataAccount(tokenmeta.UpdateMetadataAccountParam{
					MetadataAccount:     metadataAccountPubKey,
					UpdateAuthority:     s.feePayer.PublicKey,
					Data:                updatedMetaData,
					NewUpdateAuthority:  nil,
					PrimarySaleHappened: nil,
				}),
			},
		}),
	})
	txId, err := s.client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Printf("failed sending updated metadata transaction, err: %v", err)
		return "", err
	}
	return txId, nil
}
