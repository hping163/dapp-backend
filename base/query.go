package base

import (
	"context"
	"fmt"
	"log"
	"math/big"
)

// 查询指定区块号的区块信息
func QueryBlock() {
	// 获取以太坊客户端
	client := GetETHClient()

	// 查询指定区块号的区块信息
	block, err := client.BlockByNumber(context.Background(), big.NewInt(9173482))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("区块哈希：", block.Hash().Hex())
	fmt.Println("区块时间：", block.Time())
	fmt.Println("交易数量：", block.Transactions().Len())
}
