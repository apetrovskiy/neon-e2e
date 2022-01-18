package go_eth

import (
	"github.com/dailymotion/allure-go"
	"github.com/ethereum/go-ethereum/ethclient"
)

func connect() (*ethclient.Client, error) {
	var client *ethclient.Client
	var err error

	allure.Step(allure.Description("Connecting to network "+GetConfig().ProxyURL), allure.Action(func() {
		client, err = ethclient.Dial(GetConfig().ProxyURL)
		ReportErrorInAllure(err)

	}))
	return client, err
}
func getClient() *ethclient.Client {
	client, err := connect()
	if err != nil {
		ReportErrorInAllure(err)
	}
	return client
}
