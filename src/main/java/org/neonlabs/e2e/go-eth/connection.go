package go_eth

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func connect() (*ethclient.Client, error) {
	client, err := ethclient.Dial(GetConfig().ProxyURL)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}
