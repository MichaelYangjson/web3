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
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0x50C321Bc6303299fA18d993495f94DfdE4B49F78")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	fmt.Println(balance)

	//var wallet handler.Wallet
	//prikey, publikey, address := wallet.GetWallet()
	//fmt.Println("prikey:" + prikey + " publickey:" + publikey + " address:" + address)

	//block := handler.GetHeaderByNumber(client)
	//handler.GetTransferInfo(client, block)

	prikey := "0x242f4ef89454b99be3854804b188510d9e181951d7c384a87993351b35d47f06"
	toAddress := "0x4b2a33175fB938B7e6112437eeecD0E7717D5541"
	handler.EthTransfer(prikey[2:], client, toAddress)

}
