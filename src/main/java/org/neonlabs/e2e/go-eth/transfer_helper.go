package go_eth

import (
	"context"
	"crypto/ecdsa"

	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dailymotion/allure-go"
	"github.com/ethereum/go-ethereum"
)

func getPendingNonce(client *ethclient.Client, senderAccount Account) uint64 {
	var nonce uint64
	var err error

	allure.Step(allure.Description("Getting pending nonce"), allure.Action(func() {

		nonce, err = client.PendingNonceAt(context.Background(), senderAccount.Address)
		ReportErrorInAllure(err)
	}))

	return nonce

}

func getChainId(client *ethclient.Client) *big.Int {
	var chainId *big.Int
	var err error

	allure.Step(allure.Description("Getting chain id"), allure.Action(func() {

		chainId, err = client.NetworkID(context.Background())
		ReportErrorInAllure(err)

	}))

	return chainId
}

func signTransaction(tx *types.Transaction, signer types.EIP155Signer, privateKey *ecdsa.PrivateKey) *types.Transaction {
	var signedTx *types.Transaction
	var err error

	allure.Step(allure.Description("Siging transaction"), allure.Action(func() {

		signedTx, err = types.SignTx(tx, signer, privateKey)
		ReportErrorInAllure(err)

	}))

	return signedTx
}

func sendTransaction(client *ethclient.Client, signedTx *types.Transaction) {

	allure.Step(allure.Description("Sending transaction"), allure.Action(func() {

		err := client.SendTransaction(context.Background(), signedTx)
		ReportErrorInAllure(err)

	}))

}

func estimateGasLimit(client *ethclient.Client, data []byte, tokenAddress common.Address) uint64 {
	var gasLimit uint64
	var err error

	allure.Step(allure.Description("Estimating gas limit"), allure.Action(func() {

		gasLimit, err = client.EstimateGas(context.Background(), ethereum.CallMsg{
			To:   &tokenAddress,
			Data: data,
		})
		ReportErrorInAllure(err)
		log.Println(gasLimit) // 23256
	}))

	return gasLimit
}
