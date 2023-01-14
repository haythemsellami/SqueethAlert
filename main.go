package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/haythemsellami/SqueethAlert/handler"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	PROVIDER_URL := os.Getenv("PROVIDER_URL")

	router := gin.Default()

	client, err := ethclient.Dial(PROVIDER_URL)
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/jumbo_crab_queued", handler.PostJumboCrabHandler(client))

	router.Run("localhost:8080")
}
