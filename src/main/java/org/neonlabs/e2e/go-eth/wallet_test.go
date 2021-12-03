package go_eth

import (
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestNewWallet(t *testing.T) {

	allure.Test(t,
		allure.Epic("go-ethereum"),
		allure.Feature("go-ethereum"),
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

// TODO: clean it up
// func TestConnection111(t *testing.T) {

// 	allure.Test(t,
// 		allure.Epic("go-ethereum"),
// 		allure.Feature("go-ethereum"),
// 		allure.Story("Connection"),
// 		allure.Description("Connection to network 111"),
// 		allure.Action(func() {
// 			_, err := connect()
// 			if err != nil {
// 				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
// 			}
// 		}))
// }

// TODO: clean it up
// func TestConnection333(t *testing.T) {

// 	allure.Test(t,
// 		allure.Epic("go-ethereum"),
// 		allure.Feature("go-ethereum"),
// 		allure.Story("Connection 333"),
// 		allure.Description("Connection to network 333"),
// 		allure.Action(func() {
// 			_, err := connect()
// 			if err != nil {
// 				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
// 			}
// 		}))
// }

// TODO: clean it up
// func TestNewWallet333(t *testing.T) {

// 	allure.Test(t,
// 		allure.Epic("go-ethereum"),
// 		allure.Feature("go-ethereum"),
// 		allure.Story("Wallet 333"),
// 		allure.Description("Creating a new wallet 333"),
// 		allure.Action(func() {
// 			account := createWallet()
// 			if len(account.Address.Hash()) == 0 {
// 				t.Error("Failed to create a new wallet")
// 			}
// 			t.Log("Wallet has been created")
// 		}))
// }
