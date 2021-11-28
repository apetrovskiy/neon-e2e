package go_eth

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {

	allure.Test(t,
		allure.Epic("go-ethereum"),
		allure.Feature("go-ethereum"),
		allure.Story("go-ethereum"),
		allure.Description("Creating a new wallet"),
		allure.Action(func() {
			client, err := connect()
			if err != nil {
				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
			}

			hex := createWallet()
			if len(hex) == 0 {
				t.Error("Failed to create a new wallet")
			}

			balance := getLastBlockBalance(client, hex)
			fmt.Println(balance)
			assert.Equal(t, GetConfig().InitialBalance, balance, "The initial balance is wrong")
		}))
}
