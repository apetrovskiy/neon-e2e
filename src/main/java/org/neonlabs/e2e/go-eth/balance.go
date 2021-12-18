package go_eth

import (
	"context"
	"math/big"

	"github.com/dailymotion/allure-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func getLastBlockBalance(client *ethclient.Client, hex string) *big.Int {
	var balance *big.Int
	var err error

	allure.Step(allure.Description("Getting last block balance"), allure.Action(func() {

		account := common.HexToAddress(hex)
		balance, err = client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			log.Fatal(err)
		}
	}))
	return balance
}
func getSpecificBlockBalance(client *ethclient.Client, hex string, blockNumber *big.Int) *big.Int {
	var balance *big.Int
	var err error

	allure.Step(allure.Description("Getting a specific block balance"), allure.Action(func() {

		account := common.HexToAddress(hex)
		balance, err = client.BalanceAt(context.Background(), account, blockNumber)
		if err != nil {
			log.Fatal(err)
		}
	}))
	return balance
}
func getPendingBalance(client *ethclient.Client, hex string) *big.Int {
	var pendingBalance *big.Int
	var err error

	allure.Step(allure.Description("Getting penging balance"), allure.Action(func() {

		account := common.HexToAddress(hex)
		pendingBalance, err = client.PendingBalanceAt(context.Background(), account)
		if err != nil {
			log.Fatal(err)
		}
	}))
	return pendingBalance
}
