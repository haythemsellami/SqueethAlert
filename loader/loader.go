package loader

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func GetContractAbi(contractName string) abi.ABI {

	contractABI, err := abi.JSON(strings.NewReader(getLocalABI(contractName)))

	if err != nil {
		log.Fatal(err)
	}

	return contractABI
}

func getLocalABI(abiName string) string {
	abiPath := "../abis/" + abiName + ".json"

	abiFile, err := os.Open(abiPath)
	if err != nil {
		log.Fatal(err)
	}

	defer abiFile.Close()

	result, err := io.ReadAll(abiFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}
