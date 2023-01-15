package handler

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/haythemsellami/SqueethAlert/common"
	"github.com/haythemsellami/SqueethAlert/decoder"
	"github.com/haythemsellami/SqueethAlert/encoder"
	"github.com/haythemsellami/SqueethAlert/loader"
	"github.com/haythemsellami/SqueethAlert/notification"
)

func PostJumboCrabHandler(client *ethclient.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var req common.BlocknativePayload

		err := c.ShouldBind(&req)

		if err != nil {
			log.Fatal(err)
		}

		jumboCrabAbi := loader.GetContractAbi("JumboCrab")
		tx := decoder.GetTransactionByHash(client, req.Hash)
		txRceipt := decoder.GetTransactionReceipt(client, tx.Hash())

		txLogs := decoder.DecodeTransactionLogs(txRceipt, &jumboCrabAbi)

		eventsStructs := encoder.PackLogIntoEventStruct(&jumboCrabAbi, txLogs)

		// fmt.Printf("%+v\n", eventsStructs)

		notification.JumboCrabNotify(eventsStructs)

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}

	return gin.HandlerFunc(fn)
}
