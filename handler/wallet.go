package handler

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

type Wallet struct {
}

func (w *Wallet) GetWallet() (priKey, pubKey, address string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	priKey = hexutil.Encode(privateKeyBytes)[2:]
	fmt.Println(priKey) //

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", "", ""
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]
	fmt.Println(pubKey) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	tmpAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// 0x96216849c49358B10257cb55b28eA603c874b05E
	fmt.Println(tmpAddress)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	address = hexutil.Encode(hash.Sum(nil)[12:])
	fmt.Println(address)
	return priKey, pubKey, address
}

func GetPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	return privateKey
}
