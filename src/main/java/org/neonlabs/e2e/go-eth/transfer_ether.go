package go_eth

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func transferEther(client *ethclient.Client, senderAccount Account, recipientAccount Account, amount string) {
	nonce := getPendingNonce(client, senderAccount)

	// value := big.NewInt(1000000000000000000) // in wei (1 eth)
	value := new(big.Int)
	value, _ = value.SetString(amount, 0)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		// TODO: logging
		fmt.Println("=================================================================== 02 e")
		fmt.Println(err)
		log.Fatal(err)
	}
	// gasPrice := big.NewInt(21000)

	// TODO: clean it up
	// toAddress := common.HexToAddress(recipientAccount.Address.Hex())
	var data []byte
	// tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTransaction(nonce, recipientAccount.Address, value, gasLimit, gasPrice, data)

	chainID := getChainId(client)

	signedTx := signTransaction(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)

	sendTransaction(client, signedTx)

	// TODO: logging
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
