package main

import (
	"context"
	"eth_test/handler"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0x47e31E940f2d1F8FeC009d07e266F893db154A53")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	fmt.Println(balance)

	var wallet handler.Wallet
	prikey, publikey, address := wallet.GetWallet()
	fmt.Println("prikey:" + prikey + " publickey:" + publikey + " address:" + address)

	block := handler.GetHeaderByNumber(client)
	handler.GetTransferInfo(client, block)

}
