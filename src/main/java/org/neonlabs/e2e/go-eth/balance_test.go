package go_eth

import (
	"log"
	"math/big"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestBalance(t *testing.T) {

	allure.Test(t,
		allure.Epic(Epic),
		allure.Feature(FeatureExternallyOwnedAccounts),
		allure.Story(StoryBalance),
		allure.Description("Get the latest block balance"),
		allure.Action(func() {
			client, err := connect()
			if err != nil {
				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
			}

			account := createWallet()
			if len(account.Address) == 0 {
				t.Error("Failed to create a new wallet")
			}

			balance := getLastBlockBalance(client, account.Address.Hex())
			log.Println(balance)
			assert.Equal(t, GetConfig().InitialBalance, balance, "The initial balance is wrong")
		}))
}
func TestGetSpecificBlockBalance(t *testing.T) {
	allure.Test(t,
		allure.Epic(Epic),
		allure.Feature(FeatureExternallyOwnedAccounts),
		allure.Story(StoryBalance),
		allure.Description("Get a specific block balance"),
		allure.Action(func() {
			client, err := connect()
			if err != nil {
				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
			}

			account := createWallet()
			if len(account.Address) == 0 {
				t.Error("Failed to create a new wallet")
			}

			blockNumber := big.NewInt(1001)
			balance := getSpecificBlockBalance(client, account.Address.Hex(), blockNumber)
			log.Println(balance)
			assert.Equal(t, GetConfig().InitialBalance, balance, "The initial balance is wrong")
		}))
}
func TestGetPendingBalance(t *testing.T) {
	allure.Test(t,
		allure.Epic(Epic),
		allure.Feature(FeatureExternallyOwnedAccounts),
		allure.Story(StoryBalance),
		allure.Description("Get pending balance"),
		allure.Action(func() {
			client, err := connect()
			if err != nil {
				t.Errorf("Failed to connect to %s: %o", GetConfig().ProxyURL, err)
			}

			account := createWallet()
			if len(account.Address) == 0 {
				t.Error("Failed to create a new wallet")
			}

			balance := getPendingBalance(client, account.Address.Hex())
			log.Println(balance)
			assert.Equal(t, GetConfig().InitialBalance, balance, "The initial balance is wrong")
		}))
}
