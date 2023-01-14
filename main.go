package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/haythemsellami/SqueethAlert/common"
	"github.com/haythemsellami/SqueethAlert/decoder"
	"github.com/haythemsellami/SqueethAlert/encoder"
	"github.com/haythemsellami/SqueethAlert/loader"
)

func postJumboCrabHandler(client *ethclient.Client) gin.HandlerFunc {
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

		fmt.Printf("%+v\n", eventsStructs)

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}

	return gin.HandlerFunc(fn)
}

func main() {
	const PROVIDER_URL = "https://mainnet.infura.io/v3/9a1eacc6b18f436dab839c1713616fd1"

	router := gin.Default()

	client, err := ethclient.Dial(PROVIDER_URL)
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/jumbo_crab_queued", postJumboCrabHandler(client))

	router.Run("localhost:8080")
}
