package main

import (
	"fmt"

	"github.com/opynfinance/MonitoringAPI-Go/loader"
	"github.com/opynfinance/gin"
)

func postJumboCrabQueued(c *gin.Context) {
	fmt.Println("Hello, world!", c)

	contractABI := loader.GetContractAbi("JumboCrab")

	fmt.Println("ABI!", contractABI)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.POST("/jumbo_crab_queued", postJumboCrabQueued)

	router.Run("localhost:8080")
}
