package base

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// 发送交易
func SendTransaction() {
	// 获取以太坊客户端
	client := GetETHClient()

	// 获取私钥
	privateKey := GetPrivateKey()
	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法将公钥转换为ECDSA格式")
	}
	// 获取公钥地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 获取交易Noce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 获取交易GasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 获取交易GasLimit
	gasLimit := uint64(100000)
	// 获取交易To地址
	toAddress := common.HexToAddress("0x16999678059a70B1438A46F523be1A3D1480eE40")
	// 获取交易Value
	value := big.NewInt(1000000000000000)
	// 创建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	// 获取交易ChainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("交易哈希：", signedTx.Hash().Hex()) // 0x06cf16b672ec7dcafa4f388aa03a7d4d581f15c6b8ff78bc21e9e97c1e95176b
}
