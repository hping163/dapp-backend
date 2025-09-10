package base

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// GetETHClient 获取以太坊客户端（https）
func GetETHClient() *ethclient.Client {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// GetWSClient 获取以太坊客户端（websocket）
func GetWSClient() *ethclient.Client {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
