package main

import (
	"dapp-backend-task/base"
	"dapp-backend-task/counter"
)

func main() {
	// 查询指定区块号的区块信息
	base.QueryBlock()
	// 发送交易
	base.SendTransaction()
	// 部署合约
	counter.DeployContract()
	// 调用合约累加方法
	counter.CallContractIncrement()
	// 调用合约查询计数方法
	counter.CallContractGetCount()
}
