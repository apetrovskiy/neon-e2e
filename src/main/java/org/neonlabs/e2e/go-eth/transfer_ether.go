package go_eth

import (
	"context"
	// "crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// "golang.org/x/crypto/sha3"
)

func getPendingNonce() {

}

func transferEther(client *ethclient.Client, senderAccount Account, recipientAccount Account, amount string) {
	fromAddress := crypto.PubkeyToAddress(*senderAccount.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("=================================================================== 01 e")
		log.Fatal(err)
	}

	// value := big.NewInt(1000000000000000000) // in wei (1 eth)
	value := new(big.Int)
	value, _ = value.SetString(amount, 0)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("=================================================================== 02 e")
		fmt.Println(err)
		log.Fatal(err)
	}
	// gasPrice := big.NewInt(21000)

	toAddress := common.HexToAddress(recipientAccount.Address.Hex())
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("=================================================================== 03 e")
		fmt.Println(err)
		log.Fatal(err)
	}
	// chainID := big.NewInt(111)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)
	if err != nil {
		fmt.Println("=================================================================== 04 e")
		fmt.Println(err)
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("=================================================================== 05 e")
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
