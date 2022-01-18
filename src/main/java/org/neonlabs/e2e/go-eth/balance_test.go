package go_eth

import (
	"math/big"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestBalance(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureExternallyOwnedAccounts),
		allure.Feature(StoryBalance),
		// allure.Story(StoryBalance),
		allure.Description("Get the latest block balance"),
		allure.Action(func() {

			client := getClient()

			account := createWallet(GetConfig().RequestAmount)
			assert.NotEqual(t, 0, len(account.Address.Hash()), FailedToCreateWallet)
			if len(account.Address) == 0 {
				t.Error(FailedToCreateWallet)
			}

			balance := getLastBlockBalance(client, account.Address.Hex())
			log.Println(balance)

			expectedBalance := ToWei(GetConfig().RequestAmount, 18)
			allure.Step(allure.Description("Checking that the last block balance "+balance.String()+" equals "+expectedBalance.String()), allure.Action(func() {
				assert.Equal(t, expectedBalance, balance, InitialBalanceIsWrong)
			}))
		}))
}
func TestGetSpecificBlockBalance(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureExternallyOwnedAccounts),
		allure.Feature(StoryBalance),
		// allure.Story(StoryBalance),
		allure.Description("Get a specific block balance"),
		allure.Action(func() {

			client := getClient()

			account := createWallet(GetConfig().RequestAmount)
			assert.NotEqual(t, 0, len(account.Address.Hash()), FailedToCreateWallet)
			if len(account.Address) == 0 {
				t.Error(FailedToCreateWallet)
			}

			blockNumber := big.NewInt(1001)
			balance := getSpecificBlockBalance(client, account.Address.Hex(), blockNumber)
			log.Println(balance)

			expectedBalance := ToWei(GetConfig().RequestAmount, 18)
			allure.Step(allure.Description("Checking that the specific block balance "+balance.String()+" equals "+expectedBalance.String()), allure.Action(func() {
				assert.Equal(t, expectedBalance, balance, ResultingBalanceIsWrong)
			}))
		}))
}
func TestGetPendingBalance(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureExternallyOwnedAccounts),
		allure.Feature(StoryBalance),
		// allure.Story(StoryBalance),
		allure.Description("Get pending balance"),
		allure.Action(func() {

			client := getClient()

			account := createWallet(GetConfig().RequestAmount)
			assert.NotEqual(t, 0, len(account.Address.Hash()), FailedToCreateWallet)
			if len(account.Address) == 0 {
				t.Error(FailedToCreateWallet)
			}

			balance := getPendingBalance(client, account.Address.Hex())
			log.Println(balance)

			expectedBalance := ToWei(GetConfig().RequestAmount, 18)
			allure.Step(allure.Description("Checking that the pending balance "+balance.String()+" equals "+expectedBalance.String()), allure.Action(func() {
				assert.Equal(t, expectedBalance, balance, ResultingBalanceIsWrong)
			}))
		}))
}
