package go_eth

import (
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestNewWallet(t *testing.T) {

	allure.Test(t,
		allure.Epic("go-ethereum"),
		allure.Feature("Externally Owned Accounts"),
		allure.Story("Wallet"),
		allure.Description("Creating a new wallet"),
		allure.Action(func() {
			account := createWallet()
			assert.NotEqual(t, 0, len(account.Address.Hash()), "Failed to create a new wallet")
			if len(account.Address.Hash()) == 0 {

				t.Error("Failed to create a new wallet")
			}
			t.Log("Wallet has been created")

		}))
}
