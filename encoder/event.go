package encoder

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"

	mcommon "github.com/haythemsellami/SqueethAlert/common"
)

func PackLogIntoEventStruct(contractAbi *abi.ABI, transactionLogs []mcommon.TransactionLogs) []interface{} {
	var packedStructs []interface{}

	for _, event := range transactionLogs {
		switch event.Event.Name {
		case "USDCQueued":
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
		case "CrabQueued":
			var encodedEvent mcommon.CrabQueuedEvent

			err := contractAbi.UnpackIntoInterface(&encodedEvent, "CrabQueued", event.NonIndexedParams)
			if err != nil {
				panic(err)
			}

			// unpack indexed params
			encodedEvent.Withdrawer = event.IndexedParams[0]
			// TODO: fix this and convert the right amount
			encodedEvent.ReceiptIndex = big.NewInt(0)

			packedStructs = append(packedStructs, encodedEvent)
		default:
			fmt.Println("Couldn't encode event struct")
		}
	}

	return packedStructs
}
