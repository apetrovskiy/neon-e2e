package go_eth

import (
	"testing"

	"github.com/dailymotion/allure-go"
)

func TestNewWalletBalance(t *testing.T) {
	allure.Epic("go-ethereum")
	allure.Story("go-ethereum")
	allure.Test(t,
		allure.Description("Creating a new wallet"),
		allure.Action(func() {
			hash := createAccount()
			if hash == nil {
				t.Error("Failed to create a new wallet")
			}
		}))
}
