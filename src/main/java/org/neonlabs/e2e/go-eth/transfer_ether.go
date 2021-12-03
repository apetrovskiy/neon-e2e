package go_eth

import (
	"context"
	"crypto/ecdsa"
	// "crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	// "golang.org/x/crypto/sha3"
	"github.com/dailymotion/allure-go"
)

func getPendingNonce(client *ethclient.Client, senderAccount Account) uint64 {
	var nonce uint64
	var err error

	allure.Step(allure.Description("Getting pending nonce"), allure.Action(func() {

		nonce, err = client.PendingNonceAt(context.Background(), senderAccount.Address)
		if err != nil {
			fmt.Println("=================================================================== 01 e")
			log.Fatal(err)
		}
	}))

	return nonce

}

func getChainId(client *ethclient.Client) *big.Int {
	var chainId *big.Int
	var err error

	allure.Step(allure.Description("Getting chain id"), allure.Action(func() {

		chainId, err = client.NetworkID(context.Background())
		if err != nil {
			fmt.Println("=================================================================== 03 e")
			fmt.Println(err)
			log.Fatal(err)
		}

	}))

	return chainId
}

func signTransaction(tx *types.Transaction, signer types.EIP155Signer, privateKey *ecdsa.PrivateKey) *types.Transaction {
	var signedTx *types.Transaction
	var err error

	allure.Step(allure.Description("Siging transaction"), allure.Action(func() {

		signedTx, err = types.SignTx(tx, signer, privateKey)
		if err != nil {
			fmt.Println("=================================================================== 04 e")
			fmt.Println(err)
			log.Fatal(err)
		}

	}))

	return signedTx
}

func sendTransaction(client *ethclient.Client, signedTx *types.Transaction) {

	allure.Step(allure.Description("Sending transaction"), allure.Action(func() {

		err := client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Println("=================================================================== 05 e")
			fmt.Println(err)
			log.Fatal(err)
		}

	}))

}

func transferEther(client *ethclient.Client, senderAccount Account, recipientAccount Account, amount string) {
	// fromAddress := crypto.PubkeyToAddress(*senderAccount.PublicKey)
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	/*
		  nonce, err := client.PendingNonceAt(context.Background(), senderAccount.Address)
			if err != nil {
				fmt.Println("=================================================================== 01 e")
				log.Fatal(err)
			}
	*/
	nonce := getPendingNonce(client, senderAccount)

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

	// toAddress := common.HexToAddress(recipientAccount.Address.Hex())
	var data []byte
	// tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTransaction(nonce, recipientAccount.Address, value, gasLimit, gasPrice, data)

	/*
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			fmt.Println("=================================================================== 03 e")
			fmt.Println(err)
			log.Fatal(err)
		}
	*/
	// temporarily
	// chainID := big.NewInt(111)
	chainID := getChainId(client)

	/*
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)
		if err != nil {
			fmt.Println("=================================================================== 04 e")
			fmt.Println(err)
			log.Fatal(err)
		}
	*/

	signedTx := signTransaction(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)

	/*
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Println("=================================================================== 05 e")
			fmt.Println(err)
			log.Fatal(err)
		}
	*/
	sendTransaction(client, signedTx)

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
