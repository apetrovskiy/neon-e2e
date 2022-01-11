package go_eth

import (
	"context"
	"math/big"

	"github.com/dailymotion/allure-go"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func transferEther(client *ethclient.Client, senderAccount Account, recipientAccount Account, amount string) {
	allure.Step(allure.Description("Transferring Ether"), allure.Action(func() {
		nonce := getPendingNonce(client, senderAccount)

		value := new(big.Int)
		value, _ = value.SetString(amount, 0)
		gasLimit := uint64(21000) // in units
		gasPrice, err := client.SuggestGasPrice(context.Background())
		ReportErrorInAllure(err)

		var data []byte
		var tx, signedTx *types.Transaction
		var chainID *big.Int
		allure.Step(allure.Description("Creating a transaction"), allure.Action(func() {
			tx = types.NewTransaction(nonce, recipientAccount.Address, value, gasLimit, gasPrice, data)
		}))
		allure.Step(allure.Description("Getting chain id"), allure.Action(func() {
			chainID = getChainId(client)
		}))

		allure.Step(allure.Description("Signing the transaction"), allure.Action(func() {
			signedTx = signTransaction(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)
		}))

		allure.Step(allure.Description("Sending the transaction"), allure.Action(func() {
			sendTransaction(client, signedTx)
		}))

		allure.Step(allure.Description("Logging the result"), allure.Action(func() {
			log.Printf("tx sent: %s", signedTx.Hash().Hex())
		}))
	}))
}
