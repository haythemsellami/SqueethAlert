package decoder

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	mcommon "github.com/haythemsellami/SqueethAlert/common"
)

func GetTransactionByHash(client *ethclient.Client, txHashStr string) *types.Transaction {
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

func DecodeTransactionInputData(contractAbi *abi.ABI, data []byte) {
	// The first 4 bytes of the t represent the ID of the method in the ABI
	// https://docs.soliditylang.org/en/v0.5.3/abi-spec.html#function-selector
	methodSigData := data[:4]
	method, err := contractAbi.MethodById(methodSigData)
	if err != nil {
		log.Fatal(err)
	}

	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Method Name: %s\n", method.Name)
	// fmt.Printf("Method inputs: %v\n", MapToJson(inputsMap))
}

func GetTransactionReceipt(client *ethclient.Client, txHash common.Hash) *types.Receipt {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	return receipt
}

func DecodeTransactionLogs(receipt *types.Receipt, contractAbi *abi.ABI) []mcommon.TransactionLogs {
	var transactionLogs []mcommon.TransactionLogs

	for _, vLog := range receipt.Logs {
		var transactionLog mcommon.TransactionLogs

		event, err := contractAbi.EventByID(vLog.Topics[0])
		if err != nil {
			// ignoring this as mostly the eventId is not located in the given ABI
			continue
		}
		transactionLog.Event = event
		// topic[1:] is other indexed params in event
		if len(vLog.Topics) > 1 {
			var IndexedParamsArray []common.Address

			for _, param := range vLog.Topics[1:] {
				IndexedParamsArray = append(IndexedParamsArray, common.HexToAddress(param.Hex()))
			}
			transactionLog.IndexedParams = IndexedParamsArray
		}

		// if len(vLog.Data) > 0 {
		// 	fmt.Println("vLog.Data", vLog.Data)
		// 	outputDataMap := make(map[string]interface{})

		// 	err = contractAbi.UnpackIntoMap(outputDataMap, transactionLog.Event.Name, vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	for _, logValue := range outputDataMap {
		// 		switch logValue.(type) {
		// 		case *big.Int:
		// 			transactionLog.NonIndexedParams = append(transactionLog.NonIndexedParams, logValue.(*big.Int).String())
		// 		default:
		// 			transactionLog.NonIndexedParams = append(transactionLog.NonIndexedParams, logValue.(string))
		// 		}
		// 	}
		// }

		transactionLog.NonIndexedParams = vLog.Data
		transactionLogs = append(transactionLogs, transactionLog)
	}

	return transactionLogs
}
