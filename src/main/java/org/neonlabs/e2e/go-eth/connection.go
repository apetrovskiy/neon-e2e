package go_eth

import (
	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/ethereum/go-ethereum/ethclient"
)

func connect() (*ethclient.Client, error) {
	var client *ethclient.Client
	var err error

	allure.Step(allure.Description("Connecting to network "+GetConfig().ProxyURL), allure.Action(func() {
		client, err = ethclient.Dial(GetConfig().ProxyURL)
		if err != nil {
			log.Fatal(err)
		}

	}))
	return client, err
}
