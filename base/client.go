package base

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// GetETHClient 获取以太坊客户端（https）
func GetETHClient() *ethclient.Client {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/4fdeb7812b9546b4b9027a9187e82bbc")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// GetWSClient 获取以太坊客户端（websocket）
func GetWSClient() *ethclient.Client {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/4fdeb7812b9546b4b9027a9187e82bbc")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
