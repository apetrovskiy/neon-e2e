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
func getSpecificBlockBalance(client *ethclient.Client, hex string, blockNumber *big.Int) *big.Int {
	account := common.HexToAddress(hex)
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}
func getPendingBalance(client *ethclient.Client, hex string) *big.Int {
	account := common.HexToAddress(hex)
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	return pendingBalance
}
