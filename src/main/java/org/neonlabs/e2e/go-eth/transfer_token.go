package go_eth

import (
	"github.com/dailymotion/allure-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func buildGasEstimationRequest(toAddress common.Address) []byte {
	var data []byte

	allure.Step(allure.Description("Preparing gas estimation request"), allure.Action(func() {

		transferFnSignature := []byte("transfer(address,uint256)")
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]
		log.Println(hexutil.Encode(methodID)) // 0xa9059cbb

		paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
		log.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

		amount := new(big.Int)
		amount.SetString("1000000000000000000000", 10) // sets the value to 1000 tokens, in the token denomination

		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
		log.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)

	}))

	return data
}

func transferToken(client *ethclient.Client, senderAccount Account, recipientAccount Account, tokenAmount string) {
	/*
		fromAddress := crypto.PubkeyToAddress(*senderAccount.PublicKey)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Println("=================================================================== 01 e")
			log.Fatal(err)
		}
	*/
	nonce := getPendingNonce(client, senderAccount)

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Println("=================================================================== 02 e")
	//   log.Println(err)
	// 	log.Fatal(err)
	// }
	gasPrice := big.NewInt(21000)

	/*
		toAddress := common.HexToAddress(recipientAccount.Address.Hex())
		var data []byte
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

		// chainID, err := client.NetworkID(context.Background())
		// if err != nil {
		// 	log.Println("=================================================================== 03 e")
		//   log.Println(err)
		// 	log.Fatal(err)
		// }
		chainID := big.NewInt(111)

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)
		if err != nil {
			log.Println("=================================================================== 04 e")
			log.Println(err)
			log.Fatal(err)
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Println("=================================================================== 05 e")
			log.Println(err)
			log.Fatal(err)
		}

		log.Printf("tx sent: %s", signedTx.Hash().Hex())
	*/

	toAddress := common.HexToAddress(senderAccount.Address.Hex())
	tokenAddress := common.HexToAddress(recipientAccount.Address.Hex())

	/*
		transferFnSignature := []byte("transfer(address,uint256)")
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]
		log.Println(hexutil.Encode(methodID)) // 0xa9059cbb

		paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
		log.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

		amount := new(big.Int)
		amount.SetString("1000000000000000000000", 10) // sets the value to 1000 tokens, in the token denomination

		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
		log.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

		var data []byte
		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)
	*/
	data := buildGasEstimationRequest(toAddress)

	/*
		gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			To:   &tokenAddress,
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(gasLimit) // 23256
	*/
	gasLimit := estimateGasLimit(client, data, tokenAddress)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	/*
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	*/
	chainID := getChainId(client)

	/*
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)
		if err != nil {
			log.Fatal(err)
		}
	*/
	signedTx := signTransaction(tx, types.NewEIP155Signer(chainID), senderAccount.PrivateKey)

	/*
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
	*/
	sendTransaction(client, signedTx)

	log.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
