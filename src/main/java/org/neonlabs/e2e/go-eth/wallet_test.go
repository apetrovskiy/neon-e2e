package go_eth

import (
	"testing"

	"github.com/dailymotion/allure-go"
)

func TestNewWalletBalance(t *testing.T) {

	allure.Test(t,
		allure.Epic("go-ethereum"),
		allure.Feature("go-ethereum"),
		allure.Story("Wallet"),
		allure.Description("Creating a new wallet"),
		allure.Action(func() {
			account := createWallet()
			if len(account.Address.Hash()) == 0 {
				t.Error("Failed to create a new wallet")
			}
			t.Log("Wallet has been created")
		}))
}

func TestConnection111(t *testing.T) {

	allure.Test(t,
		allure.Epic("go-ethereum"),
		allure.Feature("go-ethereum"),
		allure.Story("Connection"),
		allure.Description("Connection to network 111"),
		allure.Action(func() {
			_, err := connect()
			if err != nil {
				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
			}
		}))
}
