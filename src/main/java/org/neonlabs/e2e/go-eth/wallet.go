package go_eth

import (
	"crypto/ecdsa"
	"hash"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type Account struct {
	Address         common.Address
	PrivateKey      *ecdsa.PrivateKey
	PrivateKeyBytes []byte
	PrivateKeyHex   string
	PublicKey       *ecdsa.PublicKey
	PublicKeyBytes  []byte
	Hash            hash.Hash
}

func createWallet() *Account {

	var accountData Account

	allure.Step(allure.Description("Creating a wallet"), allure.Action(func() {

		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		hexString := hexutil.Encode(privateKeyBytes)[:2]
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		address := crypto.PubkeyToAddress(*publicKeyECDSA)
		hash := sha3.NewLegacyKeccak256()
		hash.Write(publicKeyBytes[1:])

		log.Println(hexString)
		log.Println(publicKeyECDSA)
		log.Println(publicKeyBytes)
		log.Println(address)
		log.Println(hash)
		log.Println(hexutil.Encode(hash.Sum(nil)[12:]))

		accountData = Account{
			Address:         address,
			PrivateKey:      privateKey,
			PrivateKeyBytes: privateKeyBytes,
			PrivateKeyHex:   hexString,
			PublicKey:       publicKeyECDSA,
			PublicKeyBytes:  publicKeyBytes,
			Hash:            hash,
		}
	}))

	return &accountData
}
