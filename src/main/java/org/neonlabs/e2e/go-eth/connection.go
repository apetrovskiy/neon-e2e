package go_eth

import "github.com/ethereum/go-ethereum/ethclient"

func connect() (*ethclient.Client, error) {
	client, err := ethclient.Dial(GetConfig().ProxyURL)
	return client, err
}
