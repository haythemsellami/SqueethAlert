package encoder

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"

	mcommon "github.com/haythemsellami/SqueethAlert/common"
)

func PackLogIntoEventStruct(contractAbi *abi.ABI, transactionLogs []mcommon.TransactionLogs) []interface{} {
	var packedStructs []interface{}

	for _, event := range transactionLogs {
		if event.Event.Name == "USDCQueued" {
			var encodedEvent mcommon.USDCQueuedEvent

			err := contractAbi.UnpackIntoInterface(&encodedEvent, "USDCQueued", event.NonIndexedParams)
			if err != nil {
				panic(err)
			}

			// unpack indexed params
			encodedEvent.Depositor = event.IndexedParams[0]
			// TODO: fix this
			encodedEvent.ReceiptIndex = big.NewInt(0)

			packedStructs = append(packedStructs, encodedEvent)
		}
	}

	return packedStructs
}
