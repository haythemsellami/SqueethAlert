package decoder

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getTransactionByHash(client *ethclient.Client, txHashStr string) *types.Transaction {
	txHash := common.HexToHash(txHashStr)

	var tx *types.Transaction
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	return tx
}

func GetTransactionMessage(tx *types.Transaction) types.Message {
	msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}
