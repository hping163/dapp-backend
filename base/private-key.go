package base

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

// GetPrivateKey 获取私钥
func GetPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("6855591763d1b4166eef0a211fccf52d42ecad8bc792326c64e4281352df6a65")
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}
