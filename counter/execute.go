package counter

import (
	"context"
	"crypto/ecdsa"
	"dapp-backend-task/base"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/**
 * @description: 部署counter合约
 * @return {*}
 */
func DeployContract() {
	// 获取到ETH客户端
	client := base.GetETHClient()
	// 获取私钥
	privateKey := base.GetPrivateKey()
	// 获取公钥
	publicKey := privateKey.Public()
	// 获取公钥ECDSA
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法将公钥转换为ECDSA格式")
	}
	// 获取公钥地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取Nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 获取GasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 获取GasLimit
	gasLimit := uint64(1000000)
	// 获取区块ID
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个交易
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	// 设置Nonce
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = gasLimit
	auth.Value = big.NewInt(0)
	// 部署合约
	contractAddress, tx, _, err := DeployCounter(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	// 等待合约部署完成
	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	// 打印合约地址
	fmt.Println("合约地址:", contractAddress.Hex()) // 0x9FE5d51DCdD6bE6027221b7540F81A8C4fe16a98
	fmt.Println("交易哈希:", tx.Hash().Hex())       // 0xf48857ed91bb1d6cfd724e5912d2ea40b992b3ae65497edf2ed4d2b45e46707b

}

/**
 * @description: 调用合约Increment方法
 * @return {*}
 */
func CallContractIncrement() {
	// 获取到ETH客户端
	client := base.GetETHClient()
	// 获取合约实例
	counter, err := NewCounter(common.HexToAddress("0x9FE5d51DCdD6bE6027221b7540F81A8C4fe16a98"), client)
	if err != nil {
		log.Fatal(err)
	}
	// 获取私钥
	privateKey := base.GetPrivateKey()
	// 获取区块ID
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 获取opt
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	// 执行合约方法
	tx, err := counter.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("交易哈希:", tx.Hash().Hex())
}

/**
 * @description: 调用合约GetCount方法
 * @return {*}
 */
func CallContractGetCount() {
	// 获取到ETH客户端
	client := base.GetETHClient()
	// 获取合约实例
	counter, err := NewCounter(common.HexToAddress("0x9FE5d51DCdD6bE6027221b7540F81A8C4fe16a98"), client)
	if err != nil {
		log.Fatal(err)
	}
	// 调用合约方法
	callopt := &bind.CallOpts{Context: context.Background()}
	count, err := counter.GetCount(callopt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("当前计数:", count)
}
