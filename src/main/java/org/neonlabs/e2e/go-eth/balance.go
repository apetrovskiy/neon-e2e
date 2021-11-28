package go_eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"log"

	"math/big"
)

func getLastBlockBalance(client *ethclient.Client, hex string) *big.Int {
	account := common.HexToAddress(hex)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}
