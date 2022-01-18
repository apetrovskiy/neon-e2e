package go_eth

import (
	"strconv"
	"testing"

	"github.com/dailymotion/allure-go"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	firstRequestAmount  = 5
	secondRequestAmount = 3
)

func TestSingleFaucetRequest(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureExternallyOwnedAccounts),
		allure.Feature(StoryFaucet),
		// allure.Story(StoryFaucet),
		allure.Description("Single request to faucet"),
		allure.Action(func() {

			client := getClient()

			account := createWallet()
			assert.NotEqual(t, 0, len(account.Address.Hash()), FailedToCreateWallet)
			if len(account.Address.Hash()) == 0 {
				t.Error(FailedToCreateWallet)
			}
			t.Log("Wallet has been created")

			allure.Step(allure.Description("Requesting faucet for "+strconv.Itoa(firstRequestAmount)), allure.Action(func() {
				requestFaucet(account.Address.String(), firstRequestAmount)
			}))

			balance := getLastBlockBalance(client, account.Address.Hex())
			log.Println(balance)

			expectedBalance := ToWei(firstRequestAmount, 18)

			allure.Step(allure.Description("Checking that the last block balance "+balance.String()+" equals "+expectedBalance.String()), allure.Action(func() {
				assert.Equal(t, expectedBalance, balance, ResultingBalanceIsWrong)
			}))
		}))
}
func TestSubsequestFaucetRequest(t *testing.T) {
	t.Parallel()
	allure.Test(t,
		allure.Epic(FeatureExternallyOwnedAccounts),
		allure.Feature(StoryFaucet),
		// allure.Story(StoryFaucet),
		allure.Description("Subsequent requests to faucet"),
		allure.Action(func() {

			client := getClient()

			account := createWallet()
			assert.NotEqual(t, 0, len(account.Address.Hash()), FailedToCreateWallet)
			if len(account.Address.Hash()) == 0 {
				t.Error(FailedToCreateWallet)
			}
			t.Log("Wallet has been created")

			allure.Step(allure.Description("Requesting faucet for "+strconv.Itoa(firstRequestAmount)), allure.Action(func() {
				requestFaucet(account.Address.String(), firstRequestAmount)
			}))
			allure.Step(allure.Description("Requesting faucet for "+strconv.Itoa(secondRequestAmount)), allure.Action(func() {
				requestFaucet(account.Address.String(), secondRequestAmount)
			}))

			balance := getLastBlockBalance(client, account.Address.Hex())
			log.Println(balance)

			expectedBalance := ToWei(firstRequestAmount+secondRequestAmount, 18)

			allure.Step(allure.Description("Checking that the last block balance "+balance.String()+" equals "+expectedBalance.String()), allure.Action(func() {
				assert.Equal(t, expectedBalance, balance, ResultingBalanceIsWrong)
			}))
		}))
}
