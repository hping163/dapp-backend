package base

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

// GetPrivateKey 获取私钥
func GetPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}
