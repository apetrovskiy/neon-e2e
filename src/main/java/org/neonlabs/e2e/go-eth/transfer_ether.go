package go_eth

import (
	"context"
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func transferEther(client *ethclient.Client, senderAccount Account, recipientAccount Account, amount string) {
	nonce := getPendingNonce(client, senderAccount)

	value := new(big.Int)
	value, _ = value.SetString(amount, 0)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	tx := types.NewTransaction(nonce, recipientAccount.Address, value, gasLimit, gasPrice, data)

	chainID := getChainId(client)

	signedTx := signTransaction(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)

	sendTransaction(client, signedTx)

	log.Printf("tx sent: %s", signedTx.Hash().Hex())
}
