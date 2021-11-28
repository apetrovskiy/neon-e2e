package go_eth

import (
	"testing"

	"github.com/dailymotion/allure-go"
)

func TestNewWalletBalance(t *testing.T) {

	allure.Test(t,
		allure.Epic("go-ethereum"),
		allure.Feature("go-ethereum"),
		allure.Story("go-ethereum"),
		allure.Description("Creating a new wallet"),
		allure.Action(func() {
			accountAddress := createWallet()
			if len(accountAddress.Hash()) == 0 {
				t.Error("Failed to create a new wallet")
			}
		}))
}
